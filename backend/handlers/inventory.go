package handlers

import (
	"mes-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// --- 库存管理 ---

// GetStockList 获取库存列表
func GetStockList(c *gin.Context) {
	var stocks []models.Stock
	if err := models.DB.Find(&stocks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取库存失败"})
		return
	}
	c.JSON(http.StatusOK, stocks)
}

// --- 收料管理 ---

// CreateReceipt 创建收料记录 (对应图片中的入库确认单初始化)
func CreateReceipt(c *gin.Context) {
	var order models.InboundOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	// 初始状态
	order.IsQCConfirmed = false
	order.IsInbound = false

	if err := models.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建收料记录失败"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// --- 来料检验 (IQC) ---

// ConfirmIQC 品质确认 (IQC)
func ConfirmIQC(c *gin.Context) {
	id := c.Param("id")
	type IQCReq struct {
		QCNo           string  `json:"qc_no" binding:"required"`
		QualifiedQty   float64 `json:"qualified_qty"`
		UnqualifiedQty float64 `json:"unqualified_qty"`
	}

	var req IQCReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	var order models.InboundOrder
	if err := models.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "单据不存在"})
		return
	}

	username, _ := c.Get("username")
	now := time.Now().Format("2006-01-02 15:04:05")

	order.QCNo = req.QCNo
	order.QualifiedQty = req.QualifiedQty
	order.UnqualifiedQty = req.UnqualifiedQty
	order.IsQCConfirmed = true
	order.QCUser = username.(string)
	order.QCTime = &now

	if err := models.DB.Save(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "品质确认失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "品质确认成功", "data": order})
}

// --- 入库管理 ---

// GetInboundOrders 获取入库单列表
func GetInboundOrders(c *gin.Context) {
	var orders []models.InboundOrder
	if err := models.DB.Order("created_at desc").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取入库单失败"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// ConfirmInbound 确认入库
func ConfirmInbound(c *gin.Context) {
	id := c.Param("id")
	var order models.InboundOrder

	if err := models.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "入库单不存在"})
		return
	}

	if !order.IsQCConfirmed {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请先进行品质确认"})
		return
	}

	if order.IsInbound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该单据已入库"})
		return
	}

	tx := models.DB.Begin()

	// 1. 更新入库状态
	username, _ := c.Get("username")
	now := time.Now().Format("2006-01-02 15:04:05")
	order.IsInbound = true
	order.InboundUser = username.(string)
	order.InboundTime = &now

	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新入库状态失败"})
		return
	}

	// 2. 更新库存
	var stock models.Stock
	err := tx.Where("material_code = ?", order.MaterialCode).First(&stock).Error

	beforeQty := 0.0
	if err == nil {
		beforeQty = stock.Quantity
		stock.Quantity += order.QualifiedQty
		if err := tx.Save(&stock).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "更新库存失败"})
			return
		}
	} else {
		stock = models.Stock{
			MaterialCode: order.MaterialCode,
			Quantity:     order.QualifiedQty,
			Warehouse:    "默认仓库",
		}
		if err := tx.Create(&stock).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "创建库存记录失败"})
			return
		}
	}

	// 3. 记录流水
	log := models.StockLog{
		MaterialCode: order.MaterialCode,
		Type:         "inbound",
		Quantity:     order.QualifiedQty,
		BeforeQty:    beforeQty,
		AfterQty:     beforeQty + order.QualifiedQty,
		ReferenceNo:  order.InboundNo,
		Operator:     username.(string),
	}
	if err := tx.Create(&log).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "记录库存流水失败"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "入库成功"})
}

// --- 出库管理 ---

// CreateOutboundOrder 创建并执行出库
func CreateOutboundOrder(c *gin.Context) {
	var req models.OutboundOrder
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	tx := models.DB.Begin()

	// 1. 检查库存
	var stock models.Stock
	if err := tx.Where("material_code = ?", req.MaterialCode).First(&stock).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "该物料无库存"})
		return
	}

	if stock.Quantity < req.Quantity {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "库存不足"})
		return
	}

	// 2. 扣减库存
	beforeQty := stock.Quantity
	stock.Quantity -= req.Quantity
	if err := tx.Save(&stock).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "库存扣减失败"})
		return
	}

	// 3. 创建出库单
	username, _ := c.Get("username")
	now := time.Now().Format("2006-01-02 15:04:05")
	req.Operator = username.(string)
	req.OutboundTime = &now
	req.Status = "completed"

	if err := tx.Create(&req).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建出库单失败"})
		return
	}

	// 4. 记录流水
	log := models.StockLog{
		MaterialCode: req.MaterialCode,
		Type:         "outbound",
		Quantity:     req.Quantity,
		BeforeQty:    beforeQty,
		AfterQty:     stock.Quantity,
		ReferenceNo:  req.OutboundNo,
		Operator:     username.(string),
	}
	if err := tx.Create(&log).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "记录库存流水失败"})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"message": "出库成功"})
}

// GetOutboundOrders 获取出库单列表
func GetOutboundOrders(c *gin.Context) {
	var orders []models.OutboundOrder
	if err := models.DB.Order("created_at desc").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取出库单失败"})
		return
	}
	c.JSON(http.StatusOK, orders)
}
