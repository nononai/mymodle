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
	cacheMinyunCardOnlineUserIdPrefix = "cache:minyun:cardOnlineUser:id:"
)

type (
	cardOnlineUserModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *CardOnlineUser) error

		FindOne(ctx context.Context, id int64) (*CardOnlineUser, error)
		Update(ctx context.Context, tx *gorm.DB, data *CardOnlineUser) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultCardOnlineUserModel struct {
		gormc.CachedConn
		table string
	}

	CardOnlineUser struct {
		Id          int64           `gorm:"column:id"`
		UserId      int64   `gorm:"column:user_id"`      // userbasic外键id
		CardId      int64   `gorm:"column:card_id"`      // 系统创建卡的id外键
		CardType    int64   `gorm:"column:card_type"`    // 1,储值卡2时长卡3次数卡4电量卡（储值卡等同于用户余额）
		AbleEnergy  float64 `gorm:"column:able_energy"`  // 单位 度
		AbleTimes   int64   `gorm:"column:able_times"`   // 单位分钟
		AbleCount   int64   `gorm:"column:able_count"`   // 次
		LimitPower  int64   `gorm:"column:limit_power"`  // 卡限制功率
		ExpiredTime int64   `gorm:"column:expired_time"` // 到期时间
		CreatedAt   time.Time  `gorm:"column:created_at"`
		UpdatedAt   time.Time    `gorm:"column:updated_at"`
		DeletedAt   gorm.DeletedAt  `gorm:"column:deleted_at;index"`
		ChannelId   int64   `gorm:"column:channel_id"`
		ShopId      int64   `gorm:"column:shop_id"`
		SiteId      int64   `gorm:"column:site_id"`
		Remark      string  `gorm:"column:remark"`
		BasicMoney  int64   `gorm:"column:basic_money"` // 单位 分 -cardtype=1此字段有意义，类似用户余额。消费先消费基本余额后在消费赠送金额，退款赠送金额清零
		GiveMoney   int64   `gorm:"column:give_money"`  // 单位 分 赠送金额  于上相同
		State       int64   `gorm:"column:state"`       // 卡状态，1，正常，2，到期，
		Site 		SiteCharge 		`gorm:"foreignKey:SiteId"`
	}
)

func (CardOnlineUser) TableName() string {
	return "`card_online_user`"
}

func newCardOnlineUserModel(conn *gorm.DB, c cache.CacheConf) *defaultCardOnlineUserModel {
	return &defaultCardOnlineUserModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`card_online_user`",
	}
}

func (m *defaultCardOnlineUserModel) Insert(ctx context.Context, tx *gorm.DB, data *CardOnlineUser) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultCardOnlineUserModel) FindOne(ctx context.Context, id int64) (*CardOnlineUser, error) {
	minyunCardOnlineUserIdKey := fmt.Sprintf("%s%v", cacheMinyunCardOnlineUserIdPrefix, id)
	var resp CardOnlineUser
	err := m.QueryCtx(ctx, &resp, minyunCardOnlineUserIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&CardOnlineUser{}).Where("`id` = ?", id).First(&resp).Error
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

func (m *defaultCardOnlineUserModel) Update(ctx context.Context, tx *gorm.DB, data *CardOnlineUser) error {
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

func (m *defaultCardOnlineUserModel) getCacheKeys(data *CardOnlineUser) []string {
	if data == nil {
		return []string{}
	}
	minyunCardOnlineUserIdKey := fmt.Sprintf("%s%v", cacheMinyunCardOnlineUserIdPrefix, data.Id)
	cacheKeys := []string{
		minyunCardOnlineUserIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultCardOnlineUserModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
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
		return db.Delete(&CardOnlineUser{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultCardOnlineUserModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultCardOnlineUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMinyunCardOnlineUserIdPrefix, primary)
}

func (m *defaultCardOnlineUserModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&CardOnlineUser{}).Where("`id` = ?", primary).Take(v).Error
}