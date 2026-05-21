package handlers

import (
	"mes-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetStockList 获取库存列表
func GetStockList(c *gin.Context) {
	var stocks []models.Stock
	if err := models.DB.Find(&stocks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取库存失败"})
		return
	}
	c.JSON(http.StatusOK, stocks)
}

// GetInboundOrders 获取入库单列表
func GetInboundOrders(c *gin.Context) {
	var orders []models.InboundOrder
	if err := models.DB.Order("created_at desc").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取入库单失败"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

// CreateInboundOrder 创建入库单 (模拟收料/来料过程)
func CreateInboundOrder(c *gin.Context) {
	var order models.InboundOrder
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	if err := models.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建入库单失败"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// ConfirmInbound 确认入库 (更新库存)
func ConfirmInbound(c *gin.Context) {
	id := c.Param("id")
	var order models.InboundOrder

	if err := models.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "入库单不存在"})
		return
	}

	if order.IsInbound {
		c.JSON(http.StatusBadRequest, gin.H{"error": "该单据已入库"})
		return
	}

	// 开启事务
	tx := models.DB.Begin()

	// 1. 更新入库单状态
	order.IsInbound = true
	username, _ := c.Get("username")
	order.InboundUser = username.(string)
	now := "2026-05-22 10:00:00" // 模拟当前时间
	order.InboundTime = &now

	if err := tx.Save(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "更新入库单失败"})
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
		// 新物料入库
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
	c.JSON(http.StatusOK, gin.H{"message": "入库确认成功"})
}
