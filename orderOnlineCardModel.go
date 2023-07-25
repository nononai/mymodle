package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ OrderOnlineCardModel = (*customOrderOnlineCardModel)(nil)

type (
	// OrderOnlineCardModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderOnlineCardModel.
	OrderOnlineCardModel interface {
		orderOnlineCardModel
		customOrderOnlineCardLogicModel
	}

	customOrderOnlineCardModel struct {
		*defaultOrderOnlineCardModel
	}

	customOrderOnlineCardLogicModel interface {
	}
)

// NewOrderOnlineCardModel returns a model for the database table.
func NewOrderOnlineCardModel(conn *gorm.DB, c cache.CacheConf) OrderOnlineCardModel {
	return &customOrderOnlineCardModel{
		defaultOrderOnlineCardModel: newOrderOnlineCardModel(conn, c),
	}
}

func (m *defaultOrderOnlineCardModel) customCacheKeys(data *OrderOnlineCard) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
