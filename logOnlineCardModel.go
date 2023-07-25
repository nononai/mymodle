package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ LogOnlineCardModel = (*customLogOnlineCardModel)(nil)

type (
	// LogOnlineCardModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLogOnlineCardModel.
	LogOnlineCardModel interface {
		logOnlineCardModel
		customLogOnlineCardLogicModel
	}

	customLogOnlineCardModel struct {
		*defaultLogOnlineCardModel
	}

	customLogOnlineCardLogicModel interface {
	}
)

// NewLogOnlineCardModel returns a model for the database table.
func NewLogOnlineCardModel(conn *gorm.DB, c cache.CacheConf) LogOnlineCardModel {
	return &customLogOnlineCardModel{
		defaultLogOnlineCardModel: newLogOnlineCardModel(conn, c),
	}
}

func (m *defaultLogOnlineCardModel) customCacheKeys(data *LogOnlineCard) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
