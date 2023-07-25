package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ OrderChargeModel = (*customOrderChargeModel)(nil)

type (
	// OrderChargeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOrderChargeModel.
	OrderChargeModel interface {
		orderChargeModel
		customOrderChargeLogicModel
	}

	customOrderChargeModel struct {
		*defaultOrderChargeModel
	}

	customOrderChargeLogicModel interface {
		GetOrderChargeByOrderId(ctx context.Context, orderId string) (*OrderCharge, error)
		GetOrderChargeByStatusAndUserId(ctx context.Context, userId, status int, page, pageSize int) ([]*OrderCharge, error)
	}
)

// NewOrderChargeModel returns a model for the database table.
func NewOrderChargeModel(conn *gorm.DB, c cache.CacheConf) OrderChargeModel {
	return &customOrderChargeModel{
		defaultOrderChargeModel: newOrderChargeModel(conn, c),
	}
}

func (m *defaultOrderChargeModel) customCacheKeys(data *OrderCharge) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
func (m *defaultOrderChargeModel) GetOrderChargeByOrderId(ctx context.Context, orderId string) (*OrderCharge, error) {
	var resp OrderCharge
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&OrderCharge{}).Where("order_id = ?", orderId).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, err
	}
}

// 通过订单状态和用户id分页获取充电订单信息
func (m *defaultOrderChargeModel) GetOrderChargeByStatusAndUserId(ctx context.Context, userId, status int, page, pageSize int) ([]*OrderCharge, error) {
	var resp []*OrderCharge
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&OrderCharge{}).Where("charge_stat = ? and user_id=?", status, userId).Offset((page - 1) * pageSize).
			Limit(pageSize).Preload("Device").Preload("Site").Preload("Area").Find(&resp).Error
	})
	switch err {
	case nil:
		return resp, nil
	case gorm.ErrRecordNotFound:
		return nil, nil
	default:
		return nil, err
	}
}
