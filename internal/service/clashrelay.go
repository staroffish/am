package service

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode"

	"gopkg.in/yaml.v3"

	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/store"
)

const (
	// ClashRelayAirportGroup 是"装着机场所有节点"的 proxy-groups 固定名称，
	// 只被 VPS 节点的 dialer-proxy 引用，不被规则引用。
	ClashRelayAirportGroup = "机场节点"

	clashRelayDefaultUA     = "clash.meta"
	clashRelayMaxBody       = 8 << 20
	clashRelayDefaultKeep   = 10
	clashRelayDefaultNode   = "VPS-SS"
	clashRelayDefaultGroup  = "PROXY"
	clashRelayFetchTimeout  = 60 * time.Second
	clashRelayVersionPull   = "pull"
	clashRelayVersionManual = "manual"
)

// ValidationError 表示用户可修正的问题（配置不完整、YAML 非法、校验不通过）。
// handler 会把它映射成 400，而不是 500。
type ValidationError struct {
	Items []string
}

func (e *ValidationError) Error() string { return strings.Join(e.Items, "；") }

// ClashRelayService 负责拉取机场配置、改造成链式配置并管理版本。
type ClashRelayService struct {
	db     *store.DB
	log    *log.Logger
	mu     sync.Mutex
	client *http.Client
}

func NewClashRelayService(db *store.DB, logger *log.Logger) *ClashRelayService {
	return &ClashRelayService{
		db:  db,
		log: logger,
		client: &http.Client{
			Timeout: clashRelayFetchTimeout,
		},
	}
}

// DefaultClashRelaySettings 返回一份带默认值的空配置。
func DefaultClashRelaySettings() *model.ClashRelaySettings {
	return &model.ClashRelaySettings{
		AirportUA:    clashRelayDefaultUA,
		VPSNodeName:  clashRelayDefaultNode,
		VPSGroupName: clashRelayDefaultGroup,
		KeepVersions: clashRelayDefaultKeep,
	}
}

func normalizeClashRelaySettings(s *model.ClashRelaySettings) {
	s.AirportURL = strings.TrimSpace(s.AirportURL)
	s.AirportUA = strings.TrimSpace(s.AirportUA)
	s.VPSHost = strings.TrimSpace(s.VPSHost)
	s.VPSCipher = strings.TrimSpace(s.VPSCipher)
	s.VPSPassword = strings.TrimSpace(s.VPSPassword)
	s.VPSNodeName = strings.TrimSpace(s.VPSNodeName)
	s.VPSGroupName = strings.TrimSpace(s.VPSGroupName)
	s.RuleProvidersText = strings.TrimSpace(s.RuleProvidersText)

	if s.AirportUA == "" {
		s.AirportUA = clashRelayDefaultUA
	}
	if s.VPSNodeName == "" {
		s.VPSNodeName = clashRelayDefaultNode
	}
	if s.VPSGroupName == "" {
		s.VPSGroupName = clashRelayDefaultGroup
	}
	if s.KeepVersions < 1 {
		s.KeepVersions = clashRelayDefaultKeep
	}
}

// Settings 读取配置；没有记录时返回默认值（不落库）。
func (s *ClashRelayService) Settings(ctx context.Context) (*model.ClashRelaySettings, error) {
	st, err := s.db.GetClashRelaySettings(ctx)
	if err != nil {
		return nil, err
	}
	if st == nil {
		st = DefaultClashRelaySettings()
	}
	normalizeClashRelaySettings(st)
	return st, nil
}

// SaveSettings 合并并保存配置，保留"上次拉取"字段。
func (s *ClashRelayService) SaveSettings(ctx context.Context, in *model.ClashRelaySettings) (*model.ClashRelaySettings, error) {
	cur, err := s.db.GetClashRelaySettings(ctx)
	if err != nil {
		return nil, err
	}
	if cur == nil {
		cur = DefaultClashRelaySettings()
	}

	cur.AirportURL = in.AirportURL
	cur.AirportUA = in.AirportUA
	cur.VPSHost = in.VPSHost
	cur.VPSPort = in.VPSPort
	cur.VPSCipher = in.VPSCipher
	cur.VPSPassword = in.VPSPassword
	cur.VPSNodeName = in.VPSNodeName
	cur.VPSGroupName = in.VPSGroupName
	cur.KeepVersions = in.KeepVersions
	cur.RulesText = in.RulesText
	cur.RuleProvidersEnabled = in.RuleProvidersEnabled
	cur.RuleProvidersText = in.RuleProvidersText
	normalizeClashRelaySettings(cur)

	// 保存时只校验"已填写值"的格式，允许分步保存（先存订阅源，再存 VPS / rules）。
	if errs := validateClashRelaySettingsFormat(cur); len(errs) > 0 {
		return nil, &ValidationError{Items: errs}
	}
	if err := s.db.SaveClashRelaySettings(ctx, cur); err != nil {
		return nil, err
	}
	return cur, nil
}

// validateClashRelaySettingsFormat 只校验已填写值的格式，空值一律放行。
// 这样「订阅源」和「转换设置」可以分别保存，不会互相阻塞。
func validateClashRelaySettingsFormat(st *model.ClashRelaySettings) []string {
	var errs []string

	if st.AirportURL != "" {
		if u, err := url.Parse(st.AirportURL); err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
			errs = append(errs, "机场订阅 URL 必须是合法的 http(s) 地址")
		}
	}
	if st.VPSPort != 0 && (st.VPSPort < 1 || st.VPSPort > 65535) {
		errs = append(errs, "VPS 端口必须在 1-65535 之间")
	}
	if st.VPSNodeName != "" && st.VPSNodeName == ClashRelayAirportGroup {
		errs = append(errs, "VPS 节点名不能与固定组名「"+ClashRelayAirportGroup+"」相同")
	}
	if st.VPSGroupName != "" && st.VPSGroupName == ClashRelayAirportGroup {
		errs = append(errs, "VPS 组名不能与固定组名「"+ClashRelayAirportGroup+"」相同")
	}
	if st.VPSNodeName != "" && st.VPSNodeName == st.VPSGroupName {
		errs = append(errs, "VPS 节点名不能与组名相同")
	}
	if strings.TrimSpace(st.RulesText) != "" {
		if _, err := ParseRules(st.RulesText); err != nil {
			errs = append(errs, err.Error())
		}
	}
	if strings.TrimSpace(st.RuleProvidersText) != "" {
		if _, _, err := ParseRuleSets(st.RuleProvidersText); err != nil {
			errs = append(errs, err.Error())
		}
	}
	return errs
}

// validateClashRelaySettingsComplete 在「拉取并生成」前校验完整配置。
func validateClashRelaySettingsComplete(st *model.ClashRelaySettings) []string {
	errs := validateClashRelaySettingsFormat(st)

	if st.AirportURL == "" {
		errs = append(errs, "机场订阅 URL 不能为空")
	}
	if st.VPSHost == "" {
		errs = append(errs, "VPS 地址不能为空")
	}
	if st.VPSPort == 0 {
		errs = append(errs, "VPS 端口必须在 1-65535 之间")
	}
	if st.VPSCipher == "" {
		errs = append(errs, "SS 加密方式不能为空")
	}
	if st.VPSPassword == "" {
		errs = append(errs, "SS 密码不能为空")
	}
	if strings.TrimSpace(st.RulesText) == "" {
		errs = append(errs, "rules 不能为空")
	}
	if st.RuleProvidersEnabled && strings.TrimSpace(st.RuleProvidersText) == "" {
		errs = append(errs, "已启用自定义 rule-providers，但内容为空")
	}
	return errs
}

// PullDraft 是拉取后尚未入库的草稿。
type PullDraft struct {
	Content     string `json:"content"`
	AirportHash string `json:"airport_hash"`
	NodeCount   int    `json:"node_count"`
}

// PullResult 是一次手动拉取的结果（不落库）。
type PullResult struct {
	Draft     *PullDraft `json:"draft"`
	Unchanged bool       `json:"unchanged"`
	Message   string     `json:"message"`
}

// Pull 拉取机场配置、改造、校验并生成版本。
func (s *ClashRelayService) Pull(ctx context.Context) (*PullResult, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	st, err := s.Settings(ctx)
	if err != nil {
		return nil, err
	}
	if errs := validateClashRelaySettingsComplete(st); len(errs) > 0 {
		return nil, &ValidationError{Items: errs}
	}

	at := time.Now()
	raw, err := s.fetch(ctx, st.AirportURL, st.AirportUA)
	if err != nil {
		s.markPull(ctx, at, "error", err.Error())
		return nil, err
	}

	cfg, err := TransformAirportConfig(raw, st)
	if err != nil {
		s.markPull(ctx, at, "error", err.Error())
		return nil, err
	}
	content, err := MarshalConfigWithRules(cfg, st.RulesText)
	if err != nil {
		s.markPull(ctx, at, "error", err.Error())
		return nil, err
	}
	if errs := ValidateConfig(cfg, st); len(errs) > 0 {
		ve := &ValidationError{Items: errs}
		s.markPull(ctx, at, "error", ve.Error())
		return nil, ve
	}

	draft := &PullDraft{
		Content:     content,
		AirportHash: hashBytes(raw),
		NodeCount:   countAirportNodes(cfg),
	}

	cur, err := s.db.CurrentClashRelayVersion(ctx)
	if err != nil {
		s.markPull(ctx, at, "error", err.Error())
		return nil, err
	}
	unchanged := cur != nil && cur.ContentHash == hashString(content)

	msg := "已生成草稿，点击「保存为新版本」后才会入库"
	status := "draft"
	if unchanged {
		msg = "草稿与当前版本内容一致，保存不会产生新版本"
		status = "unchanged"
	}
	s.markPull(ctx, at, status, "")

	return &PullResult{Draft: draft, Unchanged: unchanged, Message: msg}, nil
}

// SaveVersion 保存整份配置，生成一个新版本。
// source 为 pull 或 manual；手动编辑已有版本时 airportHash 会从父版本继承。
func (s *ClashRelayService) SaveVersion(ctx context.Context, content string, parentID *uint, source, airportHash string) (*model.ClashRelayVersion, bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if source == "" {
		source = clashRelayVersionManual
	}
	if source != clashRelayVersionPull && source != clashRelayVersionManual {
		return nil, false, &ValidationError{Items: []string{"非法的版本来源"}}
	}

	st, err := s.Settings(ctx)
	if err != nil {
		return nil, false, err
	}

	cfg, err := ParseConfig(content)
	if err != nil {
		return nil, false, &ValidationError{Items: []string{"YAML 解析失败: " + err.Error()}}
	}
	if errs := ValidateConfig(cfg, st); len(errs) > 0 {
		return nil, false, &ValidationError{Items: errs}
	}

	if parentID != nil {
		if parent, err := s.db.GetClashRelayVersion(ctx, *parentID); err == nil {
			airportHash = parent.AirportHash
		}
	}

	version, unchanged, err := s.saveVersion(ctx, source, parentID, airportHash, content,
		countAirportNodes(cfg), st.KeepVersions)
	if err != nil {
		return nil, false, err
	}
	return version, unchanged, nil
}

func (s *ClashRelayService) saveVersion(ctx context.Context, source string, parentID *uint, airportHash, content string, nodeCount, keep int) (*model.ClashRelayVersion, bool, error) {
	contentHash := hashString(content)

	cur, err := s.db.CurrentClashRelayVersion(ctx)
	if err != nil {
		return nil, false, err
	}
	if cur != nil && cur.ContentHash == contentHash {
		return cur, true, nil
	}

	v := &model.ClashRelayVersion{
		Source:      source,
		ParentID:    parentID,
		AirportHash: airportHash,
		Content:     content,
		ContentHash: contentHash,
		NodeCount:   nodeCount,
	}
	if err := s.db.CreateClashRelayVersion(ctx, v, keep); err != nil {
		return nil, false, err
	}
	return v, false, nil
}

// Versions 返回版本列表（不含正文，避免列表接口过大）。
func (s *ClashRelayService) Versions(ctx context.Context) ([]model.ClashRelayVersion, error) {
	list, err := s.db.ListClashRelayVersions(ctx)
	if err != nil {
		return nil, err
	}
	for i := range list {
		list[i].Content = ""
	}
	return list, nil
}

// Version 返回指定版本的完整内容。
func (s *ClashRelayService) Version(ctx context.Context, id uint) (*model.ClashRelayVersion, error) {
	return s.db.GetClashRelayVersion(ctx, id)
}

// Activate 把指定版本设为当前。
func (s *ClashRelayService) Activate(ctx context.Context, id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.db.ActivateClashRelayVersion(ctx, id)
}

// CurrentContent 返回当前版本的完整 YAML；没有当前版本时返回空串。
func (s *ClashRelayService) CurrentContent(ctx context.Context) (string, error) {
	cur, err := s.db.CurrentClashRelayVersion(ctx)
	if err != nil {
		return "", err
	}
	if cur == nil {
		return "", nil
	}
	return cur.Content, nil
}

// StatusResult 是顶部状态卡需要的数据。
type StatusResult struct {
	Settings     *model.ClashRelaySettings `json:"settings"`
	Current      *model.ClashRelayVersion  `json:"current"`
	VersionCount int64                     `json:"version_count"`
}

func (s *ClashRelayService) Status(ctx context.Context) (*StatusResult, error) {
	st, err := s.Settings(ctx)
	if err != nil {
		return nil, err
	}
	cur, err := s.db.CurrentClashRelayVersion(ctx)
	if err != nil {
		return nil, err
	}
	if cur != nil {
		cur.Content = ""
	}
	n, err := s.db.CountClashRelayVersions(ctx)
	if err != nil {
		return nil, err
	}
	return &StatusResult{Settings: st, Current: cur, VersionCount: n}, nil
}

func (s *ClashRelayService) markPull(ctx context.Context, at time.Time, status, msg string) {
	if err := s.db.UpdateClashRelayPullStatus(ctx, at, status, msg); err != nil && s.log != nil {
		s.log.Printf("clash relay: update pull status: %v", err)
	}
}

func (s *ClashRelayService) fetch(ctx context.Context, rawURL, ua string) ([]byte, error) {
	if ua == "" {
		ua = clashRelayDefaultUA
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, rawURL, nil)
	if err != nil {
		return nil, fmt.Errorf("构造订阅请求失败: %w", err)
	}
	req.Header.Set("User-Agent", ua)
	req.Header.Set("Accept", "*/*")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("拉取机场订阅失败: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, clashRelayMaxBody+1))
	if err != nil {
		return nil, fmt.Errorf("读取机场响应失败: %w", err)
	}
	if len(body) > clashRelayMaxBody {
		return nil, errors.New("机场订阅内容超过 8MB，已拒绝")
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("机场返回 HTTP %d: %s", resp.StatusCode, snippet(body))
	}
	if len(body) == 0 {
		return nil, errors.New("机场返回空内容")
	}
	return body, nil
}

func snippet(b []byte) string {
	s := strings.TrimSpace(string(b))
	if len(s) > 200 {
		return s[:200] + "…"
	}
	return s
}

// ---------------- 配置解析与改造 ----------------

// ParseConfig 把整份 YAML 解析成通用 map。
func ParseConfig(content string) (map[string]any, error) {
	var cfg map[string]any
	if err := yaml.Unmarshal([]byte(content), &cfg); err != nil {
		return nil, err
	}
	if cfg == nil {
		return nil, errors.New("配置为空")
	}
	return cfg, nil
}

// MarshalConfig 把通用 map 序列化为 YAML。
func MarshalConfig(cfg map[string]any) (string, error) {
	return marshalConfigMap(cfg)
}

// marshalConfigMap 序列化时将可打印的非 ASCII 字符（如 emoji）临时替换成
// 唯一 token，序列化后再还原。这样 yaml.v3 不会把它们输出成 "\U0001F1EE"
// 之类的转义，同时也不会误伤字面量反斜杠。
func marshalConfigMap(cfg map[string]any) (string, error) {
	tokens := make(map[string]string)
	counter := 0
	nonce := newTokenNonce()

	encodable := encodePrintableNonASCII(cfg, tokens, &counter, nonce)
	b, err := yaml.Marshal(encodable)
	if err != nil {
		return "", err
	}

	out := string(b)
	for tok, r := range tokens {
		out = strings.ReplaceAll(out, tok, r)
	}
	return out, nil
}

func newTokenNonce() string {
	var b [6]byte
	if _, err := rand.Read(b[:]); err != nil {
		return fmt.Sprintf("%x", time.Now().UnixNano())
	}
	return hex.EncodeToString(b[:])
}

func encodePrintableNonASCII(v any, tokens map[string]string, counter *int, nonce string) any {
	switch x := v.(type) {
	case string:
		return encodeStringNonASCII(x, tokens, counter, nonce)
	case map[string]any:
		m := make(map[string]any, len(x))
		for k, val := range x {
			m[k] = encodePrintableNonASCII(val, tokens, counter, nonce)
		}
		return m
	case []any:
		s := make([]any, len(x))
		for i, val := range x {
			s[i] = encodePrintableNonASCII(val, tokens, counter, nonce)
		}
		return s
	default:
		return v
	}
}

func encodeStringNonASCII(s string, tokens map[string]string, counter *int, nonce string) string {
	needs := false
	for _, r := range s {
		if r > 0x7F && unicode.IsPrint(r) {
			needs = true
			break
		}
	}
	if !needs {
		return s
	}

	var b strings.Builder
	for _, r := range s {
		if r > 0x7F && unicode.IsPrint(r) {
			*counter++
			tok := fmt.Sprintf("__yamlu_%s_%d__", nonce, *counter)
			tokens[tok] = string(r)
			b.WriteString(tok)
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

// MarshalConfigWithRules 序列化配置，并保留用户 rules 原文（含注释）。
// 做法：先把 rules 从 map 里摘掉序列化其余部分，再把原文作为 rules 块追加到末尾。
func MarshalConfigWithRules(cfg map[string]any, rulesText string) (string, error) {
	rules, hasRules := cfg["rules"]
	if !hasRules {
		return MarshalConfig(cfg)
	}

	delete(cfg, "rules")
	body, err := marshalConfigMap(cfg)
	if err != nil {
		return "", err
	}
	cfg["rules"] = rules

	return body + "rules:\n" + normalizeRulesBlock(rulesText) + "\n", nil
}

// normalizeRulesBlock 把用户输入的 rules 文本整理成可嵌在 "rules:" 之下的块。
// 保留注释与空行；如果用户贴的是完整的 "rules:" 块，则去掉那行表头。
func normalizeRulesBlock(text string) string {
	lines := strings.Split(NormalizeRulesText(text), "\n")
	out := make([]string, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			out = append(out, "")
			continue
		}
		out = append(out, "  "+line)
	}
	return strings.Join(out, "\n")
}

// TransformAirportConfig 以机场配置为底，只覆盖 proxies / proxy-groups / rules 三个块。
func TransformAirportConfig(raw []byte, st *model.ClashRelaySettings) (map[string]any, error) {
	var cfg map[string]any
	if err := yaml.Unmarshal(raw, &cfg); err != nil {
		return nil, fmt.Errorf("机场配置不是合法 YAML: %w", err)
	}
	if cfg == nil {
		return nil, errors.New("机场配置为空")
	}

	proxiesRaw, ok := cfg["proxies"].([]any)
	if !ok || len(proxiesRaw) == 0 {
		return nil, errors.New("机场配置缺少 proxies 或 proxies 为空")
	}

	names := make([]string, 0, len(proxiesRaw))
	existing := make(map[string]bool, len(proxiesRaw))
	for i, p := range proxiesRaw {
		m, ok := p.(map[string]any)
		if !ok {
			return nil, fmt.Errorf("proxies 第 %d 项不是对象", i+1)
		}
		name, _ := m["name"].(string)
		if name == "" {
			return nil, fmt.Errorf("proxies 第 %d 项缺少 name", i+1)
		}
		if existing[name] {
			return nil, fmt.Errorf("机场节点名重复: %s", name)
		}
		existing[name] = true
		names = append(names, name)
	}

	vpsName := uniqueName(strings.TrimSpace(st.VPSNodeName), existing)
	vpsProxy := map[string]any{
		"name":         vpsName,
		"type":         "ss",
		"server":       st.VPSHost,
		"port":         st.VPSPort,
		"cipher":       st.VPSCipher,
		"password":     st.VPSPassword,
		"dialer-proxy": ClashRelayAirportGroup,
	}

	proxies := make([]any, 0, len(proxiesRaw)+1)
	proxies = append(proxies, proxiesRaw...)
	proxies = append(proxies, vpsProxy)

	rules, err := ParseRules(st.RulesText)
	if err != nil {
		return nil, err
	}
	if st.RuleProvidersEnabled {
		anchors, providers, err := ParseRuleSets(st.RuleProvidersText)
		if err != nil {
			return nil, err
		}
		cfg["rule-providers"] = providers
		if anchors != nil {
			cfg["rule-anchor"] = anchors
		} else {
			// 启用自定义 providers 时，机场原有的 rule-anchor 不再需要。
			delete(cfg, "rule-anchor")
		}
	}

	members := make([]any, 0, len(names))
	for _, n := range names {
		members = append(members, n)
	}
	airportGroup := map[string]any{
		"name":    ClashRelayAirportGroup,
		"type":    "url-test",
		"proxies": members,
	}
	vpsGroup := map[string]any{
		"name":    st.VPSGroupName,
		"type":    "select",
		"proxies": []any{vpsName},
	}

	cfg["proxies"] = proxies
	cfg["proxy-groups"] = []any{airportGroup, vpsGroup}
	cfg["rules"] = rules
	return cfg, nil
}

// ParseRules 把页面上的 rules 文本解析成 YAML 列表。
// 支持两种写法：直接的 "- 规则" 列表，或完整的 "rules:" 块。
// 解析前会统一缩进，因此不要求用户手工对齐（从别处粘贴时常出现首行顶格、其余行带缩进）。
func ParseRules(text string) ([]any, error) {
	t := strings.TrimSpace(text)
	if t == "" {
		return nil, errors.New("rules 不能为空")
	}

	normalized := NormalizeRulesText(t)

	var rules []any
	if err := yaml.Unmarshal([]byte(normalized), &rules); err == nil && len(rules) > 0 {
		return rules, nil
	}

	var wrapped struct {
		Rules []any `yaml:"rules"`
	}
	if err := yaml.Unmarshal([]byte("rules:\n"+normalized+"\n"), &wrapped); err == nil && len(wrapped.Rules) > 0 {
		return wrapped.Rules, nil
	}

	return nil, errors.New("rules 不是合法的 YAML 列表（每行形如 - DOMAIN-SUFFIX,example.com,PROXY）")
}

// ParseRuleSets 解析页面上的 rule-anchor / rule-providers 文本。
// 返回 (anchors, providers)。anchors 可以为空；providers 必须非空。
// 支持完整块（同时包含 rule-anchor 与 rule-providers），也兼容只写 providers 映射。
func ParseRuleSets(text string) (map[string]any, map[string]any, error) {
	t := strings.TrimSpace(text)
	if t == "" {
		return nil, nil, errors.New("rule-providers 不能为空")
	}

	var m map[string]any
	if err := yaml.Unmarshal([]byte(t), &m); err != nil {
		return nil, nil, fmt.Errorf("rule-anchor / rule-providers 不是合法的 YAML: %w", err)
	}
	if len(m) == 0 {
		return nil, nil, errors.New("rule-providers 不能为空")
	}

	var anchors map[string]any
	if v, ok := m["rule-anchor"]; ok {
		a, ok := v.(map[string]any)
		if !ok {
			return nil, nil, errors.New("rule-anchor 必须是映射")
		}
		if len(a) > 0 {
			anchors = a
		}
		delete(m, "rule-anchor")
	}

	if v, ok := m["rule-providers"]; ok {
		if p, ok := v.(map[string]any); ok {
			if len(p) == 0 {
				return nil, nil, errors.New("rule-providers 不能为空")
			}
			return anchors, p, nil
		}
		// "rule-providers:" 为空块、用户把内容顶格粘贴：去掉空键，其余当作 providers。
		delete(m, "rule-providers")
		if len(m) == 0 {
			return nil, nil, errors.New("rule-providers 不能为空")
		}
		return anchors, m, nil
	}

	if len(m) == 0 {
		return nil, nil, errors.New("rule-providers 不能为空")
	}
	return anchors, m, nil
}

// DeleteVersion 删除指定版本（当前版本不允许删除）。
func (s *ClashRelayService) DeleteVersion(ctx context.Context, id uint) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.db.DeleteClashRelayVersion(ctx, id)
}

// NormalizeRulesText 统一 rules 文本的缩进，并去掉可能存在的 "rules:" 表头。
// Clash 规则都是单行条目，所以按行去除首尾空白是安全的；注释与空行会保留。
func NormalizeRulesText(text string) string {
	lines := strings.Split(strings.TrimSpace(text), "\n")
	out := make([]string, 0, len(lines))
	headerRemoved := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if !headerRemoved && strings.HasPrefix(trimmed, "rules:") {
			headerRemoved = true
			continue
		}
		if trimmed == "" {
			out = append(out, "")
			continue
		}
		out = append(out, trimmed)
	}
	return strings.Join(out, "\n")
}

// ValidateConfig 执行发布前校验，返回人类可读的问题列表。
func ValidateConfig(cfg map[string]any, st *model.ClashRelaySettings) []string {
	var errs []string

	proxiesRaw, ok := cfg["proxies"].([]any)
	if !ok || len(proxiesRaw) == 0 {
		return append(errs, "缺少 proxies 或 proxies 为空")
	}

	nodeNames := make(map[string]bool, len(proxiesRaw))
	vpsName := ""
	for i, p := range proxiesRaw {
		m, ok := p.(map[string]any)
		if !ok {
			errs = append(errs, fmt.Sprintf("proxies 第 %d 项不是对象", i+1))
			continue
		}
		name, _ := m["name"].(string)
		if name == "" {
			errs = append(errs, fmt.Sprintf("proxies 第 %d 项缺少 name", i+1))
			continue
		}
		if nodeNames[name] {
			errs = append(errs, "节点名重复: "+name)
		}
		nodeNames[name] = true

		typ, _ := m["type"].(string)
		dialer, _ := m["dialer-proxy"].(string)
		if typ == "ss" && dialer == ClashRelayAirportGroup {
			if vpsName != "" {
				errs = append(errs, "存在多个经「"+ClashRelayAirportGroup+"」拨号的 SS 出口节点")
			}
			vpsName = name
		}
	}
	if vpsName == "" {
		errs = append(errs, "缺少经「"+ClashRelayAirportGroup+"」拨号的 SS 出口节点")
	}

	groupsRaw, ok := cfg["proxy-groups"].([]any)
	if !ok || len(groupsRaw) == 0 {
		return append(errs, "缺少 proxy-groups 或 proxy-groups 为空")
	}

	groupNames := make(map[string]bool, len(groupsRaw))
	var airportGroup, vpsGroup map[string]any
	for i, g := range groupsRaw {
		m, ok := g.(map[string]any)
		if !ok {
			errs = append(errs, fmt.Sprintf("proxy-groups 第 %d 项不是对象", i+1))
			continue
		}
		name, _ := m["name"].(string)
		if name == "" {
			errs = append(errs, fmt.Sprintf("proxy-groups 第 %d 项缺少 name", i+1))
			continue
		}
		if groupNames[name] {
			errs = append(errs, "组名重复: "+name)
		}
		if nodeNames[name] {
			errs = append(errs, "组名与节点名冲突: "+name)
		}
		groupNames[name] = true

		switch name {
		case ClashRelayAirportGroup:
			airportGroup = m
		default:
			if name == st.VPSGroupName {
				vpsGroup = m
			}
		}
	}

	if airportGroup == nil {
		errs = append(errs, "缺少「"+ClashRelayAirportGroup+"」组")
	} else {
		if typ, _ := airportGroup["type"].(string); typ != "url-test" {
			errs = append(errs, "「"+ClashRelayAirportGroup+"」组的 type 必须是 url-test")
		}
		members := toStrings(airportGroup["proxies"])
		if len(members) == 0 {
			errs = append(errs, "「"+ClashRelayAirportGroup+"」组为空")
		}
		for _, m := range members {
			if !nodeNames[m] {
				errs = append(errs, "「"+ClashRelayAirportGroup+"」组成员不存在: "+m)
			}
			if m == vpsName {
				errs = append(errs, "「"+ClashRelayAirportGroup+"」组不能包含 VPS 节点（会成环）")
			}
		}
	}

	if vpsGroup == nil {
		errs = append(errs, "缺少「"+st.VPSGroupName+"」组")
	} else {
		if typ, _ := vpsGroup["type"].(string); typ != "select" {
			errs = append(errs, "「"+st.VPSGroupName+"」组的 type 必须是 select")
		}
		members := toStrings(vpsGroup["proxies"])
		if len(members) != 1 || members[0] != vpsName {
			errs = append(errs, "「"+st.VPSGroupName+"」组必须只包含 VPS 节点")
		}
	}

	rulesRaw, rulesOK := cfg["rules"].([]any)
	if !rulesOK || len(rulesRaw) == 0 {
		errs = append(errs, "缺少 rules 或 rules 为空")
	} else {
		providers, _ := cfg["rule-providers"].(map[string]any)
		missing := make(map[string]bool)
		for _, r := range rulesRaw {
			rule, ok := r.(string)
			if !ok || !strings.HasPrefix(rule, "RULE-SET,") {
				continue
			}
			parts := strings.Split(rule, ",")
			if len(parts) < 2 {
				continue
			}
			name := strings.TrimSpace(parts[1])
			if name == "" {
				continue
			}
			if _, ok := providers[name]; !ok {
				missing[name] = true
			}
		}
		names := make([]string, 0, len(missing))
		for n := range missing {
			names = append(names, n)
		}
		sort.Strings(names)
		for _, n := range names {
			errs = append(errs, "rules 引用了不存在的 rule-provider: "+n)
		}
	}

	return errs
}

func toStrings(v any) []string {
	switch vv := v.(type) {
	case []any:
		out := make([]string, 0, len(vv))
		for _, item := range vv {
			if s, ok := item.(string); ok {
				out = append(out, s)
			}
		}
		return out
	case []string:
		return vv
	default:
		return nil
	}
}

func countAirportNodes(cfg map[string]any) int {
	proxies, _ := cfg["proxies"].([]any)
	n := 0
	for _, p := range proxies {
		m, ok := p.(map[string]any)
		if !ok {
			continue
		}
		typ, _ := m["type"].(string)
		dialer, _ := m["dialer-proxy"].(string)
		if typ == "ss" && dialer == ClashRelayAirportGroup {
			continue
		}
		n++
	}
	return n
}

func uniqueName(base string, existing map[string]bool) string {
	if base == "" {
		base = clashRelayDefaultNode
	}
	if !existing[base] {
		return base
	}
	for i := 2; ; i++ {
		candidate := fmt.Sprintf("%s-%d", base, i)
		if !existing[candidate] {
			return candidate
		}
	}
}

func hashBytes(b []byte) string {
	sum := sha256.Sum256(b)
	return hex.EncodeToString(sum[:])
}

func hashString(s string) string {
	return hashBytes([]byte(s))
}
