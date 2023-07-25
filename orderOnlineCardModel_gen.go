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
	cacheMinyunOrderOnlineCardIdPrefix = "cache:minyun:orderOnlineCard:id:"
)

type (
	orderOnlineCardModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *OrderOnlineCard) error

		FindOne(ctx context.Context, id int64) (*OrderOnlineCard, error)
		Update(ctx context.Context, tx *gorm.DB, data *OrderOnlineCard) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultOrderOnlineCardModel struct {
		gormc.CachedConn
		table string
	}

	OrderOnlineCard struct {
		Id            int64          `gorm:"column:id"`
		OrderId       string         `gorm:"column:order_id"`    // 内部订单号
		OutNo         string         `gorm:"column:out_no"`      // 外部订单号
		OrderMoney    int64          `gorm:"column:order_money"` // 单位分
		PayType       int64          `gorm:"column:pay_type"`
		PayStat       int64          `gorm:"column:pay_stat"`
		GiveMoney     int64          `gorm:"column:give_money"` // 单位分
		PayTime       int64          `gorm:"column:pay_time"`
		RefundMoney   int64          `gorm:"column:refund_money"` // 单位分
		RefundOrderId string         `gorm:"column:refund_order_id"`
		ChannelId     int64          `gorm:"column:channel_id"`
		ShopId        int64          `gorm:"column:shop_id"`
		SiteId        int64          `gorm:"column:site_id"`
		UserId        int64          `gorm:"column:user_id"`
		CreatedAt     time.Time      `gorm:"column:created_at"`
		UpdatedAt     time.Time      `gorm:"column:updated_at"`
		DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index"`
		Remark        string         `gorm:"column:remark"`
		OrderTitle    string         `gorm:"column:order_title"`
		CardId        int64          `gorm:"column:card_id"`
	}
)

func (OrderOnlineCard) TableName() string {
	return "`order_online_card`"
}

func newOrderOnlineCardModel(conn *gorm.DB, c cache.CacheConf) *defaultOrderOnlineCardModel {
	return &defaultOrderOnlineCardModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`order_online_card`",
	}
}

func (m *defaultOrderOnlineCardModel) Insert(ctx context.Context, tx *gorm.DB, data *OrderOnlineCard) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultOrderOnlineCardModel) FindOne(ctx context.Context, id int64) (*OrderOnlineCard, error) {
	minyunOrderOnlineCardIdKey := fmt.Sprintf("%s%v", cacheMinyunOrderOnlineCardIdPrefix, id)
	var resp OrderOnlineCard
	err := m.QueryCtx(ctx, &resp, minyunOrderOnlineCardIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&OrderOnlineCard{}).Where("`id` = ?", id).First(&resp).Error
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

func (m *defaultOrderOnlineCardModel) Update(ctx context.Context, tx *gorm.DB, data *OrderOnlineCard) error {
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

func (m *defaultOrderOnlineCardModel) getCacheKeys(data *OrderOnlineCard) []string {
	if data == nil {
		return []string{}
	}
	minyunOrderOnlineCardIdKey := fmt.Sprintf("%s%v", cacheMinyunOrderOnlineCardIdPrefix, data.Id)
	cacheKeys := []string{
		minyunOrderOnlineCardIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultOrderOnlineCardModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
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
		return db.Delete(&OrderOnlineCard{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultOrderOnlineCardModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultOrderOnlineCardModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMinyunOrderOnlineCardIdPrefix, primary)
}

func (m *defaultOrderOnlineCardModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&OrderOnlineCard{}).Where("`id` = ?", primary).Take(v).Error
}