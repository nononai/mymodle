package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ LogUserShopModel = (*customLogUserShopModel)(nil)

type (
	// LogUserShopModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLogUserShopModel.
	LogUserShopModel interface {
		logUserShopModel
		customLogUserShopLogicModel
	}

	customLogUserShopModel struct {
		*defaultLogUserShopModel
	}

	customLogUserShopLogicModel interface {
	}
)

// NewLogUserShopModel returns a model for the database table.
func NewLogUserShopModel(conn *gorm.DB, c cache.CacheConf) LogUserShopModel {
	return &customLogUserShopModel{
		defaultLogUserShopModel: newLogUserShopModel(conn, c),
	}
}

func (m *defaultLogUserShopModel) customCacheKeys(data *LogUserShop) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
