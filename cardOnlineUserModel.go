package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ CardOnlineUserModel = (*customCardOnlineUserModel)(nil)

type (
	// CardOnlineUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCardOnlineUserModel.
	CardOnlineUserModel interface {
		cardOnlineUserModel
		customCardOnlineUserLogicModel
	}

	customCardOnlineUserModel struct {
		*defaultCardOnlineUserModel
	}

	customCardOnlineUserLogicModel interface {
		GetUserBalanceCard(ctx context.Context, userId, siteId int64) (*CardOnlineUser, error)
		GetUserOnlineCards(ctx context.Context, userId, siteId int64) (*CardOnlineUser, error)
		GetUserAccounts(ctx context.Context, userId int64) ([]*CardOnlineUser, error)
		GetUserMonthList(ctx context.Context, userId int64) ([]*CardOnlineUser, error)
	}
)

// NewCardOnlineUserModel returns a model for the database table.
func NewCardOnlineUserModel(conn *gorm.DB, c cache.CacheConf) CardOnlineUserModel {
	return &customCardOnlineUserModel{
		defaultCardOnlineUserModel: newCardOnlineUserModel(conn, c),
	}
}

func (m *defaultCardOnlineUserModel) customCacheKeys(data *CardOnlineUser) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
func (m *defaultCardOnlineUserModel) GetUserBalanceCard(ctx context.Context, userId, siteId int64) (*CardOnlineUser, error) {
	var resp CardOnlineUser
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&CardOnlineUser{}).Where("`user_id` = ? AND `site_id` = ? AND `card_type` = 4", userId, siteId).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultCardOnlineUserModel) GetUserOnlineCards(ctx context.Context, userId, siteId int64) (*CardOnlineUser, error) {
	var resp CardOnlineUser
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&CardOnlineUser{}).Where("`user_id` = ? AND `site_id` = ? AND `card_type` != 4", userId, siteId).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 获取用户储值卡信息list
func (m *defaultCardOnlineUserModel) GetUserAccounts(ctx context.Context, userId int64) ([]*CardOnlineUser, error) {
	var resp []*CardOnlineUser
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&CardOnlineUser{}).Where("`user_id` = ? and card_type=4", userId).Preload("Site").Find(&resp).Error
	})
	switch err {
	case nil:
		return resp, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 获取用户线上卡信息list
func (m *defaultCardOnlineUserModel) GetUserMonthList(ctx context.Context, userId int64) ([]*CardOnlineUser, error) {
	var resp []*CardOnlineUser
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&CardOnlineUser{}).Where("`user_id` = ? and card_type!=4", userId).Preload("Site").Find(&resp).Error
	})
	switch err {
	case nil:
		return resp, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
