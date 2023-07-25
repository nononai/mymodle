package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ ThirdPaymentModel = (*customThirdPaymentModel)(nil)

type (
	// ThirdPaymentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customThirdPaymentModel.
	ThirdPaymentModel interface {
		thirdPaymentModel
		customThirdPaymentLogicModel
	}

	customThirdPaymentModel struct {
		*defaultThirdPaymentModel
	}

	customThirdPaymentLogicModel interface {
	}
)

// NewThirdPaymentModel returns a model for the database table.
func NewThirdPaymentModel(conn *gorm.DB, c cache.CacheConf) ThirdPaymentModel {
	return &customThirdPaymentModel{
		defaultThirdPaymentModel: newThirdPaymentModel(conn, c),
	}
}

func (m *defaultThirdPaymentModel) customCacheKeys(data *ThirdPayment) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
