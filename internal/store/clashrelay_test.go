package store

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/staroffish/am/internal/model"
)

func TestClashRelaySettings(t *testing.T) {
	ctx := context.Background()
	db := newTestDB(t)

	st, err := db.GetClashRelaySettings(ctx)
	if err != nil {
		t.Fatalf("GetClashRelaySettings: %v", err)
	}
	if st != nil {
		t.Fatalf("空库应返回 nil, got %#v", st)
	}

	in := &model.ClashRelaySettings{
		AirportURL:   "https://airport.example.com/sub",
		AirportUA:    "clash.meta",
		VPSHost:      "vps.example.com",
		VPSPort:      8443,
		VPSCipher:    "aes-256-gcm",
		VPSPassword:  "pass",
		VPSNodeName:  "VPS-SS",
		VPSGroupName: "PROXY",
		KeepVersions: 10,
		RulesText:    "- MATCH,PROXY",
	}
	if err := db.SaveClashRelaySettings(ctx, in); err != nil {
		t.Fatalf("SaveClashRelaySettings: %v", err)
	}

	at := time.Now()
	if err := db.UpdateClashRelayPullStatus(ctx, at, "ok", ""); err != nil {
		t.Fatalf("UpdateClashRelayPullStatus: %v", err)
	}

	got, err := db.GetClashRelaySettings(ctx)
	if err != nil {
		t.Fatalf("GetClashRelaySettings: %v", err)
	}
	if got.VPSHost != "vps.example.com" || got.VPSPort != 8443 {
		t.Errorf("配置字段被破坏: %#v", got)
	}
	if got.LastPullStatus != "ok" || got.LastPullAt == nil {
		t.Errorf("拉取状态未保存: status=%q at=%v", got.LastPullStatus, got.LastPullAt)
	}
}

func TestClashRelayVersionsKeepAndActivate(t *testing.T) {
	ctx := context.Background()
	db := newTestDB(t)

	for i := 0; i < 5; i++ {
		v := &model.ClashRelayVersion{
			Source:      "pull",
			Content:     fmt.Sprintf("content-%d", i),
			ContentHash: fmt.Sprintf("hash-%d", i),
			NodeCount:   i,
		}
		if err := db.CreateClashRelayVersion(ctx, v, 3); err != nil {
			t.Fatalf("CreateClashRelayVersion %d: %v", i, err)
		}
	}

	list, err := db.ListClashRelayVersions(ctx)
	if err != nil {
		t.Fatalf("ListClashRelayVersions: %v", err)
	}
	if len(list) != 3 {
		t.Fatalf("保留 3 个版本，实际 %d", len(list))
	}
	if list[0].ID != 5 || list[0].Content != "content-4" {
		t.Errorf("最新版本异常: %#v", list[0])
	}
	if !list[0].IsCurrent {
		t.Errorf("最新版本应为当前")
	}

	cur, err := db.CurrentClashRelayVersion(ctx)
	if err != nil || cur == nil || cur.ID != 5 {
		t.Fatalf("CurrentClashRelayVersion = %#v, err=%v", cur, err)
	}

	if err := db.ActivateClashRelayVersion(ctx, 4); err != nil {
		t.Fatalf("ActivateClashRelayVersion: %v", err)
	}
	cur, err = db.CurrentClashRelayVersion(ctx)
	if err != nil || cur == nil || cur.ID != 4 {
		t.Fatalf("切换后当前版本应为 4，实际 %#v, err=%v", cur, err)
	}

	n, err := db.CountClashRelayVersions(ctx)
	if err != nil || n != 3 {
		t.Fatalf("CountClashRelayVersions = %d, err=%v", n, err)
	}

	if _, err := db.GetClashRelayVersion(ctx, 999); err == nil {
		t.Fatalf("不存在的版本应返回错误")
	}
}

func TestClashRelayDeleteVersion(t *testing.T) {
	ctx := context.Background()
	db := newTestDB(t)

	for i := 0; i < 3; i++ {
		v := &model.ClashRelayVersion{
			Source:      "pull",
			Content:     fmt.Sprintf("content-%d", i),
			ContentHash: fmt.Sprintf("hash-%d", i),
		}
		if err := db.CreateClashRelayVersion(ctx, v, 10); err != nil {
			t.Fatalf("create %d: %v", i, err)
		}
	}
	// 当前版本是 3
	if err := db.DeleteClashRelayVersion(ctx, 2); err != nil {
		t.Fatalf("删除非当前版本应成功: %v", err)
	}
	if n, _ := db.CountClashRelayVersions(ctx); n != 2 {
		t.Fatalf("删除后数量 = %d, want 2", n)
	}
	if err := db.DeleteClashRelayVersion(ctx, 3); !errors.Is(err, ErrClashRelayVersionCurrent) {
		t.Fatalf("删除当前版本应被拒绝，实际 err = %v", err)
	}
	if err := db.DeleteClashRelayVersion(ctx, 999); err == nil {
		t.Fatalf("删除不存在的版本应报错")
	}
}
