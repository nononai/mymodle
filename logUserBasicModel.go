package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ LogUserBasicModel = (*customLogUserBasicModel)(nil)

type (
	// LogUserBasicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLogUserBasicModel.
	LogUserBasicModel interface {
		logUserBasicModel
		customLogUserBasicLogicModel
	}

	customLogUserBasicModel struct {
		*defaultLogUserBasicModel
	}

	customLogUserBasicLogicModel interface {
	}
)

// NewLogUserBasicModel returns a model for the database table.
func NewLogUserBasicModel(conn *gorm.DB, c cache.CacheConf) LogUserBasicModel {
	return &customLogUserBasicModel{
		defaultLogUserBasicModel: newLogUserBasicModel(conn, c),
	}
}

func (m *defaultLogUserBasicModel) customCacheKeys(data *LogUserBasic) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
