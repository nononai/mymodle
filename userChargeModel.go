package model

import (
	"context"
	"fmt"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UserChargeModel = (*customUserChargeModel)(nil)

type (
	// UserChargeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserChargeModel.
	UserChargeModel interface {
		userChargeModel
		customUserChargeLogicModel
	}

	customUserChargeModel struct {
		*defaultUserChargeModel
	}

	customUserChargeLogicModel interface {
		GetUserChargeByUserId(ctx context.Context, userId int64) (*UserCharge, error)
	}
)

// NewUserChargeModel returns a model for the database table.
func NewUserChargeModel(conn *gorm.DB, c cache.CacheConf) UserChargeModel {
	return &customUserChargeModel{
		defaultUserChargeModel: newUserChargeModel(conn, c),
	}
}

func (m *defaultUserChargeModel) customCacheKeys(data *UserCharge) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
func (m *defaultUserChargeModel) GetUserChargeByUserId(ctx context.Context, userId int64) (*UserCharge, error) {
	minyunUserChargeUserIdKey := fmt.Sprintf("%s%v", cacheMinyunUserChargeIdPrefix, userId)
	var resp UserCharge
	err := m.QueryCtx(ctx, &resp, minyunUserChargeUserIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&UserCharge{}).Where("`user_id` = ?", userId).First(&resp).Error
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
