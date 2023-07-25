package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UserShopModel = (*customUserShopModel)(nil)

type (
	// UserShopModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserShopModel.
	UserShopModel interface {
		userShopModel
		customUserShopLogicModel
	}

	customUserShopModel struct {
		*defaultUserShopModel
	}

	customUserShopLogicModel interface {
	}
)

// NewUserShopModel returns a model for the database table.
func NewUserShopModel(conn *gorm.DB, c cache.CacheConf) UserShopModel {
	return &customUserShopModel{
		defaultUserShopModel: newUserShopModel(conn, c),
	}
}

func (m *defaultUserShopModel) customCacheKeys(data *UserShop) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
