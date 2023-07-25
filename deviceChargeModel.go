package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ DeviceChargeModel = (*customDeviceChargeModel)(nil)
var (
	cacheMinyunDeviceChargesn     = "cache:minyun:deviceCharge:sn:"
	cacheMinyunDeviceChargeQrcode = "cache:minyun:deviceCharge:qrcode:"
)

type (
	// DeviceChargeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeviceChargeModel.
	DeviceChargeModel interface {
		deviceChargeModel
		customDeviceChargeLogicModel
	}

	customDeviceChargeModel struct {
		*defaultDeviceChargeModel
	}

	customDeviceChargeLogicModel interface {
		GetDeviceChargeByDeviceId(ctx context.Context, deviceId string) (*DeviceCharge, error)
		GetDeviceChargeByDeviceQrcode(ctx context.Context, deviceQrcode string) (*DeviceCharge, error)
	}
)

// NewDeviceChargeModel returns a model for the database table.
func NewDeviceChargeModel(conn *gorm.DB, c cache.CacheConf) DeviceChargeModel {
	return &customDeviceChargeModel{
		defaultDeviceChargeModel: newDeviceChargeModel(conn, c),
	}
}

func (m *defaultDeviceChargeModel) customCacheKeys(data *DeviceCharge) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

// 按设备id查询（sn）查找设备
func (m *defaultDeviceChargeModel) GetDeviceChargeByDeviceId(ctx context.Context, deviceId string) (*DeviceCharge, error) {
	minyunDeviceChargeDevicesnKey := fmt.Sprintf("%s%v", cacheMinyunDeviceChargesn, deviceId)
	var resp DeviceCharge
	err := m.QueryCtx(ctx, &resp, minyunDeviceChargeDevicesnKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&DeviceCharge{}).Where("`device_id` = ?", deviceId).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

// 按设备qrcode查找设备信息
func (m *defaultDeviceChargeModel) GetDeviceChargeByDeviceQrcode(ctx context.Context, deviceQrcode string) (*DeviceCharge, error) {
	var resp DeviceCharge
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&DeviceCharge{}).Where("`qrcode` = ?", deviceQrcode).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err

	}
}
