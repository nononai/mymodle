package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ DeviceChargeLogModel = (*customDeviceChargeLogModel)(nil)

type (
	// DeviceChargeLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeviceChargeLogModel.
	DeviceChargeLogModel interface {
		deviceChargeLogModel
		customDeviceChargeLogLogicModel
	}

	customDeviceChargeLogModel struct {
		*defaultDeviceChargeLogModel
	}

	customDeviceChargeLogLogicModel interface {
	}
)

// NewDeviceChargeLogModel returns a model for the database table.
func NewDeviceChargeLogModel(conn *gorm.DB, c cache.CacheConf) DeviceChargeLogModel {
	return &customDeviceChargeLogModel{
		defaultDeviceChargeLogModel: newDeviceChargeLogModel(conn, c),
	}
}

func (m *defaultDeviceChargeLogModel) customCacheKeys(data *DeviceChargeLog) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
