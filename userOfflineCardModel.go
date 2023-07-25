package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UserOfflineCardModel = (*customUserOfflineCardModel)(nil)

type (
	// UserOfflineCardModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserOfflineCardModel.
	UserOfflineCardModel interface {
		userOfflineCardModel
		customUserOfflineCardLogicModel
	}

	customUserOfflineCardModel struct {
		*defaultUserOfflineCardModel
	}

	customUserOfflineCardLogicModel interface {
		GetUserOfflineCardByUserId(ctx context.Context, userId int64) ([]*UserOfflineCard, error)
	}
)

// NewUserOfflineCardModel returns a model for the database table.
func NewUserOfflineCardModel(conn *gorm.DB, c cache.CacheConf) UserOfflineCardModel {
	return &customUserOfflineCardModel{
		defaultUserOfflineCardModel: newUserOfflineCardModel(conn, c),
	}
}

func (m *defaultUserOfflineCardModel) customCacheKeys(data *UserOfflineCard) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

// 通过用户id获取用户充电卡列表
func (m *defaultUserOfflineCardModel) GetUserOfflineCardByUserId(ctx context.Context, userId int64) ([]*UserOfflineCard, error) {
	var resp []*UserOfflineCard
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&UserOfflineCard{}).Where("`user_id` = ?", userId).Find(&resp).Error
	})
	fmt.Println("err  temp:", err)
	switch err {
	case nil:
		return resp, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
