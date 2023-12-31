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
	cacheMinyunCardOnlineIdPrefix = "cache:minyun:cardOnline:id:"
)

type (
	cardOnlineModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *CardOnline) error

		FindOne(ctx context.Context, id int64) (*CardOnline, error)
		Update(ctx context.Context, tx *gorm.DB, data *CardOnline) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultCardOnlineModel struct {
		gormc.CachedConn
		table string
	}

	CardOnline struct {
		Id           int64          `gorm:"column:id"`
		ChannelId    int64          `gorm:"column:channel_id"`
		ShopId       int64          `gorm:"column:shop_id"`
		SiteId       int64          `gorm:"column:site_id"`
		CardType     int64          `gorm:"column:card_type"` // 1,储值卡2时长卡3次数卡4电量卡（储值卡等同于用户余额）
		CardName     string         `gorm:"column:card_name"`
		TimeType     int64          `gorm:"column:time_type"`      // 1，长期有效，2固定时间，3，相对时间
		UseType      int64          `gorm:"column:use_type"`       // 卡限制，1，按每天多少次，2，按有效期合计多次次，如果是1，按照valuedays来算次数，如果是按合计，就是按实际次数ablecount来控制
		PowerPrice   string         `gorm:"column:power_price"`    // 功率-价格如300w以下：20元/时间---存json
		RechargeGive string         `gorm:"column:recharge_give"`  // 充值赠送，存json
		AbleCount    int64          `gorm:"column:able_count"`     // 次数
		AbleTimes    int64          `gorm:"column:able_times"`     // 单位分钟
		AbleEnergy   float64        `gorm:"column:able_energy"`    // 单位 度
		MoreUseCount int64          `gorm:"column:more_use_count"` // 同时允许几个订单运行
		IsTransfer   int64          `gorm:"column:is_transfer"`    // 是否允许转让，1允许，2不允许
		CardText     string         `gorm:"column:card_text"`      // 卡片说明
		ValueDays    int64          `gorm:"column:value_days"`     // 有效天数-对应3相对时间
		EndTime      int64          `gorm:"column:end_time"`       // 卡有效期截止日-对应2固定时间
		IsMoreBuy    int64          `gorm:"column:is_more_buy"`    // 是否允许重复购买-加油包理念
		MoreBuyItems int64          `gorm:"column:more_buy_items"` // 有效期内多次购买的资源
		MoreBuyPrice int64          `gorm:"column:more_buy_price"` // 有效期内多次购买的金额
		ValueByDay   int64          `gorm:"column:value_by_day"`   // 如果usetype是1的，此字段有意义，设置每日可以使用的资源次、时间、电量
		UpdatedAt    time.Time      `gorm:"column:updated_at"`
		DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;index"`
		CreatedAt    time.Time      `gorm:"column:created_at"`
		Remark       string         `gorm:"column:remark"`
		Site 		SiteCharge 		`gorm:"foreignKey:SiteId"`
	}
)

func (CardOnline) TableName() string {
	return "`card_online`"
}

func newCardOnlineModel(conn *gorm.DB, c cache.CacheConf) *defaultCardOnlineModel {
	return &defaultCardOnlineModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`card_online`",
	}
}

func (m *defaultCardOnlineModel) Insert(ctx context.Context, tx *gorm.DB, data *CardOnline) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultCardOnlineModel) FindOne(ctx context.Context, id int64) (*CardOnline, error) {
	minyunCardOnlineIdKey := fmt.Sprintf("%s%v", cacheMinyunCardOnlineIdPrefix, id)
	var resp CardOnline
	err := m.QueryCtx(ctx, &resp, minyunCardOnlineIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&CardOnline{}).Where("`id` = ?", id).First(&resp).Error
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

func (m *defaultCardOnlineModel) Update(ctx context.Context, tx *gorm.DB, data *CardOnline) error {
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

func (m *defaultCardOnlineModel) getCacheKeys(data *CardOnline) []string {
	if data == nil {
		return []string{}
	}
	minyunCardOnlineIdKey := fmt.Sprintf("%s%v", cacheMinyunCardOnlineIdPrefix, data.Id)
	cacheKeys := []string{
		minyunCardOnlineIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultCardOnlineModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
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
		return db.Delete(&CardOnline{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultCardOnlineModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultCardOnlineModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMinyunCardOnlineIdPrefix, primary)
}

func (m *defaultCardOnlineModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&CardOnline{}).Where("`id` = ?", primary).Take(v).Error
}
