package models

import (
	"gorm.io/gorm"
)

// Material 物料基础信息
type Material struct {
	gorm.Model
	MaterialCode string `gorm:"uniqueIndex;not null;size:100" json:"material_code"` // 物料编码
	MaterialName string `gorm:"size:255" json:"material_name"`                     // 物料名称
	Spec         string `gorm:"size:255" json:"spec"`                              // 物料规格
	Unit         string `gorm:"size:50" json:"unit"`                               // 单位
	MaterialType string `gorm:"size:100" json:"material_type"`                     // 特殊分类/物料类型
	ABCClass     string `gorm:"size:10" json:"abc_class"`                          // ABC分类
}

// InboundOrder 入库单
type InboundOrder struct {
	gorm.Model
	InboundNo      string  `gorm:"uniqueIndex;not null;size:100" json:"inbound_no"` // 入库确认单编号
	IsUrgent       bool    `gorm:"default:false" json:"is_urgent"`                  // 是否加急
	ConfirmType    string  `gorm:"size:100" json:"confirm_type"`                    // 确认类型 (如: 来料入库确认)
	ReceiveNo      string  `gorm:"size:100" json:"receive_no"`                      // 收料单编码
	QCNo           string  `gorm:"size:100" json:"qc_no"`                           // 检验单编码
	MaterialCode   string  `gorm:"size:100" json:"material_code"`                   // 物料编码
	TotalQty       float64 `gorm:"type:decimal(18,4)" json:"total_qty"`             // 物料数量
	QualifiedQty   float64 `gorm:"type:decimal(18,4)" json:"qualified_qty"`         // 合格数量
	UnqualifiedQty float64 `gorm:"type:decimal(18,4)" json:"unqualified_qty"`       // 不合格数量
	IsQCConfirmed  bool    `gorm:"default:false" json:"is_qc_confirmed"`            // 已品质确认
	QCUser         string  `gorm:"size:100" json:"qc_user"`                         // 品质确认人
	QCTime         *string `json:"qc_time"`                                         // 品质确认时间
	IsInbound      bool    `gorm:"default:false" json:"is_inbound"`                 // 已入库确认
	InboundUser    string  `gorm:"size:100" json:"inbound_user"`                    // 入库确认人
	InboundTime    *string `json:"inbound_time"`                                    // 入库确认时间
	Source         string  `gorm:"size:100" json:"source"`                          // 来源 (如: 手动创建)
	OrgCode        string  `gorm:"size:100" json:"org_code"`                        // 组织编码
	OrgName        string  `gorm:"size:255" json:"org_name"`                        // 组织名称
	Remark         string  `gorm:"type:text" json:"remark"`                         // 备注
}

// Stock 库存表
type Stock struct {
	gorm.Model
	MaterialCode string  `gorm:"uniqueIndex;not null;size:100" json:"material_code"` // 物料编码
	Quantity     float64 `gorm:"type:decimal(18,4);default:0" json:"quantity"`       // 当前库存数量
	Warehouse    string  `gorm:"size:100" json:"warehouse"`                         // 仓库
}

// StockLog 库存流水
type StockLog struct {
	gorm.Model
	MaterialCode string  `gorm:"size:100" json:"material_code"`             // 物料编码
	Type         string  `gorm:"size:50" json:"type"`                       // 类型: inbound(入库), outbound(出库)
	Quantity     float64 `gorm:"type:decimal(18,4)" json:"quantity"`         // 变动数量
	BeforeQty    float64 `gorm:"type:decimal(18,4)" json:"before_qty"`       // 变动前数量
	AfterQty     float64 `gorm:"type:decimal(18,4)" json:"after_qty"`        // 变动后数量
	ReferenceNo  string  `gorm:"size:100" json:"reference_no"`              // 关联单据号
	Operator     string  `gorm:"size:100" json:"operator"`                  // 操作人
}
