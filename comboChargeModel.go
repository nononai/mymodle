package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ ComboChargeModel = (*customComboChargeModel)(nil)

type (
	// ComboChargeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customComboChargeModel.
	ComboChargeModel interface {
		comboChargeModel
		customComboChargeLogicModel
	}

	customComboChargeModel struct {
		*defaultComboChargeModel
	}

	customComboChargeLogicModel interface {
	}
)

// NewComboChargeModel returns a model for the database table.
func NewComboChargeModel(conn *gorm.DB, c cache.CacheConf) ComboChargeModel {
	return &customComboChargeModel{
		defaultComboChargeModel: newComboChargeModel(conn, c),
	}
}

func (m *defaultComboChargeModel) customCacheKeys(data *ComboCharge) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
