package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/staroffish/am/internal/config"
)

// Anthropic Messages API types

type anthropicTool struct {
	Type        string      `json:"type"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	InputSchema interface{} `json:"input_schema,omitempty"`
}

type anthropicContent struct {
	Type string `json:"type"`
	Text string `json:"text,omitempty"`
}

type anthropicMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type anthropicRequest struct {
	Model     string             `json:"model"`
	MaxTokens int                `json:"max_tokens"`
	System    string             `json:"system"`
	Messages  []anthropicMessage `json:"messages"`
	Tools     []anthropicTool    `json:"tools,omitempty"`
	Stream    bool               `json:"stream,omitempty"`
}

type anthropicResponse struct {
	Content    []anthropicContent `json:"content"`
	StopReason string             `json:"stop_reason"`
	Usage      struct {
		InputTokens  int `json:"input_tokens"`
		OutputTokens int `json:"output_tokens"`
	} `json:"usage"`
	Error *struct {
		Type    string `json:"type,omitempty"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

const defaultMaxTokens = 8192

type Client struct {
	baseURL    string
	apiKey     string
	model      string
	maxTokens  int
	httpClient *http.Client
	log        *log.Logger
}

func NewClient(cfg config.AIConfig, logger *log.Logger) *Client {
	if cfg.BaseURL == "" || cfg.APIKey == "" {
		return nil
	}
	maxTokens := cfg.MaxTokens
	if maxTokens <= 0 {
		maxTokens = defaultMaxTokens
	}
	return &Client{
		baseURL:   strings.TrimRight(cfg.BaseURL, "/"),
		apiKey:    cfg.APIKey,
		model:     cfg.Model,
		maxTokens: maxTokens,
		httpClient: &http.Client{
			Timeout: 120 * time.Second,
		},
		log: logger,
	}
}

// webSearchTool is the DeepSeek-compatible web search tool definition.
// DeepSeek supports server-side web search via the Anthropic Messages API.
var webSearchTool = anthropicTool{
	Type: "web_search_20250305",
	Name: "web_search",
}

// Chat sends a message without web search.
func (c *Client) Chat(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	return c.chat(ctx, systemPrompt, userPrompt, nil)
}

// ChatWithSearch sends a message with web search tool enabled.
func (c *Client) ChatWithSearch(ctx context.Context, systemPrompt, userPrompt string) (string, error) {
	return c.chat(ctx, systemPrompt, userPrompt, []anthropicTool{webSearchTool})
}

func (c *Client) chat(ctx context.Context, systemPrompt, userPrompt string, tools []anthropicTool) (string, error) {
	reqBody := anthropicRequest{
		Model:     c.model,
		MaxTokens: c.maxTokens,
		System:    systemPrompt,
		Messages: []anthropicMessage{
			{Role: "user", Content: userPrompt},
		},
		Tools:  tools,
		Stream: false,
	}

	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("marshal request: %w", err)
	}

	url := c.baseURL + "/v1/messages"

	// 完整记录一次会话，方便排查：请求（含 system/user 原文与请求参数）与响应（原始 JSON）。
	c.logAI(">>> REQUEST url=%s model=%s max_tokens=%d web_search=%t stream=%t",
		url, c.model, c.maxTokens, len(tools) > 0, reqBody.Stream)
	c.logAI(">>> SYSTEM:\n%s", systemPrompt)
	c.logAI(">>> USER:\n%s", userPrompt)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	start := time.Now()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logAI("<<< TRANSPORT ERROR after %v: %v", time.Since(start), err)
		return "", fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		c.logAI("<<< READ ERROR (status=%d) after %v: %v", resp.StatusCode, time.Since(start), err)
		return "", fmt.Errorf("read response: %w", err)
	}

	// 原始响应体始终落盘（含 thinking / tool_use / error 等所有内容块）。
	c.logAI("<<< HTTP %d in %v, raw response body (%d bytes):\n%s",
		resp.StatusCode, time.Since(start), len(respBytes), string(respBytes))

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error %d: %s", resp.StatusCode, string(respBytes))
	}

	var result anthropicResponse
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", fmt.Errorf("unmarshal response: %w (raw: %s)", err, string(respBytes))
	}

	if result.Error != nil {
		return "", fmt.Errorf("API error: %s", result.Error.Message)
	}

	// Extract text from response. With web search, the response may contain
	// server_tool_use, thinking and web_search_tool_result blocks before the text block.
	var textParts []string
	for _, block := range result.Content {
		if block.Type == "text" {
			textParts = append(textParts, block.Text)
		}
	}
	content := strings.Join(textParts, "\n")
	if content == "" {
		c.logAI("<<< EMPTY TEXT: stop_reason=%q usage(in=%d,out=%d) block_types=[%s]",
			result.StopReason, result.Usage.InputTokens, result.Usage.OutputTokens, c.blockTypes(result))
		return "", fmt.Errorf("empty response content (stop_reason=%s, usage out=%d%s)",
			c.stopReasonText(result.StopReason), result.Usage.OutputTokens, c.blockTypeHint(result))
	}

	c.logAI("<<< RESPONSE TEXT (stop_reason=%q usage in=%d out=%d):\n%s",
		result.StopReason, result.Usage.InputTokens, result.Usage.OutputTokens, content)
	return content, nil
}

// blockTypes 汇总响应里的内容块类型，用于判断"空的"响应到底回了什么（如 thinking）。
func (c *Client) blockTypes(result anthropicResponse) string {
	if len(result.Content) == 0 {
		return "none"
	}
	counts := make(map[string]int, len(result.Content))
	for _, block := range result.Content {
		counts[block.Type]++
	}
	types := make([]string, 0, len(counts))
	for t, n := range counts {
		types = append(types, fmt.Sprintf("%s×%d", t, n))
	}
	sort.Strings(types)
	return strings.Join(types, ",")
}

func (c *Client) blockTypeHint(result anthropicResponse) string {
	types := c.blockTypes(result)
	if types == "none" {
		return ""
	}
	return ", blocks=" + types
}

// stopReasonText 把 stop_reason 翻成人话，便于直接看日志定位。
func (c *Client) stopReasonText(reason string) string {
	switch reason {
	case "":
		return "未返回"
	case "end_turn":
		return "end_turn(正常结束但无文本)"
	case "max_tokens":
		return "max_tokens(输出被 max_tokens 截断)"
	case "stop_sequence":
		return "stop_sequence"
	case "tool_use":
		return "tool_use(只发起了工具调用)"
	case "pause_turn":
		return "pause_turn(工具调用后暂停)"
	case "refusal":
		return "refusal(模型拒答)"
	default:
		return reason
	}
}

func (c *Client) logAI(format string, args ...interface{}) {
	if c.log != nil {
		c.log.Printf(format, args...)
	}
}
