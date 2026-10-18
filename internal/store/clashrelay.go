package store

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/staroffish/am/internal/model"
)

// ---------------- Clash 中转设置 ----------------

// GetClashRelaySettings 返回单行设置；不存在时返回 (nil, nil)。
func (d *DB) GetClashRelaySettings(ctx context.Context) (*model.ClashRelaySettings, error) {
	var s model.ClashRelaySettings
	err := d.gorm.WithContext(ctx).First(&s, 1).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &s, nil
}

// SaveClashRelaySettings 以 id=1 覆盖保存设置。
func (d *DB) SaveClashRelaySettings(ctx context.Context, s *model.ClashRelaySettings) error {
	s.ID = 1
	s.UpdatedAt = time.Now()
	return d.gorm.WithContext(ctx).Save(s).Error
}

// UpdateClashRelayPullStatus 只更新"上次拉取"相关字段，不影响配置。
func (d *DB) UpdateClashRelayPullStatus(ctx context.Context, at time.Time, status, errMsg string) error {
	return d.gorm.WithContext(ctx).Model(&model.ClashRelaySettings{}).
		Where("id = ?", 1).
		Updates(map[string]any{
			"last_pull_at":     at,
			"last_pull_status": status,
			"last_pull_error":  errMsg,
		}).Error
}

// ---------------- Clash 中转版本 ----------------

// CreateClashRelayVersion 在单个事务里：取消旧当前版本 → 插入新版本并设为当前 → 清理超出 keep 的旧版本。
func (d *DB) CreateClashRelayVersion(ctx context.Context, v *model.ClashRelayVersion, keep int) error {
	if keep < 1 {
		keep = 10
	}
	return d.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&model.ClashRelayVersion{}).
			Where("is_current = ?", true).
			Update("is_current", false).Error; err != nil {
			return err
		}

		v.IsCurrent = true
		if v.CreatedAt.IsZero() {
			v.CreatedAt = time.Now()
		}
		if err := tx.Create(v).Error; err != nil {
			return err
		}

		var keepIDs []uint
		if err := tx.Model(&model.ClashRelayVersion{}).
			Order("id desc").Limit(keep).Pluck("id", &keepIDs).Error; err != nil {
			return err
		}
		if len(keepIDs) == 0 {
			return nil
		}
		return tx.Where("id NOT IN ?", keepIDs).Delete(&model.ClashRelayVersion{}).Error
	})
}

// ListClashRelayVersions 按版本号倒序返回全部版本（含 content，由调用方按需裁剪）。
func (d *DB) ListClashRelayVersions(ctx context.Context) ([]model.ClashRelayVersion, error) {
	var list []model.ClashRelayVersion
	if err := d.gorm.WithContext(ctx).Order("id desc").Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (d *DB) GetClashRelayVersion(ctx context.Context, id uint) (*model.ClashRelayVersion, error) {
	var v model.ClashRelayVersion
	if err := d.gorm.WithContext(ctx).First(&v, id).Error; err != nil {
		return nil, err
	}
	return &v, nil
}

// CurrentClashRelayVersion 返回当前版本；没有时返回 (nil, nil)。
func (d *DB) CurrentClashRelayVersion(ctx context.Context) (*model.ClashRelayVersion, error) {
	var v model.ClashRelayVersion
	err := d.gorm.WithContext(ctx).Where("is_current = ?", true).Order("id desc").First(&v).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &v, nil
}

// ActivateClashRelayVersion 把指定版本设为当前。
func (d *DB) ActivateClashRelayVersion(ctx context.Context, id uint) error {
	return d.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var v model.ClashRelayVersion
		if err := tx.First(&v, id).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.ClashRelayVersion{}).
			Where("is_current = ?", true).
			Update("is_current", false).Error; err != nil {
			return err
		}
		return tx.Model(&model.ClashRelayVersion{}).
			Where("id = ?", id).
			Update("is_current", true).Error
	})
}

// ErrClashRelayVersionCurrent 表示试图删除当前版本。
var ErrClashRelayVersionCurrent = errors.New("当前版本不能删除，请先切换到其他版本")

// DeleteClashRelayVersion 删除指定版本；当前版本不允许删除。
func (d *DB) DeleteClashRelayVersion(ctx context.Context, id uint) error {
	return d.gorm.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var v model.ClashRelayVersion
		if err := tx.First(&v, id).Error; err != nil {
			return err
		}
		if v.IsCurrent {
			return ErrClashRelayVersionCurrent
		}
		return tx.Delete(&model.ClashRelayVersion{}, id).Error
	})
}

func (d *DB) CountClashRelayVersions(ctx context.Context) (int64, error) {
	var n int64
	if err := d.gorm.WithContext(ctx).Model(&model.ClashRelayVersion{}).Count(&n).Error; err != nil {
		return 0, err
	}
	return n, nil
}
