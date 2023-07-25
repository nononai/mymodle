package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ LogShopChargeModel = (*customLogShopChargeModel)(nil)

type (
	// LogShopChargeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLogShopChargeModel.
	LogShopChargeModel interface {
		logShopChargeModel
		customLogShopChargeLogicModel
	}

	customLogShopChargeModel struct {
		*defaultLogShopChargeModel
	}

	customLogShopChargeLogicModel interface {
	}
)

// NewLogShopChargeModel returns a model for the database table.
func NewLogShopChargeModel(conn *gorm.DB, c cache.CacheConf) LogShopChargeModel {
	return &customLogShopChargeModel{
		defaultLogShopChargeModel: newLogShopChargeModel(conn, c),
	}
}

func (m *defaultLogShopChargeModel) customCacheKeys(data *LogShopCharge) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
