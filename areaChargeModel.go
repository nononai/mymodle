package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ AreaChargeModel = (*customAreaChargeModel)(nil)

type (
	// AreaChargeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAreaChargeModel.
	AreaChargeModel interface {
		areaChargeModel
		customAreaChargeLogicModel
	}

	customAreaChargeModel struct {
		*defaultAreaChargeModel
	}

	customAreaChargeLogicModel interface {
	}
)

// NewAreaChargeModel returns a model for the database table.
func NewAreaChargeModel(conn *gorm.DB, c cache.CacheConf) AreaChargeModel {
	return &customAreaChargeModel{
		defaultAreaChargeModel: newAreaChargeModel(conn, c),
	}
}

func (m *defaultAreaChargeModel) customCacheKeys(data *AreaCharge) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
