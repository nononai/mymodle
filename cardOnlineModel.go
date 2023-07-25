package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ CardOnlineModel = (*customCardOnlineModel)(nil)

type (
	// CardOnlineModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCardOnlineModel.
	CardOnlineModel interface {
		cardOnlineModel
		customCardOnlineLogicModel
	}

	customCardOnlineModel struct {
		*defaultCardOnlineModel
	}

	customCardOnlineLogicModel interface {
		GetCardOnlineBySiteIdAndCardType(ctx context.Context, siteId int64, cardType int64) (*CardOnline, error)
	}
)

// NewCardOnlineModel returns a model for the database table.
func NewCardOnlineModel(conn *gorm.DB, c cache.CacheConf) CardOnlineModel {
	return &customCardOnlineModel{
		defaultCardOnlineModel: newCardOnlineModel(conn, c),
	}
}

func (m *defaultCardOnlineModel) customCacheKeys(data *CardOnline) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

// 通过站点id和卡类型获取卡信息
func (m *defaultCardOnlineModel) GetCardOnlineBySiteIdAndCardType(ctx context.Context, siteId int64, cardType int64) (*CardOnline, error) {
	var resp CardOnline
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&CardOnline{}).Where("`site_id` = ? AND `card_type` = ?", siteId, cardType).Preload("Site").First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
