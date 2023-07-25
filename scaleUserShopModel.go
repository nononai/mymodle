package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ ScaleUserShopModel = (*customScaleUserShopModel)(nil)

type (
	// ScaleUserShopModel is an interface to be customized, add more methods here,
	// and implement the added methods in customScaleUserShopModel.
	ScaleUserShopModel interface {
		scaleUserShopModel
		customScaleUserShopLogicModel
	}

	customScaleUserShopModel struct {
		*defaultScaleUserShopModel
	}

	customScaleUserShopLogicModel interface {
	}
)

// NewScaleUserShopModel returns a model for the database table.
func NewScaleUserShopModel(conn *gorm.DB, c cache.CacheConf) ScaleUserShopModel {
	return &customScaleUserShopModel{
		defaultScaleUserShopModel: newScaleUserShopModel(conn, c),
	}
}

func (m *defaultScaleUserShopModel) customCacheKeys(data *ScaleUserShop) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
