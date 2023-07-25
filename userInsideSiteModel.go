package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UserInsideSiteModel = (*customUserInsideSiteModel)(nil)

type (
	// UserInsideSiteModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserInsideSiteModel.
	UserInsideSiteModel interface {
		userInsideSiteModel
		customUserInsideSiteLogicModel
	}

	customUserInsideSiteModel struct {
		*defaultUserInsideSiteModel
	}

	customUserInsideSiteLogicModel interface {
	}
)

// NewUserInsideSiteModel returns a model for the database table.
func NewUserInsideSiteModel(conn *gorm.DB, c cache.CacheConf) UserInsideSiteModel {
	return &customUserInsideSiteModel{
		defaultUserInsideSiteModel: newUserInsideSiteModel(conn, c),
	}
}

func (m *defaultUserInsideSiteModel) customCacheKeys(data *UserInsideSite) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
