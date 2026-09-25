package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	Content []anthropicContent `json:"content"`
	Error   *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

type Client struct {
	baseURL    string
	apiKey     string
	model      string
	httpClient *http.Client
	log        *log.Logger
}

func NewClient(cfg config.AIConfig, logger *log.Logger) *Client {
	if cfg.BaseURL == "" || cfg.APIKey == "" {
		return nil
	}
	return &Client{
		baseURL: strings.TrimRight(cfg.BaseURL, "/"),
		apiKey:  cfg.APIKey,
		model:   cfg.Model,
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
	c.logAI(">>> SYSTEM:\n%s", systemPrompt)
	c.logAI(">>> USER:\n%s", userPrompt)

	reqBody := anthropicRequest{
		Model:     c.model,
		MaxTokens: 4096,
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
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(bodyBytes))
	if err != nil {
		return "", fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("http request: %w", err)
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		c.logAI("<<< ERROR %d: %s", resp.StatusCode, string(respBytes))
		return "", fmt.Errorf("API error %d: %s", resp.StatusCode, string(respBytes))
	}

	var result anthropicResponse
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return "", fmt.Errorf("unmarshal response: %w", err)
	}

	if result.Error != nil {
		return "", fmt.Errorf("API error: %s", result.Error.Message)
	}

	// Extract text from response. With web search, the response may contain
	// server_tool_use and web_search_tool_result blocks before the text block.
	var textParts []string
	for _, block := range result.Content {
		if block.Type == "text" {
			textParts = append(textParts, block.Text)
		}
	}
	content := strings.Join(textParts, "\n")
	if content == "" {
		return "", fmt.Errorf("empty response content")
	}

	c.logAI("<<< RESPONSE:\n%s", content)
	return content, nil
}

func (c *Client) logAI(format string, args ...interface{}) {
	if c.log != nil {
		c.log.Printf(format, args...)
	}
}
