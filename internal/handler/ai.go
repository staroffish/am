package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/staroffish/am/internal/ai"
)

type AIHandler struct {
	aiClient *ai.Client
	log      *log.Logger
}

func NewAIHandler(aiClient *ai.Client, logger *log.Logger) *AIHandler {
	return &AIHandler{aiClient: aiClient, log: logger}
}

func (h *AIHandler) Register(e *echo.Group) {
	e.POST("/generate-rule", h.GenerateRule)
}

type aiRuleResult struct {
	JapaneseName string `json:"japanese_name"`
	Regex        string `json:"regex"`
	IsInvalid    bool   `json:"is_invalid,omitempty"`
	RawResponse  string `json:"raw_response,omitempty"`
}

func (h *AIHandler) GenerateRule(c echo.Context) error {
	var req struct {
		MagnetName string `json:"magnet_name"`
		UserName   string `json:"user_name"`
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if req.MagnetName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "magnet_name is required")
	}

	if h.aiClient == nil {
		return echo.NewHTTPError(http.StatusNotImplemented, "AI not configured")
	}

	ctx := c.Request().Context()

	// === Step 1: Web search for Japanese name ===
	h.log.Printf("AI step1 search-name: magnet=%s", req.MagnetName)
	start := time.Now()
	japaneseName, err := h.searchJapaneseName(ctx, req.MagnetName, req.UserName)
	elapsed := time.Since(start)
	h.log.Printf("AI step1 took %v, name=%s", elapsed, japaneseName)
	if err != nil {
		h.log.Printf("AI step1 error: %v", err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("AI name search failed: %v", err))
	}

	// === Step 2: Generate regex ===
	h.log.Printf("AI step2 generate-regex: name=%s, magnet=%s", japaneseName, req.MagnetName)
	start = time.Now()
	regexStr, err := h.generateRegexWithRetry(ctx, japaneseName, req.MagnetName)
	elapsed = time.Since(start)
	h.log.Printf("AI step2 took %v, regex=%s", elapsed, regexStr)
	if err != nil {
		h.log.Printf("AI step2 error: %v", err)
		return c.JSON(http.StatusOK, aiRuleResult{
			JapaneseName: japaneseName,
			Regex:        "",
			IsInvalid:    true,
			RawResponse:  fmt.Sprintf("regex generation failed: %v", err),
		})
	}

	result := aiRuleResult{JapaneseName: japaneseName, Regex: regexStr}
	if regexStr != "" {
		testStr := fmt.Sprintf(regexStr, 1)
		if _, err := regexp.Compile(testStr); err != nil {
			result.IsInvalid = true
			result.RawResponse = fmt.Sprintf("regex compile error: %v", err)
		}
	}

	return c.JSON(http.StatusOK, result)
}

// searchJapaneseName calls AI to find the Japanese name.
func (h *AIHandler) searchJapaneseName(ctx context.Context, magnetName, userName string) (string, error) {
	query := fmt.Sprintf("%s\n帮我看下这个标题中动漫标题的日语名是什么，返回JSON格式：{\"name\":\"日语名\"}。如果找不到返回：{\"name\":\"\"}", magnetName)
	if userName != "" {
		query += fmt.Sprintf("（中文名: %s）", userName)
	}

	response, err := h.aiClient.ChatWithSearch(ctx, searchNameSystemPrompt, query)
	if err != nil {
		return "", err
	}

	raw := strings.TrimSpace(response)
	h.log.Printf("AI search raw response: %s", raw)

	// Try to parse JSON response
	name := parseJSONName(raw)
	if name == "" {
		return "", fmt.Errorf("AI未找到该动漫，原始响应: %s", raw)
	}

	return name, nil
}

// searchNameSystemPrompt 约束第一步的输出格式：必须是纯 JSON、必须用 ASCII 直引号。
// 注意：不要用代码兜底去修全角引号，靠提示词让模型直接输出合规格式。
const searchNameSystemPrompt = `你是动漫标题查询工具。输入是一条动漫种子（磁链）名称，你要用联网搜索查出这部动漫的官方日语原名。

搜索与判断：
- 优先查日文维基百科、bangumi.tv、萌娘百科、anidb.net
- 忽略种子名里的集数、分辨率、字幕组、来源标签（如 Baha/CR/ABEMA/1080p/MKV/简繁内封等），只认作品名
- 必须命中同一部作品；找不到就当作未找到，绝不猜测、绝不编造、绝不用记忆代替搜索

输出（严格遵守）：
- 只输出一个 JSON 对象，不要 markdown 代码块、不要前后缀、不要解释
- key 固定为 name：{"name":"日语原名"}
- 引号必须使用英文半角双引号 " ，禁止使用中文全角引号 “ ” 、禁止单引号
- 日语原名必须来自你搜到的网页标题，原样照抄，不要加书名号、括号、罗马音、中文翻译或任何说明文字
- 找不到任何匹配 → 只输出 {"name":""}

示例（正确）：{"name":"野生のラスボスが現れた！"}
示例（错误，禁止）：{“name”:“野生のラスボスが現れた！”}`

// parseJSONName extracts "name" field from a JSON response like {"name":"xxx"}
func parseJSONName(raw string) string {
	// Try to find {"name":"..."} pattern
	idx := strings.Index(raw, `"name"`)
	if idx < 0 {
		return ""
	}
	// Find the value after "name":
	colon := strings.Index(raw[idx:], ":")
	if colon < 0 {
		return ""
	}
	val := strings.TrimSpace(raw[idx+colon+1:])
	// Strip quotes
	val = strings.Trim(val, `"`)
	// Take everything until next quote or }
	end := strings.IndexAny(val, `"}`)
	if end > 0 {
		val = val[:end]
	}
	val = strings.TrimSpace(val)
	return val
}

// generateRegex generates a regex pattern using the confirmed Japanese name and magnet name.
func (h *AIHandler) generateRegex(ctx context.Context, japaneseName, magnetName string) (string, error) {
	user := fmt.Sprintf("动漫日文名: %s\n磁链名称: %s\n\n请为这条磁链生成一个Go正则模板。只返回正则字符串，不要markdown、JSON或解释。", japaneseName, magnetName)

	return h.aiClient.Chat(ctx, regexSystemPrompt, user)
}

// generateRegexWithRetry 处理模型偶尔返回空正文（stop_reason=max_tokens 或只回思考块）的情况：
// 重试一次；仍失败则把"动漫名 + 原始错误"带出去，便于定位是模型问题还是 prompt 问题。
func (h *AIHandler) generateRegexWithRetry(ctx context.Context, japaneseName, magnetName string) (string, error) {
	const attempts = 2
	var lastErr error
	for i := 1; i <= attempts; i++ {
		if i > 1 {
			h.log.Printf("AI step2 retry %d/%d: name=%s", i-1, attempts-1, japaneseName)
		}
		regexStr, err := h.generateRegex(ctx, japaneseName, magnetName)
		if err == nil && strings.TrimSpace(regexStr) != "" {
			return strings.TrimSpace(regexStr), nil
		}
		if err != nil {
			lastErr = err
		} else {
			lastErr = fmt.Errorf("empty regex returned")
		}
		if ctx.Err() != nil {
			break
		}
	}
	return "", fmt.Errorf("动漫「%s」的正则生成失败（已重试 %d 次）: %w", japaneseName, attempts-1, lastErr)
}

// regexSystemPrompt 是第二步的提示词。除原有风格约束外，显式要求"必须输出非空正则"，
// 避免模型把整个输出预算花在思考上、正文留空。
const regexSystemPrompt = `生成一个Go fmt.Sprintf正则模板，用于匹配动漫种子磁链名称。

输出格式（最重要，违反即作废）：
- 必须输出且只输出一行正则字符串本身，不能为空
- 禁止 markdown、代码块、反引号、JSON、引号包裹、前后缀说明、任何解释文字
- 无论分析过程多长，最后都必须输出这一行正则

格式说明：%d（或%02d、%03d）会在扫描时被替换为集数。

规则：
- 捕获组1 = 集数: (%02d) 或 (\d+)，集数必须写在方括号或圆括号内，如 [(%02d)] 或 \(%02d\)
- 可选捕获组2 = 结束集数（合集时使用）: (\d{2})?
- 匹配磁链中的实际文字（罗马音/英文/中文），不要用日文标题
- 包含磁链中的字幕组标签（如 [ANi]、[Nix-Raws]、[黒ネズミたち]）
- 对 [ ] ( ) . + * ? ^ $ { } | \ / 等特殊字符做转义
- 零填充: %02d 匹配 "01"，%d/\d+ 匹配无填充的 "1"
- 参考格式: \[字幕组\]  ?.*?罗马音名.+\[(%02d)\]\[1080P\]\[字幕类型\].+来源
- 关键：磁链中的版本标签和来源标签必须原样保留精确匹配。如 [简繁内封]/[简体内嵌]/[CHT]/[MultiSub] 等版本标签，Baha/CR/WEB-DL/NF/AMZN 等来源标签都必须写到正则中
- 只有分辨率后面的纯文件格式细节（如 AVC AAC MKV/HEVC EAC3 MP4 等）可以用 .+? 模糊化
- 参考格式: \[字幕组\]  ?.*?罗马音名.+\[(%02d)\]\[1080P\]\[版本标签\].+来源
- 中文必须同时兼容简繁体，如 [无無]、[战戰]、[转転] 等

现有规则参考（必须按照同样的风格生成）：
` + hardcodedExamples + `

关键：正则匹配的是磁链文本，不是日文标题。例如：
- 磁链为 "[...] 转学后... / Tenbin - 01 (CR..." → 正则中用 "Tenbin"
- 磁链为 "[桜都字幕组] ...Onaji Zemi... [01][1080P][简繁内封]" → \[桜都字幕组\].*Onaji Zemi.+\[(%02d)\]\[1080P\]\[简繁内封\].*
- 磁链为 "[Nix-Raws] ...Tenbin S01E01 [CR WEB-DL 1080p AVC AAC]" → \[Nix-Raws\].*Tenbin.+S01E(%02d).+CR
- 版本和来源标签都是硬匹配，不能模糊

只返回正则字符串。不要任何markdown、JSON、代码块、解释。`

// 从历史库里整理出来的正则样例，作为 AI 生成正则时的固定参考。
const hardcodedExamples = `1. ` + "`" + `\[黒ネズミたち\]  ?.*?Ryoumin 0-nin Start no Henkyou Ryoushu-sama.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
2. ` + "`" + `\[黒ネズミたち\]  ?.*?Kore Kaite Shine.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
3. ` + "`" + `\[Nix-Raws\]  ?.*?Neko to Ryuu.+S01E(%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?CR` + "`" + `
4. ` + "`" + `\[黒ネズミたち\]  ?.*?Tensei shitara Slime Datta Ken 4th Season.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
5. ` + "`" + `\[黒ネズミたち\]  ?.*?Kami no Shizuku.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
6. ` + "`" + `\[Nix-Raws\]  ?.*?Heroine Seijo Iie All Works Maid desu Ko.+S01E(%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?CR` + "`" + `
7. ` + "`" + `\[Nix-Raws\]  ?.*?LV999 no Murabito.+S01E(%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?CR` + "`" + `
8. ` + "`" + `\[黒ネズミたち\]  ?.*?Reiwa no Dara-san.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?CR` + "`" + `
9. ` + "`" + `\[黒ネズミたち\]  ?.*?Rakudai Kenja no Gakuin Musou.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?CR` + "`" + `
10. ` + "`" + `\[黒ネズミたち\]  ?.*?World Is Dancing.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
11. ` + "`" + `\[黒ネズミたち\]  ?.*?Hidarikiki no Eren.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
12. ` + "`" + `\[ANi\]  ?.*?I Want to End This Love Game.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?` + "`" + `
13. ` + "`" + `\[黒ネズミたち\]  ?.*?Replica datte, Koi wo Suru.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
14. ` + "`" + `\[黒ネズミたち\]  ?.*?Liar Game.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?CR` + "`" + `
15. ` + "`" + `\[黒ネズミたち\]  ?.*?Mao - (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
16. ` + "`" + `\[黒ネズミたち\]  ?.*? Kuroneko to Majo no Kyoushitsu.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
17. ` + "`" + `\[黒ネズミたち\]  ?.*? Yozakura-san Chi no Daisakusen 2nd Season.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
18. ` + "`" + `\[黒ネズミたち\]  ?.*?Tsue to Tsurugi no Wistoria Season 2.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
19. ` + "`" + `\[黒ネズミたち\]  ?.*?Shunkashuutou Daikousha: Haru no Mai.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
20. ` + "`" + `\[黒ネズミたち\]  ?.*?Yowayowa Sensei.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
21. ` + "`" + `\[黒ネズミたち\]  ?.*?Honzuki no Gekokujou 4th Season.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?CR` + "`" + `
22. ` + "`" + `\[黒ネズミたち\]  ?.*?Reincarnation no Kaben.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
23. ` + "`" + `【今晚月色真美】 ?.*?Otaku ni Yasashii Gal wa Inai.+\[(%02d)\]` + "`" + `
24. ` + "`" + `\[黒ネズミたち\]  ?.*?Youzitsu 4th Season.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
25. ` + "`" + `\[黒ネズミたち\]  ?.*?Rokkotsu.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
26. ` + "`" + `\[黒ネズミたち\]  ?.*?Lastame Season 2.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
27. ` + "`" + `\[黒ネズミたち\]  ?.*?Hyakki Yakoushou.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
28. ` + "`" + `\[Erai-raws\] Nippon Sangoku - (%02d) \[1080p AMZN WEBRip HEVC EAC3\]\[MultiSub\]` + "`" + `
29. ` + "`" + `\[黒ネズミたち\]  ?.*?Kami no Niwatsuki Kusunoki-tei.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
30. ` + "`" + `\[黒ネズミたち\]  ?.*?Yomi no Tsugai.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
31. ` + "`" + `\[黒ネズミたち\]  ?.*?Matakoro.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
32. ` + "`" + `\[黒ネズミたち\]  ?.*?Re:Zero 4th Season.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?Baha` + "`" + `
33. ` + "`" + `Witch Hat Atelier S01E(%02d) 1080p NF WEB-DL AAC2\.0 H 264-VARYG \(Tongari Boushi no Atelier, Multi-Subs\)` + "`" + `
34. ` + "`" + `\[ANi\]  ?.*?Takopis Original Sin.+- (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?` + "`" + `
35. ` + "`" + `\[LoliHouse\]  ?.*?Fate/strange Fake.+- (%02d)(?:\(Whispers of Dawn\))?(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))?.+?简繁内封字幕` + "`" + `
36. ` + "`" + `\[Skymoon-Raws\] Enter the Garden - (%02d)(?:.|&amp;)?(\d{2})?(?: (?:完|END|End|Fin|FIN))?(?: ?(?:(?:v|V)\d))? \[ViuTV` + "`"
