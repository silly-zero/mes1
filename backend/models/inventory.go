package models

import (
	"gorm.io/gorm"
)

// Material 物料基础信息
type Material struct {
	gorm.Model
	MaterialCode string `gorm:"uniqueIndex;not null;size:100" json:"material_code"` // 物料编码
	MaterialName string `gorm:"size:255" json:"material_name"`                      // 物料名称
	Spec         string `gorm:"size:255" json:"spec"`                               // 物料规格
	Unit         string `gorm:"size:50" json:"unit"`                                // 单位
	MaterialType string `gorm:"size:100" json:"material_type"`                      // 特殊分类/物料类型
	ABCClass     string `gorm:"size:10" json:"abc_class"`                           // ABC分类
}

// InboundOrder 入库单 (包含收料、检验、入库全流程)
type InboundOrder struct {
	gorm.Model
	InboundNo        string  `gorm:"uniqueIndex;not null;size:100" json:"inbound_no"` // 入库确认单编号
	IsUrgent         bool    `gorm:"default:false" json:"is_urgent"`                  // 是否加急
	ConfirmType      string  `gorm:"size:100" json:"confirm_type"`                    // 确认类型 (如: 来料入库确认)
	ReceiveNo        string  `gorm:"size:100" json:"receive_no"`                      // 收料单编码
	ReceiptTypeCode  string  `gorm:"size:100" json:"receipt_type_code"`               // 收料类型编码
	ReceiptTypeName  string  `gorm:"size:100" json:"receipt_type_name"`               // 收料类型名称
	SupplierCode     string  `gorm:"size:100" json:"supplier_code"`                   // 供应商编码
	SupplierName     string  `gorm:"size:255" json:"supplier_name"`                   // 供应商名称
	CustomerCode     string  `gorm:"size:100" json:"customer_code"`                   // 客户编码
	CustomerName     string  `gorm:"size:255" json:"customer_name"`                   // 客户名称
	ASN              string  `gorm:"size:100" json:"asn"`                             // ASN
	ERPOrderTypeCode string  `gorm:"size:100" json:"erp_order_type_code"`             // ERP单据类型编码
	ERPOrderTypeName string  `gorm:"size:100" json:"erp_order_type_name"`             // ERP单据类型名称
	DeptCode         string  `gorm:"size:100" json:"dept_code"`                       // 部门编码
	DeptName         string  `gorm:"size:255" json:"dept_name"`                       // 部门名称
	ProjectCode      string  `gorm:"size:100" json:"project_code"`                    // 项目编码
	ProjectName      string  `gorm:"size:255" json:"project_name"`                    // 项目名称
	ProjectAbbr      string  `gorm:"size:100" json:"project_abbr"`                    // 项目简称
	QCNo             string  `gorm:"size:100" json:"qc_no"`                           // 检验单编码
	MaterialCode     string  `gorm:"size:100" json:"material_code"`                   // 物料编码
	TotalQty         float64 `gorm:"type:decimal(18,4)" json:"total_qty"`             // 物料数量
	QualifiedQty     float64 `gorm:"type:decimal(18,4)" json:"qualified_qty"`         // 合格数量
	UnqualifiedQty   float64 `gorm:"type:decimal(18,4)" json:"unqualified_qty"`       // 不合格数量
	IsQCConfirmed    bool    `gorm:"default:false" json:"is_qc_confirmed"`            // 已品质确认
	QCUser           string  `gorm:"size:100" json:"qc_user"`                         // 品质确认人
	QCTime           *string `json:"qc_time"`                                         // 品质确认时间
	IsInbound        bool    `gorm:"default:false" json:"is_inbound"`                 // 已入库确认
	InboundUser      string  `gorm:"size:100" json:"inbound_user"`                    // 入库确认人
	InboundTime      *string `json:"inbound_time"`                                    // 入库确认时间
	Status           string  `gorm:"size:50;default:'未完成'" json:"status"`             // 收料单状态
	Source           string  `gorm:"size:100" json:"source"`                          // 来源 (如: 手动创建)
	OrgCode          string  `gorm:"size:100" json:"org_code"`                        // 组织编码
	OrgName          string  `gorm:"size:255" json:"org_name"`                        // 组织名称
	Remark           string  `gorm:"type:text" json:"remark"`                         // 备注
	ModifierID       string  `gorm:"size:100" json:"modifier_id"`                     // 修改人ID
	ModifierName     string  `gorm:"size:100" json:"modifier_name"`                   // 修改人名称
	ModifyTime       *string `json:"modify_time"`                                     // 修改时间
}

// OutboundOrder 出库单
type OutboundOrder struct {
	gorm.Model
	OutboundNo   string  `gorm:"uniqueIndex;not null;size:100" json:"outbound_no"` // 出库单号
	MaterialCode string  `gorm:"not null;size:100" json:"material_code"`           // 物料编码
	Quantity     float64 `gorm:"type:decimal(18,4)" json:"quantity"`               // 出库数量
	Warehouse    string  `gorm:"size:100" json:"warehouse"`                        // 出库仓库
	Receiver     string  `gorm:"size:100" json:"receiver"`                         // 领料人/接收人
	Status       string  `gorm:"size:20;default:'pending'" json:"status"`          // 状态: pending, completed
	Operator     string  `gorm:"size:100" json:"operator"`                         // 操作人
	OutboundTime *string `json:"outbound_time"`                                    // 出库时间
}

// Stock 库存表
type Stock struct {
	gorm.Model
	MaterialCode string  `gorm:"uniqueIndex;not null;size:100" json:"material_code"` // 物料编码
	Quantity     float64 `gorm:"type:decimal(18,4);default:0" json:"quantity"`       // 当前库存数量
	Warehouse    string  `gorm:"size:100" json:"warehouse"`                          // 仓库
}

// StockLog 库存流水
type StockLog struct {
	gorm.Model
	MaterialCode string  `gorm:"size:100" json:"material_code"`        // 物料编码
	Type         string  `gorm:"size:50" json:"type"`                  // 类型: inbound(入库), outbound(出库)
	Quantity     float64 `gorm:"type:decimal(18,4)" json:"quantity"`   // 变动数量
	BeforeQty    float64 `gorm:"type:decimal(18,4)" json:"before_qty"` // 变动前数量
	AfterQty     float64 `gorm:"type:decimal(18,4)" json:"after_qty"`  // 变动后数量
	ReferenceNo  string  `gorm:"size:100" json:"reference_no"`         // 关联单据号
	Operator     string  `gorm:"size:100" json:"operator"`             // 操作人
}
