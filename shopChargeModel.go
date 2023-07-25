package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ ShopChargeModel = (*customShopChargeModel)(nil)

type (
	// ShopChargeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShopChargeModel.
	ShopChargeModel interface {
		shopChargeModel
		customShopChargeLogicModel
	}

	customShopChargeModel struct {
		*defaultShopChargeModel
	}

	customShopChargeLogicModel interface {
	}
)

// NewShopChargeModel returns a model for the database table.
func NewShopChargeModel(conn *gorm.DB, c cache.CacheConf) ShopChargeModel {
	return &customShopChargeModel{
		defaultShopChargeModel: newShopChargeModel(conn, c),
	}
}

func (m *defaultShopChargeModel) customCacheKeys(data *ShopCharge) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
