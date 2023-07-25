package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ ChannelPaymentModel = (*customChannelPaymentModel)(nil)

type (
	// ChannelPaymentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChannelPaymentModel.
	ChannelPaymentModel interface {
		channelPaymentModel
		customChannelPaymentLogicModel
	}

	customChannelPaymentModel struct {
		*defaultChannelPaymentModel
	}

	customChannelPaymentLogicModel interface {
	}
)

// NewChannelPaymentModel returns a model for the database table.
func NewChannelPaymentModel(conn *gorm.DB, c cache.CacheConf) ChannelPaymentModel {
	return &customChannelPaymentModel{
		defaultChannelPaymentModel: newChannelPaymentModel(conn, c),
	}
}

func (m *defaultChannelPaymentModel) customCacheKeys(data *ChannelPayment) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
