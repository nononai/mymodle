package model

import (
	"context"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ UserBasicModel = (*customUserBasicModel)(nil)

// var (
// 	cacheMinyunUserBasicOpenIdPrefix     = "cache:minyun:OpenId:openid:"
// 	cacheMinyunUserBasicUnionIdPrefix    = "cache:minyun:UnionId:unionid:"
// 	cacheMinyunUserBaseicAliOpenIdPrefix = "cache:minyun:AliOpenId:aliopenid:"
// )

type (
	// UserBasicModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserBasicModel.
	UserBasicModel interface {
		userBasicModel
		customUserBasicLogicModel
	}

	customUserBasicModel struct {
		*defaultUserBasicModel
	}

	customUserBasicLogicModel interface {
		GetUserBasicByOpenId(ctx context.Context, openId string) (*UserBasic, error)
		GetUserBasicByUnionId(ctx context.Context, unionId string) (*UserBasic, error)
		GetUserBasicByAliOpenId(ctx context.Context, aliOpenId string) (*UserBasic, error)
		GetUserBasicByMpOpenId(ctx context.Context, mpOpenId string) (*UserBasic, error)
	}
)

// NewUserBasicModel returns a model for the database table.
func NewUserBasicModel(conn *gorm.DB, c cache.CacheConf) UserBasicModel {
	return &customUserBasicModel{
		defaultUserBasicModel: newUserBasicModel(conn, c),
	}
}

func (m *defaultUserBasicModel) customCacheKeys(data *UserBasic) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
func (m *defaultUserBasicModel) GetUserBasicByOpenId(ctx context.Context, openId string) (*UserBasic, error) {
	var resp UserBasic
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&UserBasic{}).Where("`applet_openid` = ?", openId).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound

	default:
		return nil, err
	}
}
func (m *defaultUserBasicModel) GetUserBasicByUnionId(ctx context.Context, unionId string) (*UserBasic, error) {
	// minyunUserBasicUnionIdKey := fmt.Sprintf("%s%v", cacheMinyunUserBasicUnionIdPrefix, unionId)
	var resp UserBasic
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&UserBasic{}).Where("`union_id` = ?", unionId).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultUserBasicModel) GetUserBasicByAliOpenId(ctx context.Context, aliOpenId string) (*UserBasic, error) {
	var resp UserBasic
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&UserBasic{}).Where("`ali_openid` = ?", aliOpenId).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// 通过用户的mp_openid获取用户的基本信息
func (m *defaultUserBasicModel) GetUserBasicByMpOpenId(ctx context.Context, mpOpenId string) (*UserBasic, error) {
	var resp UserBasic
	err := m.QueryNoCacheCtx(ctx, &resp, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&UserBasic{}).Where("`mp_openid` = ?", mpOpenId).First(&resp).Error
	})
	switch err {
	case nil:
		return &resp, nil
	case gormc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
