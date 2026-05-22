package models

import (
	"gorm.io/gorm"
	"time"
)

// MSDRecord 湿敏物料记录
type MSDRecord struct {
	gorm.Model
	MaterialCode     string     `gorm:"size:100;not null" json:"material_code"`      // 物料编码
	LotNo            string     `gorm:"size:100;uniqueIndex;not null" json:"lot_no"` // 批次号/卷盘号
	MSLLevel         string     `gorm:"size:20" json:"msl_level"`                   // MSL等级 (1, 2, 2a, 3, 4, 5, 5a, 6)
	TotalFloorLife   float64    `json:"total_floor_life"`                           // 总车间寿命 (小时)
	RemainingLife    float64    `json:"remaining_life"`                             // 剩余寿命 (小时)
	Status           string     `gorm:"size:50;default:'InBag'" json:"status"`      // 状态: InBag(袋装), Exposed(暴露), Baking(烘烤), Expired(过期)
	LastOpenTime     *time.Time `json:"last_open_time"`                             // 最近一次开封时间
	LastCloseTime    *time.Time `json:"last_close_time"`                            // 最近一次封袋/开始烘烤时间
	TotalExposedTime float64    `json:"total_exposed_time"`                         // 累计暴露时间 (小时)
	Warehouse        string     `gorm:"size:100" json:"warehouse"`                  // 所在仓库/位置
	Operator         string     `gorm:"size:100" json:"operator"`                   // 最后操作人
}

// MSLLifeMap 定义不同MSL等级在标准条件下的车间寿命 (小时)
var MSLLifeMap = map[string]float64{
	"1":  -1,   // 无限制
	"2":  8760, // 1年
	"2a": 672,  // 4周
	"3":  168,  // 168小时 (1周)
	"4":  72,   // 72小时
	"5":  48,   // 48小时
	"5a": 24,   // 24小时
	"6":  6,    // 强制烘烤后使用
}
