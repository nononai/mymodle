// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"time"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (
	cacheMinyunLogShopChargeIdPrefix = "cache:minyun:logShopCharge:id:"
)

type (
	logShopChargeModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *LogShopCharge) error

		FindOne(ctx context.Context, id int64) (*LogShopCharge, error)
		Update(ctx context.Context, tx *gorm.DB, data *LogShopCharge) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultLogShopChargeModel struct {
		gormc.CachedConn
		table string
	}

	LogShopCharge struct {
		Id          int64          `gorm:"column:id"`
		ShopId      int64          `gorm:"column:shop_id"`
		LogType     int64          `gorm:"column:log_type"`     // 1,订单收入2，订单退款3，账户提现4，手续费扣除5，服务购买
		LogValue    int64          `gorm:"column:log_value"`    // 单位 分
		LogName     string         `gorm:"column:log_name"`     // 操作项目名称，如 充电订单收入等
		Symbol      int64          `gorm:"column:symbol"`       // 操作类型：1，增加，2减少
		TriggerTime int64          `gorm:"column:trigger_time"` // 触发时间
		OperatId    int64          `gorm:"column:operat_id"`    // 操作者id：0为系统自动 非0则为操作人的id
		CreatedAt   time.Time      `gorm:"column:created_at"`
		UpdatedAt   time.Time      `gorm:"column:updated_at"`
		DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index"`
		UpingValue  string         `gorm:"column:uping_value"` // 操作前账户剩余资源
		Remark      string         `gorm:"column:remark"`
		SiteId      int64          `gorm:"column:site_id"`
	}
)

func (LogShopCharge) TableName() string {
	return "`log_shop_charge`"
}

func newLogShopChargeModel(conn *gorm.DB, c cache.CacheConf) *defaultLogShopChargeModel {
	return &defaultLogShopChargeModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`log_shop_charge`",
	}
}

func (m *defaultLogShopChargeModel) Insert(ctx context.Context, tx *gorm.DB, data *LogShopCharge) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultLogShopChargeModel) FindOne(ctx context.Context, id int64) (*LogShopCharge, error) {
	minyunLogShopChargeIdKey := fmt.Sprintf("%s%v", cacheMinyunLogShopChargeIdPrefix, id)
	var resp LogShopCharge
	err := m.QueryCtx(ctx, &resp, minyunLogShopChargeIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&LogShopCharge{}).Where("`id` = ?", id).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLogShopChargeModel) Update(ctx context.Context, tx *gorm.DB, data *LogShopCharge) error {
	old, err := m.FindOne(ctx, data.Id)
	if err != nil && err != ErrNotFound {
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(data).Error
	}, m.getCacheKeys(old)...)
	return err
}

func (m *defaultLogShopChargeModel) getCacheKeys(data *LogShopCharge) []string {
	if data == nil {
		return []string{}
	}
	minyunLogShopChargeIdKey := fmt.Sprintf("%s%v", cacheMinyunLogShopChargeIdPrefix, data.Id)
	cacheKeys := []string{
		minyunLogShopChargeIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultLogShopChargeModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		if err == ErrNotFound {
			return nil
		}
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&LogShopCharge{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultLogShopChargeModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultLogShopChargeModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMinyunLogShopChargeIdPrefix, primary)
}

func (m *defaultLogShopChargeModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&LogShopCharge{}).Where("`id` = ?", primary).Take(v).Error
}