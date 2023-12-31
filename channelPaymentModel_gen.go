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
	cacheMinyunChannelPaymentIdPrefix = "cache:minyun:channelPayment:id:"
)

type (
	channelPaymentModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *ChannelPayment) error

		FindOne(ctx context.Context, id int64) (*ChannelPayment, error)
		Update(ctx context.Context, tx *gorm.DB, data *ChannelPayment) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultChannelPaymentModel struct {
		gormc.CachedConn
		table string
	}

	ChannelPayment struct {
		Id               int64          `gorm:"column:id"`
		PaymentName      string         `gorm:"column:payment_name"`
		PlatId           int64          `gorm:"column:plat_id"`
		ChannelId        int64          `gorm:"column:channel_id"`
		AliCertUrl       string         `gorm:"column:ali_cert_url"`        // 支付宝证书
		AliRootCertUrl   string         `gorm:"column:ali_root_cert_url"`   // 支付宝根证书
		AliPublicCertUrl string         `gorm:"column:ali_public_cert_url"` // 支付宝公钥证书url
		AliPublicKey     string         `gorm:"column:ali_public_key"`      // 支付宝公钥
		AliPriviteKey    string         `gorm:"column:ali_privite_key"`     // 支付宝私钥
		AliAes           string         `gorm:"column:ali_aes"`             // 支付宝aes
		WxKeyUrl         string         `gorm:"column:wx_key_url"`          // 微信key.pem url
		WxApiV3          string         `gorm:"column:wx_api_v3"`           // 微信v3api序列号
		WxApiV3Key       string         `gorm:"column:wx_api_v3_key"`       // 微信v3api key
		WxNotifyUrl      string         `gorm:"column:wx_notify_url"`       // 回调地址
		AliNotifyUrl     string         `gorm:"column:ali_notify_url"`      // 支付宝回调地址
		WxMchId          string         `gorm:"column:wx_mch_id"`
		WxMchSecret      string         `gorm:"column:wx_mch_secret"`
		AliMchId         string         `gorm:"column:ali_mch_id"`
		AliMchSecret     string         `gorm:"column:ali_mch_secret"`
		CreatedAt        time.Time   `gorm:"column:created_at"`
		UpdatedAt        time.Time   `gorm:"column:updated_at"`
		DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at;index"`
	}
)

func (ChannelPayment) TableName() string {
	return "`channel_payment`"
}

func newChannelPaymentModel(conn *gorm.DB, c cache.CacheConf) *defaultChannelPaymentModel {
	return &defaultChannelPaymentModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`channel_payment`",
	}
}

func (m *defaultChannelPaymentModel) Insert(ctx context.Context, tx *gorm.DB, data *ChannelPayment) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultChannelPaymentModel) FindOne(ctx context.Context, id int64) (*ChannelPayment, error) {
	minyunChannelPaymentIdKey := fmt.Sprintf("%s%v", cacheMinyunChannelPaymentIdPrefix, id)
	var resp ChannelPayment
	err := m.QueryCtx(ctx, &resp, minyunChannelPaymentIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&ChannelPayment{}).Where("`id` = ?", id).First(&resp).Error
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

func (m *defaultChannelPaymentModel) Update(ctx context.Context, tx *gorm.DB, data *ChannelPayment) error {
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

func (m *defaultChannelPaymentModel) getCacheKeys(data *ChannelPayment) []string {
	if data == nil {
		return []string{}
	}
	minyunChannelPaymentIdKey := fmt.Sprintf("%s%v", cacheMinyunChannelPaymentIdPrefix, data.Id)
	cacheKeys := []string{
		minyunChannelPaymentIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultChannelPaymentModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
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
		return db.Delete(&ChannelPayment{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultChannelPaymentModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultChannelPaymentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMinyunChannelPaymentIdPrefix, primary)
}

func (m *defaultChannelPaymentModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&ChannelPayment{}).Where("`id` = ?", primary).Take(v).Error
}
