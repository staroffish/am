package service

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"

	"github.com/staroffish/am/internal/config"
	"github.com/staroffish/am/internal/model"
	"github.com/staroffish/am/internal/store"
)

const sampleAirportYAML = `mixed-port: 7890
allow-lan: false
mode: rule
log-level: info
external-controller: 127.0.0.1:9090
secret: "airport-secret"
dns:
  enable: true
  nameserver:
    - 1.1.1.1
tun:
  enable: false
proxies:
  - name: 香港01
    type: ss
    server: hk1.example.com
    port: 443
    cipher: aes-256-gcm
    password: pw1
  - name: 日本02
    type: ss
    server: jp2.example.com
    port: 443
    cipher: aes-256-gcm
    password: pw2
proxy-groups:
  - name: 机场自带组
    type: select
    proxies:
      - 香港01
rules:
  - MATCH,机场自带组
rule-providers:
  foo:
    type: http
    behavior: domain
    url: https://example.com/foo.yaml
    path: ./rules/foo.yaml
    interval: 86400
`

func testSettings() *model.ClashRelaySettings {
	return &model.ClashRelaySettings{
		AirportURL:   "https://airport.example.com/sub?token=x",
		AirportUA:    "clash.meta",
		VPSHost:      "vps.example.com",
		VPSPort:      8443,
		VPSCipher:    "aes-256-gcm",
		VPSPassword:  "vps-pass",
		VPSNodeName:  "VPS-SS",
		VPSGroupName: "PROXY",
		KeepVersions: 10,
		RulesText: `- DOMAIN-SUFFIX,google.com,PROXY
- GEOIP,CN,DIRECT
- MATCH,PROXY`,
	}
}

func TestTransformAirportConfig(t *testing.T) {
	st := testSettings()
	cfg, err := TransformAirportConfig([]byte(sampleAirportYAML), st)
	if err != nil {
		t.Fatalf("transform: %v", err)
	}
	if errs := ValidateConfig(cfg, st); len(errs) > 0 {
		t.Fatalf("validate failed: %v", errs)
	}

	// 未覆盖的字段必须原样保留
	for _, key := range []string{"mixed-port", "dns", "tun", "external-controller", "secret", "rule-providers"} {
		if _, ok := cfg[key]; !ok {
			t.Errorf("字段 %q 应被保留", key)
		}
	}
	if cfg["mixed-port"] != 7890 {
		t.Errorf("mixed-port = %v, want 7890", cfg["mixed-port"])
	}

	proxies, _ := cfg["proxies"].([]any)
	if len(proxies) != 3 {
		t.Fatalf("proxies 数量 = %d, want 3", len(proxies))
	}
	vps, _ := proxies[2].(map[string]any)
	if vps["name"] != "VPS-SS" || vps["type"] != "ss" {
		t.Errorf("VPS 节点异常: %#v", vps)
	}
	if vps["dialer-proxy"] != ClashRelayAirportGroup {
		t.Errorf("dialer-proxy = %v, want %s", vps["dialer-proxy"], ClashRelayAirportGroup)
	}
	if vps["server"] != "vps.example.com" || vps["port"] != 8443 {
		t.Errorf("VPS 连接参数异常: %#v", vps)
	}

	groups, _ := cfg["proxy-groups"].([]any)
	if len(groups) != 2 {
		t.Fatalf("proxy-groups 数量 = %d, want 2", len(groups))
	}
	airportGroup, _ := groups[0].(map[string]any)
	vpsGroup, _ := groups[1].(map[string]any)
	if airportGroup["name"] != ClashRelayAirportGroup || airportGroup["type"] != "url-test" {
		t.Errorf("机场组异常: %#v", airportGroup)
	}
	if got := toStrings(airportGroup["proxies"]); len(got) != 2 || got[0] != "香港01" || got[1] != "日本02" {
		t.Errorf("机场组成员 = %v", got)
	}
	if vpsGroup["name"] != "PROXY" || vpsGroup["type"] != "select" {
		t.Errorf("VPS 组异常: %#v", vpsGroup)
	}
	if got := toStrings(vpsGroup["proxies"]); len(got) != 1 || got[0] != "VPS-SS" {
		t.Errorf("VPS 组成员 = %v", got)
	}

	rules, _ := cfg["rules"].([]any)
	if len(rules) != 3 {
		t.Fatalf("rules 数量 = %d, want 3", len(rules))
	}
	if rules[0] != "DOMAIN-SUFFIX,google.com,PROXY" {
		t.Errorf("rules[0] = %v", rules[0])
	}
}

func TestTransformVPSNameCollision(t *testing.T) {
	yaml := strings.Replace(sampleAirportYAML, "name: 香港01", "name: VPS-SS", 1)
	st := testSettings()
	cfg, err := TransformAirportConfig([]byte(yaml), st)
	if err != nil {
		t.Fatalf("transform: %v", err)
	}
	if errs := ValidateConfig(cfg, st); len(errs) > 0 {
		t.Fatalf("validate failed: %v", errs)
	}
	proxies, _ := cfg["proxies"].([]any)
	vps, _ := proxies[len(proxies)-1].(map[string]any)
	if vps["name"] != "VPS-SS-2" {
		t.Errorf("同名冲突时应自动改名，实际 = %v", vps["name"])
	}
	groups, _ := cfg["proxy-groups"].([]any)
	vpsGroup, _ := groups[1].(map[string]any)
	if got := toStrings(vpsGroup["proxies"]); len(got) != 1 || got[0] != "VPS-SS-2" {
		t.Errorf("VPS 组应引用改名后的节点，实际 = %v", got)
	}
}

func TestTransformMissingProxies(t *testing.T) {
	_, err := TransformAirportConfig([]byte("rules:\n  - MATCH,DIRECT\n"), testSettings())
	if err == nil || !strings.Contains(err.Error(), "proxies") {
		t.Fatalf("缺少 proxies 时应报错，实际 err = %v", err)
	}
}

func TestParseRules(t *testing.T) {
	cases := []struct {
		name    string
		text    string
		wantLen int
		wantErr bool
	}{
		{"direct list", "- MATCH,PROXY\n- GEOIP,CN,DIRECT", 2, false},
		{"wrapped rules block", "rules:\n  - MATCH,PROXY\n  - GEOIP,CN,DIRECT", 2, false},
		{"empty", "   ", 0, true},
		{"garbage", "this is not a list", 0, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rules, err := ParseRules(tc.text)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got %v", rules)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(rules) != tc.wantLen {
				t.Fatalf("len = %d, want %d", len(rules), tc.wantLen)
			}
		})
	}
}

func TestValidateConfigErrors(t *testing.T) {
	st := testSettings()

	base, err := TransformAirportConfig([]byte(sampleAirportYAML), st)
	if err != nil {
		t.Fatalf("transform: %v", err)
	}

	t.Run("missing vps node", func(t *testing.T) {
		cfg := cloneMap(t, base)
		proxies, _ := cfg["proxies"].([]any)
		cfg["proxies"] = proxies[:len(proxies)-1]
		if errs := ValidateConfig(cfg, st); !containsSub(errs, "缺少经") {
			t.Fatalf("应提示缺少出口节点，实际 = %v", errs)
		}
	})

	t.Run("loop: vps in airport group", func(t *testing.T) {
		cfg := cloneMap(t, base)
		groups, _ := cfg["proxy-groups"].([]any)
		airportGroup, _ := groups[0].(map[string]any)
		airportGroup["proxies"] = []any{"香港01", "VPS-SS"}
		if errs := ValidateConfig(cfg, st); !containsSub(errs, "成环") {
			t.Fatalf("应提示成环，实际 = %v", errs)
		}
	})

	t.Run("vps group wrong member", func(t *testing.T) {
		cfg := cloneMap(t, base)
		groups, _ := cfg["proxy-groups"].([]any)
		vpsGroup, _ := groups[1].(map[string]any)
		vpsGroup["proxies"] = []any{"香港01"}
		if errs := ValidateConfig(cfg, st); !containsSub(errs, "只包含 VPS 节点") {
			t.Fatalf("应提示 VPS 组成员错误，实际 = %v", errs)
		}
	})

	t.Run("empty rules", func(t *testing.T) {
		cfg := cloneMap(t, base)
		cfg["rules"] = []any{}
		if errs := ValidateConfig(cfg, st); !containsSub(errs, "rules") {
			t.Fatalf("应提示 rules 为空，实际 = %v", errs)
		}
	})

	t.Run("duplicate node name", func(t *testing.T) {
		cfg := cloneMap(t, base)
		proxies, _ := cfg["proxies"].([]any)
		dup, _ := proxies[0].(map[string]any)
		copyProxy := map[string]any{"name": dup["name"], "type": "ss"}
		cfg["proxies"] = append(proxies, copyProxy)
		if errs := ValidateConfig(cfg, st); !containsSub(errs, "节点名重复") {
			t.Fatalf("应提示节点名重复，实际 = %v", errs)
		}
	})
}

func TestValidateSettingsFormat(t *testing.T) {
	// 允许分步保存：空配置在"格式校验"下必须通过。
	empty := &model.ClashRelaySettings{}
	if errs := validateClashRelaySettingsFormat(empty); len(errs) > 0 {
		t.Fatalf("空配置应通过格式校验（供分步保存），实际 = %v", errs)
	}

	// 只填订阅源也应通过。
	partial := &model.ClashRelaySettings{AirportURL: "https://airport.example.com/sub?token=x"}
	if errs := validateClashRelaySettingsFormat(partial); len(errs) > 0 {
		t.Fatalf("只填订阅源应通过格式校验，实际 = %v", errs)
	}

	bad := &model.ClashRelaySettings{AirportURL: "ftp://x", VPSPort: 70000, RulesText: "garbage"}
	errs := validateClashRelaySettingsFormat(bad)
	if !containsSub(errs, "http(s)") || !containsSub(errs, "端口") || !containsSub(errs, "rules") {
		t.Fatalf("格式校验应报告 URL/端口/rules，实际 = %v", errs)
	}

	// 完整校验仍然要求 VPS 与 rules。
	if errs := validateClashRelaySettingsComplete(partial); !containsSub(errs, "VPS 地址") || !containsSub(errs, "rules") {
		t.Fatalf("完整校验应提示 VPS/rules 缺失，实际 = %v", errs)
	}
}

func TestValidateSettings(t *testing.T) {
	st := testSettings()
	if errs := validateClashRelaySettingsComplete(st); len(errs) > 0 {
		t.Fatalf("默认配置应合法: %v", errs)
	}

	bad := testSettings()
	bad.VPSPort = 0
	bad.VPSGroupName = ClashRelayAirportGroup
	bad.RulesText = ""
	errs := validateClashRelaySettingsComplete(bad)
	if !containsSub(errs, "端口") || !containsSub(errs, "固定组名") || !containsSub(errs, "rules") {
		t.Fatalf("应报告端口/组名/rules 问题，实际 = %v", errs)
	}
}

func cloneMap(t *testing.T, in map[string]any) map[string]any {
	t.Helper()
	raw, err := MarshalConfig(in)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	out, err := ParseConfig(raw)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	return out
}

func containsSub(items []string, sub string) bool {
	for _, s := range items {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

func TestMarshalConfigWithRulesKeepsComments(t *testing.T) {
	st := testSettings()
	st.RulesText = `# 国内直连
- GEOIP,CN,DIRECT
# 其余走代理
- MATCH,PROXY`

	cfg, err := TransformAirportConfig([]byte(sampleAirportYAML), st)
	if err != nil {
		t.Fatalf("transform: %v", err)
	}
	content, err := MarshalConfigWithRules(cfg, st.RulesText)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	for _, comment := range []string{"# 国内直连", "# 其余走代理"} {
		if !strings.Contains(content, comment) {
			t.Errorf("生成内容应保留注释 %q，实际:\n%s", comment, content)
		}
	}
	parsed, err := ParseConfig(content)
	if err != nil {
		t.Fatalf("生成内容不是合法 YAML: %v\n%s", err, content)
	}
	if errs := ValidateConfig(parsed, st); len(errs) > 0 {
		t.Fatalf("生成内容校验失败: %v", errs)
	}
	rules, _ := parsed["rules"].([]any)
	if len(rules) != 2 {
		t.Fatalf("rules 数量 = %d, want 2", len(rules))
	}
	if rules[0] != "GEOIP,CN,DIRECT" || rules[1] != "MATCH,PROXY" {
		t.Errorf("rules 内容异常: %v", rules)
	}
}

func TestMarshalConfigWithRulesFullBlock(t *testing.T) {
	st := testSettings()
	st.RulesText = `rules:
  # 国内
  - GEOIP,CN,DIRECT
  - MATCH,PROXY`

	cfg, err := TransformAirportConfig([]byte(sampleAirportYAML), st)
	if err != nil {
		t.Fatalf("transform: %v", err)
	}
	content, err := MarshalConfigWithRules(cfg, st.RulesText)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	parsed, err := ParseConfig(content)
	if err != nil {
		t.Fatalf("生成内容不是合法 YAML: %v\n%s", err, content)
	}
	rules, _ := parsed["rules"].([]any)
	if len(rules) != 2 {
		t.Fatalf("rules 数量 = %d, want 2\n%s", len(rules), content)
	}
}

func TestParseRulesInconsistentIndent(t *testing.T) {
	// 用户从别处粘贴时的典型情况：首行顶格，其余行带两个空格。
	text := `- SRC-IP-CIDR,192.168.1.233/32,DIRECT
  - DOMAIN-KEYWORD,cursor,auto
  - DOMAIN-KEYWORD,docker.*,auto
  # 注释
  #- DOMAIN-KEYWORD,deepseek,auto
  - GEOIP,!CN,auto
  - IP-CIDR,192.168.0.0/16,DIRECT
  - MATCH,auto`

	var raw []any
	if err := yaml.Unmarshal([]byte(text), &raw); err == nil {
		t.Fatalf("原始文本缩进不一致，应该解析失败")
	}

	rules, err := ParseRules(text)
	if err != nil {
		t.Fatalf("统一缩进后应解析成功: %v", err)
	}
	if len(rules) != 6 {
		t.Fatalf("rules 数量 = %d, want 6", len(rules))
	}

	// 生成的 rules 块也必须是合法 YAML。
	block := normalizeRulesBlock(text)
	var parsed []any
	if err := yaml.Unmarshal([]byte(block), &parsed); err != nil {
		t.Fatalf("生成的 rules 块不是合法 YAML: %v\n%s", err, block)
	}
	if len(parsed) != 6 {
		t.Fatalf("生成块 rules 数量 = %d, want 6", len(parsed))
	}
}

func TestMarshalConfigUnescapesEmoji(t *testing.T) {
	cfg := map[string]any{
		"proxies": []any{
			map[string]any{"name": "🇮🇸 冰岛", "type": "ss"},
			map[string]any{"name": "🇯🇵 日本", "type": "ss"},
		},
		"literal": `\U0001F1EE`,
	}
	out, err := MarshalConfig(cfg)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !strings.Contains(out, "🇮🇸 冰岛") || !strings.Contains(out, "🇯🇵 日本") {
		t.Fatalf("生成内容应包含原始 emoji:\n%s", out)
	}

	parsed, err := ParseConfig(out)
	if err != nil {
		t.Fatalf("生成内容不是合法 YAML: %v\n%s", err, out)
	}
	proxies, _ := parsed["proxies"].([]any)
	name, _ := proxies[0].(map[string]any)["name"].(string)
	if name != "🇮🇸 冰岛" {
		t.Fatalf("emoji 名称被破坏: %q", name)
	}
	if parsed["literal"] != `\U0001F1EE` {
		t.Fatalf("字面量反斜杠被错误还原: %#v", parsed["literal"])
	}
}

func TestTransformKeepsEmojiInProxies(t *testing.T) {
	airport := strings.Replace(sampleAirportYAML, "name: 香港01", "name: 🇮🇸 冰岛", 1)
	st := testSettings()

	cfg, err := TransformAirportConfig([]byte(airport), st)
	if err != nil {
		t.Fatalf("transform: %v", err)
	}
	content, err := MarshalConfigWithRules(cfg, st.RulesText)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	if !strings.Contains(content, "🇮🇸 冰岛") {
		t.Fatalf("生成内容应直接展示 emoji，而不是 \\U 转义:\n%s", content)
	}
	if strings.Contains(content, `\U0001F1EE`) {
		t.Fatalf("生成内容仍包含 \\U 转义:\n%s", content)
	}
	if _, err := ParseConfig(content); err != nil {
		t.Fatalf("生成内容不是合法 YAML: %v", err)
	}
	if errs := ValidateConfig(cfg, st); len(errs) > 0 {
		t.Fatalf("校验失败: %v", errs)
	}
}

func TestParseRuleSets(t *testing.T) {
	inner := `proxy:
  type: http
  behavior: domain
  url: https://example.com/proxy.yaml
  path: ./proxy.yaml
  interval: 86400`
	if _, m, err := ParseRuleSets(inner); err != nil || len(m) != 1 {
		t.Fatalf("inner map: m=%v err=%v", m, err)
	}

	full := "rule-providers:\n" + inner
	if _, m, err := ParseRuleSets(full); err != nil || len(m) != 1 {
		t.Fatalf("full block: m=%v err=%v", m, err)
	}

	if _, _, err := ParseRuleSets("   "); err == nil {
		t.Fatalf("空内容应报错")
	}
	if _, _, err := ParseRuleSets(":::bad"); err == nil {
		t.Fatalf("非法 YAML 应报错")
	}
}

func TestParseRuleSetsWithAnchor(t *testing.T) {
	text := `rule-anchor:
  ip: &ip {type: http, interval: 86400, behavior: ipcidr, format: mrs}
  domain: &domain {type: http, interval: 86400, behavior: domain, format: mrs}
rule-providers:
  direct:
    type: http
    behavior: domain
    url: "https://example.com/direct.txt"
    path: ./ruleset/direct.yaml
    interval: 86400
  cn_domain:
    <<: *domain
    url: "https://example.com/cn_domain.mrs"
  cn_ip:
    <<: *ip
    url: "https://example.com/cn_ip.mrs"
`
	anchors, providers, err := ParseRuleSets(text)
	if err != nil {
		t.Fatalf("ParseRuleSets: %v", err)
	}
	if len(anchors) != 2 {
		t.Fatalf("anchors 数量 = %d, want 2", len(anchors))
	}
	if len(providers) != 3 {
		t.Fatalf("providers 数量 = %d, want 3", len(providers))
	}
	cnDomain, _ := providers["cn_domain"].(map[string]any)
	if cnDomain["behavior"] != "domain" || cnDomain["format"] != "mrs" {
		t.Fatalf("merge key 未展开: %#v", cnDomain)
	}
	cnIP, _ := providers["cn_ip"].(map[string]any)
	if cnIP["behavior"] != "ipcidr" {
		t.Fatalf("merge key 未展开: %#v", cnIP)
	}
}

func TestTransformWithRuleAnchor(t *testing.T) {
	st := testSettings()
	st.RuleProvidersEnabled = true
	st.RuleProvidersText = `rule-anchor:
  ip: &ip {type: http, interval: 86400, behavior: ipcidr, format: mrs}
  domain: &domain {type: http, interval: 86400, behavior: domain, format: mrs}
rule-providers:
  proxy:
    <<: *domain
    url: "https://example.com/proxy.mrs"
`
	st.RulesText = "- RULE-SET,proxy,PROXY\n- MATCH,PROXY"

	cfg, err := TransformAirportConfig([]byte(sampleAirportYAML), st)
	if err != nil {
		t.Fatalf("transform: %v", err)
	}
	if _, ok := cfg["rule-anchor"]; !ok {
		t.Fatalf("rule-anchor 应写入配置")
	}
	providers, _ := cfg["rule-providers"].(map[string]any)
	proxy, _ := providers["proxy"].(map[string]any)
	if proxy["behavior"] != "domain" {
		t.Fatalf("provider 未展开: %#v", proxy)
	}
	if errs := ValidateConfig(cfg, st); len(errs) > 0 {
		t.Fatalf("validate: %v", errs)
	}
}

func TestTransformCustomRuleProviders(t *testing.T) {
	st := testSettings()
	st.RuleProvidersEnabled = true
	st.RuleProvidersText = `proxy:
  type: http
  behavior: domain
  url: https://example.com/proxy.yaml
  path: ./proxy.yaml
  interval: 86400`
	st.RulesText = "- RULE-SET,proxy,PROXY\n- MATCH,PROXY"

	cfg, err := TransformAirportConfig([]byte(sampleAirportYAML), st)
	if err != nil {
		t.Fatalf("transform: %v", err)
	}
	providers, _ := cfg["rule-providers"].(map[string]any)
	if _, ok := providers["proxy"]; !ok {
		t.Fatalf("自定义 rule-providers 未写入: %v", providers)
	}
	if _, ok := providers["foo"]; ok {
		t.Fatalf("启用自定义后应替换机场自带的 rule-providers")
	}
	if errs := ValidateConfig(cfg, st); len(errs) > 0 {
		t.Fatalf("validate: %v", errs)
	}
}

func TestTransformKeepsAirportRuleProvidersWhenDisabled(t *testing.T) {
	st := testSettings() // 默认不启用
	cfg, err := TransformAirportConfig([]byte(sampleAirportYAML), st)
	if err != nil {
		t.Fatalf("transform: %v", err)
	}
	providers, _ := cfg["rule-providers"].(map[string]any)
	if _, ok := providers["foo"]; !ok {
		t.Fatalf("未启用时应保留机场的 rule-providers: %v", providers)
	}
}

func TestValidateMissingRuleProvider(t *testing.T) {
	st := testSettings()
	st.RulesText = "- RULE-SET,missing,PROXY\n- MATCH,PROXY"
	cfg, err := TransformAirportConfig([]byte(sampleAirportYAML), st)
	if err != nil {
		t.Fatalf("transform: %v", err)
	}
	errs := ValidateConfig(cfg, st)
	if !containsSub(errs, "rule-provider: missing") {
		t.Fatalf("应提示缺失 rule-provider，实际 = %v", errs)
	}
}

func TestValidateSettingsRuleProviders(t *testing.T) {
	st := testSettings()
	st.RuleProvidersEnabled = true
	st.RuleProvidersText = "   "
	if errs := validateClashRelaySettingsComplete(st); !containsSub(errs, "rule-providers") {
		t.Fatalf("启用但为空时应报错，实际 = %v", errs)
	}
}

func TestPullDoesNotPersistUntilSave(t *testing.T) {
	ctx := context.Background()

	db, err := store.NewDB(config.SQLiteConfig{Path: filepath.Join(t.TempDir(), "test.db")})
	if err != nil {
		t.Fatalf("NewDB: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/yaml")
		_, _ = io.WriteString(w, sampleAirportYAML)
	}))
	defer srv.Close()

	svc := NewClashRelayService(db, log.New(io.Discard, "", 0))
	st := testSettings()
	st.AirportURL = srv.URL
	st.RulesText = "- RULE-SET,foo,PROXY\n- MATCH,PROXY"
	if _, err := svc.SaveSettings(ctx, st); err != nil {
		t.Fatalf("SaveSettings: %v", err)
	}

	res, err := svc.Pull(ctx)
	if err != nil {
		t.Fatalf("Pull: %v", err)
	}
	if res.Draft == nil || strings.TrimSpace(res.Draft.Content) == "" {
		t.Fatalf("草稿内容为空")
	}
	if n, _ := db.CountClashRelayVersions(ctx); n != 0 {
		t.Fatalf("拉取后不应入库，实际版本数 = %d", n)
	}

	v, unchanged, err := svc.SaveVersion(ctx, res.Draft.Content, nil, "pull", res.Draft.AirportHash)
	if err != nil {
		t.Fatalf("SaveVersion: %v", err)
	}
	if unchanged {
		t.Fatalf("首次保存不应 unchanged")
	}
	if v.Source != "pull" {
		t.Fatalf("source = %q, want pull", v.Source)
	}
	if n, _ := db.CountClashRelayVersions(ctx); n != 1 {
		t.Fatalf("保存后版本数 = %d, want 1", n)
	}

	_, unchanged, err = svc.SaveVersion(ctx, res.Draft.Content, nil, "pull", res.Draft.AirportHash)
	if err != nil || !unchanged {
		t.Fatalf("同样内容再保存应 unchanged, err=%v unchanged=%v", err, unchanged)
	}

	v2, _, err := svc.SaveVersion(ctx, res.Draft.Content+"\n# edit\n", &v.ID, "manual", "")
	if err != nil {
		t.Fatalf("manual save: %v", err)
	}
	if v2.Source != "manual" || v2.ParentID == nil || *v2.ParentID != v.ID {
		t.Fatalf("manual 版本元信息异常: source=%q parent=%v", v2.Source, v2.ParentID)
	}
}
