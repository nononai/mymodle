package model

import (
	"context"
	"fmt"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ DeviceChargePortModel = (*customDeviceChargePortModel)(nil)
var (
	cacheMinyunDeviceCharegePoryQrcode = "cache:minyun:devicePort:qrcode:"
)

type (
	// DeviceChargePortModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDeviceChargePortModel.
	DeviceChargePortModel interface {
		deviceChargePortModel
		customDeviceChargePortLogicModel
	}

	customDeviceChargePortModel struct {
		*defaultDeviceChargePortModel
	}

	customDeviceChargePortLogicModel interface {
		GetDeviceChargePortsByDeviceId(ctx context.Context, deviceId int64) ([]*DeviceChargePort, error)
		GetDeviceChargePortByQrcode(ctx context.Context, devicePortQrcode string) (*DeviceChargePort, error)
		UpdateDeviceChargePortStatusByDeviceIdAndPort(ctx context.Context, deviceId int64, port int64, status int64) error
		GetDeviceChargePortByDeviceIdAndPort(ctx context.Context, deviceId int64, port int64) (*DeviceChargePort, error)
	}
)

// NewDeviceChargePortModel returns a model for the database table.
func NewDeviceChargePortModel(conn *gorm.DB, c cache.CacheConf) DeviceChargePortModel {
	return &customDeviceChargePortModel{
		defaultDeviceChargePortModel: newDeviceChargePortModel(conn, c),
	}
}

func (m *defaultDeviceChargePortModel) customCacheKeys(data *DeviceChargePort) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
func (m *defaultDeviceChargePortModel) GetDeviceChargePortsByDeviceId(ctx context.Context, deviceId int64) ([]*DeviceChargePort, error) {
	var resp []*DeviceChargePort
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&DeviceChargePort{}).Where("`device_id` = ?", deviceId).Find(&resp).Error
	})
	switch err {
	case nil:
		return resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err

	}
}
func (m *defaultDeviceChargePortModel) GetDeviceChargePortByQrcode(ctx context.Context, devicePortQrcode string) (*DeviceChargePort, error) {
	cacheMinyunDeviceCharegePoryQrcode := fmt.Sprintf("%s%v", cacheMinyunDeviceCharegePoryQrcode, devicePortQrcode)
	var resp DeviceChargePort
	err := m.QueryCtx(ctx, &resp, cacheMinyunDeviceCharegePoryQrcode, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&DeviceChargePort{}).Where("`qrcode` = ?", devicePortQrcode).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err

	}
}

// 根据设备id和端口号修改端口状态
func (m *defaultDeviceChargePortModel) UpdateDeviceChargePortStatusByDeviceIdAndPort(ctx context.Context, deviceId int64, port int64, status int64) error {
	return m.ExecCtx(ctx, func(conn *gorm.DB) error {
		return conn.Model(&DeviceChargePort{}).Where("`device_id` = ? and `port_id` = ?", deviceId, port).Update("status", status).Error
	})
}

// 根据设备id和端口号查找端口信息
func (m *defaultDeviceChargePortModel) GetDeviceChargePortByDeviceIdAndPort(ctx context.Context, deviceId int64, port int64) (*DeviceChargePort, error) {
	var resp DeviceChargePort
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&DeviceChargePort{}).Where("`device_id` = ? and `port_id` = ?", deviceId, port).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err

	}
}
