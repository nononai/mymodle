package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ SiteChargeModel = (*customSiteChargeModel)(nil)

type (
	// SiteChargeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSiteChargeModel.
	SiteChargeModel interface {
		siteChargeModel
		customSiteChargeLogicModel
	}

	customSiteChargeModel struct {
		*defaultSiteChargeModel
	}

	customSiteChargeLogicModel interface {
	}
)

// NewSiteChargeModel returns a model for the database table.
func NewSiteChargeModel(conn *gorm.DB, c cache.CacheConf) SiteChargeModel {
	return &customSiteChargeModel{
		defaultSiteChargeModel: newSiteChargeModel(conn, c),
	}
}

func (m *defaultSiteChargeModel) customCacheKeys(data *SiteCharge) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
