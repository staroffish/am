package model

import "time"

// ClashRelaySettings 是「Clash 中转」模块的单行配置（id 固定为 1）。
type ClashRelaySettings struct {
	ID                   uint       `gorm:"column:id;primaryKey" json:"-"`
	AirportURL           string     `gorm:"column:airport_url" json:"airport_url"`
	AirportUA            string     `gorm:"column:airport_ua" json:"airport_ua"`
	VPSHost              string     `gorm:"column:vps_host" json:"vps_host"`
	VPSPort              int        `gorm:"column:vps_port" json:"vps_port"`
	VPSCipher            string     `gorm:"column:vps_cipher" json:"vps_cipher"`
	VPSPassword          string     `gorm:"column:vps_password" json:"vps_password"`
	VPSNodeName          string     `gorm:"column:vps_node_name" json:"vps_node_name"`
	VPSGroupName         string     `gorm:"column:vps_group_name" json:"vps_group_name"`
	KeepVersions         int        `gorm:"column:keep_versions" json:"keep_versions"`
	RulesText            string     `gorm:"column:rules_text;type:text" json:"rules_text"`
	RuleProvidersEnabled bool       `gorm:"column:rule_providers_enabled" json:"rule_providers_enabled"`
	RuleProvidersText    string     `gorm:"column:rule_providers_text;type:text" json:"rule_providers_text"`
	LastPullAt           *time.Time `gorm:"column:last_pull_at" json:"last_pull_at"`
	LastPullStatus       string     `gorm:"column:last_pull_status" json:"last_pull_status"`
	LastPullError        string     `gorm:"column:last_pull_error;type:text" json:"last_pull_error"`
	UpdatedAt            time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

func (ClashRelaySettings) TableName() string { return "clash_relay_settings" }

// ClashRelayVersion 是一次生成结果的完整快照。
type ClashRelayVersion struct {
	ID          uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at;index" json:"created_at"`
	Source      string    `gorm:"column:source" json:"source"` // pull | manual
	ParentID    *uint     `gorm:"column:parent_id" json:"parent_id"`
	AirportHash string    `gorm:"column:airport_hash" json:"airport_hash"`
	Content     string    `gorm:"column:content;type:text" json:"content,omitempty"`
	ContentHash string    `gorm:"column:content_hash;index" json:"content_hash"`
	NodeCount   int       `gorm:"column:node_count" json:"node_count"`
	IsCurrent   bool      `gorm:"column:is_current;index" json:"is_current"`
}

func (ClashRelayVersion) TableName() string { return "clash_relay_versions" }
