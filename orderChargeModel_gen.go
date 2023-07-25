// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"fmt"
	"time"

	"github.com/SpectatorNan/gorm-zero/gormc"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var (
	cacheMinyunOrderChargeIdPrefix = "cache:minyun:orderCharge:id:"
)

type (
	orderChargeModel interface {
		Insert(ctx context.Context, tx *gorm.DB, data *OrderCharge) error

		FindOne(ctx context.Context, id int64) (*OrderCharge, error)
		Update(ctx context.Context, tx *gorm.DB, data *OrderCharge) error

		Delete(ctx context.Context, tx *gorm.DB, id int64) error
		Transaction(ctx context.Context, fn func(db *gorm.DB) error) error
	}

	defaultOrderChargeModel struct {
		gormc.CachedConn
		table string
	}

	OrderCharge struct {
		Id            int64          `gorm:"column:id"`
		ChannelId     int64          `gorm:"column:channel_id"`
		ShopId        int64          `gorm:"column:shop_id"`
		SiteId        int64          `gorm:"column:site_id"`
		AreaId        int64          `gorm:"column:area_id"`
		UserId        int64          `gorm:"column:user_id"`
		DeviceId      int64          `gorm:"column:device_id"`
		Qrcode        string         `gorm:"column:qrcode"`
		OrderPirce    int64          `gorm:"column:order_pirce"` // 订单原价格
		PayMoney      int64          `gorm:"column:pay_money"`   // 实际支付金额
		PayStat       int64          `gorm:"column:pay_stat"`    // 支付状态，1已支付，2，已退款，3，部分退款
		PayType       int64          `gorm:"column:pay_type"`    // 支付方式，1微信，2支付宝，3储值卡，4次数卡，5，时长卡，6，电量卡，7运维，8白名单，9刷卡支付，10线下充电卡
		PayTime       int64          `gorm:"column:pay_time"`
		CtTime        int64          `gorm:"column:ct_time"`
		DurationTime  int64          `gorm:"column:duration_time"` // 预设时间
		TimeType      int64          `gorm:"column:time_type"`     // 00计时，=01包月，=02计量，=03计次。包月、计次默认充满自停，计时、计量可手动设置时长和电量(只有1和2 其他描述为友电的下发格式)
		OrderId       string         `gorm:"column:order_id"`      // 订单号订单号（32位纯数字）订单号生产规则参考：前面14位是年月日时分秒中间8位是随机数后面10位是10进制设备号+端口号 如：20200615113914002978611105276501
		OutNo         string         `gorm:"column:out_no"`        // 外部订单号
		PortId        int64          `gorm:"column:port_id"`
		RefundMoney   int64          `gorm:"column:refund_money"`
		RefundId      string         `gorm:"column:refund_id"`
		CheckPower    int64          `gorm:"column:check_power"`  // 结算功率
		TotalEnergy   float64        `gorm:"column:total_energy"` // 用电量
		CostMoney     int64          `gorm:"column:cost_money"`   // 消费金额
		ServerFee     int64          `gorm:"column:server_fee"`   // 服务费或占位费
		EndTime       int64          `gorm:"column:end_time"`
		CheckFee      int64          `gorm:"column:check_fee"`       // 结算功率档次费用
		IsLong        int64          `gorm:"column:is_long"`         // 是否长冲模式
		EndReason     string         `gorm:"column:end_reason"`      // 停止原因 01(充满自停) 02(达到最大充电时间) 03(达到预设时间) 04（达到预设电量）05(用户拔出) 06(负载过大)，07(服务器控制停止) 08（动态过载） 09（功率过小） 0A（环境温度过高） 0B（端口温度过高） 0C（过流） 0D（用户拔出-1，可能是插座弹片卡住） 0E（无功率停止，可能是接触不良或保险丝烧断故障），0F（预检-继电器坏或保险丝断），注：0A仅适用于AP262、AP360； 0B仅适用于AP262，10=水浸断电，11=灭火结算（本端口），12=灭火结算（非本端口），13=用户密码开柜断电，14=未关好柜门，15=外部操作打开柜门
		OfflineCardId string         `gorm:"column:offline_card_id"` // 线下卡开启有效，卡id
		MonthCardId   int64          `gorm:"column:month_card_id"`   // 包月卡开启有效，卡id
		PresetEnergy  int64          `gorm:"column:preset_energy"`   // 预设电量
		ErrStep       int64          `gorm:"column:err_step"`        // 异常订单：0、未申报，1、已申报，2、已回复，3、已处理
		OrderFrom     int64          `gorm:"column:order_from"`      // 订单来源平台 1，微信小程序 2支付宝小程序 3 h5端 4，app端
		IsCheck       int64          `gorm:"column:is_check"`        // 是否已结算
		CreatedAt     time.Time      `gorm:"column:created_at"`
		UpdatedAt     time.Time      `gorm:"column:updated_at"`
		DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index"`
		MaxPower      int64          `gorm:"column:max_power"`       // 订单最大功率限制
		MaxTimes      int64          `gorm:"column:max_times"`       // 订单最大充电时长
		BalanceCardId int64          `gorm:"column:balance_card_id"` // 用户储值卡id
		ChargeStat    int64          `gorm:"column:charge_stat"`     // 充电状态1，开始充电，2结束充电3，充电开启失败
		ChargeLog     string         `gorm:"column:charge_log"`      // 充电电流记录，json
		CostGive      int64          `gorm:"column:cost_give"`       // 订单消费赠送金额
		CostBasic     int64          `gorm:"column:cost_basic"`      // 订单消费基本金额
		ChargeTime    int64          `gorm:"column:charge_time"`     // 实际充电时长
		OrderPower    int64          `gorm:"column:order_power"`     // 本订单此次功率
		Area AreaCharge `gorm:"foreignKey:AreaId`
		Site SiteCharge `gorm:"foreignKey:SiteId`
		Device DeviceCharge `gorm:"foreignKey:DeviceId`
	}
)

func (OrderCharge) TableName() string {
	return "`order_charge`"
}

func newOrderChargeModel(conn *gorm.DB, c cache.CacheConf) *defaultOrderChargeModel {
	return &defaultOrderChargeModel{
		CachedConn: gormc.NewConn(conn, c),
		table:      "`order_charge`",
	}
}

func (m *defaultOrderChargeModel) Insert(ctx context.Context, tx *gorm.DB, data *OrderCharge) error {

	err := m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(&data).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultOrderChargeModel) FindOne(ctx context.Context, id int64) (*OrderCharge, error) {
	minyunOrderChargeIdKey := fmt.Sprintf("%s%v", cacheMinyunOrderChargeIdPrefix, id)
	var resp OrderCharge
	err := m.QueryCtx(ctx, &resp, minyunOrderChargeIdKey, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&OrderCharge{}).Where("`id` = ?", id).First(&resp).Error
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

func (m *defaultOrderChargeModel) Update(ctx context.Context, tx *gorm.DB, data *OrderCharge) error {
	old, err := m.FindOne(ctx, data.Id)
	if err != nil && err != ErrNotFound {
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Save(data).Error
	}, m.getCacheKeys(old)...)
	return err
}

func (m *defaultOrderChargeModel) getCacheKeys(data *OrderCharge) []string {
	if data == nil {
		return []string{}
	}
	minyunOrderChargeIdKey := fmt.Sprintf("%s%v", cacheMinyunOrderChargeIdPrefix, data.Id)
	cacheKeys := []string{
		minyunOrderChargeIdKey,
	}
	cacheKeys = append(cacheKeys, m.customCacheKeys(data)...)
	return cacheKeys
}

func (m *defaultOrderChargeModel) Delete(ctx context.Context, tx *gorm.DB, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		if err == ErrNotFound {
			return nil
		}
		return err
	}
	err = m.ExecCtx(ctx, func(conn *gorm.DB) error {
		db := conn
		if tx != nil {
			db = tx
		}
		return db.Delete(&OrderCharge{}, id).Error
	}, m.getCacheKeys(data)...)
	return err
}

func (m *defaultOrderChargeModel) Transaction(ctx context.Context, fn func(db *gorm.DB) error) error {
	return m.TransactCtx(ctx, fn)
}

func (m *defaultOrderChargeModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMinyunOrderChargeIdPrefix, primary)
}

func (m *defaultOrderChargeModel) queryPrimary(conn *gorm.DB, v, primary interface{}) error {
	return conn.Model(&OrderCharge{}).Where("`id` = ?", primary).Take(v).Error
}
