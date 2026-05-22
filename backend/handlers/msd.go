package handlers

import (
	"mes-backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetMSDList 获取湿敏物料列表
func GetMSDList(c *gin.Context) {
	var records []models.MSDRecord
	if err := models.DB.Order("updated_at desc").Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取湿敏物料列表失败"})
		return
	}

	// 动态计算暴露时间（如果当前处于 Exposed 状态）
	now := time.Now()
	for i := range records {
		if records[i].Status == "Exposed" && records[i].LastOpenTime != nil {
			exposureHours := now.Sub(*records[i].LastOpenTime).Hours()
			records[i].RemainingLife -= exposureHours
			if records[i].RemainingLife < 0 {
				records[i].RemainingLife = 0
				records[i].Status = "Expired"
			}
		}
	}

	c.JSON(http.StatusOK, records)
}

// CreateMSDRecord 创建记录
func CreateMSDRecord(c *gin.Context) {
	var record models.MSDRecord
	if err := c.ShouldBindJSON(&record); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数错误"})
		return
	}

	// 根据 MSL 等级设置初始寿命
	if life, ok := models.MSLLifeMap[record.MSLLevel]; ok {
		record.TotalFloorLife = life
		record.RemainingLife = life
	}

	if err := models.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建记录失败"})
		return
	}

	c.JSON(http.StatusOK, record)
}

// OpenMSDBag 开封操作
func OpenMSDBag(c *gin.Context) {
	id := c.Param("id")
	var record models.MSDRecord
	if err := models.DB.First(&record, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "记录不存在"})
		return
	}

	if record.Status != "InBag" && record.Status != "Baking" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "当前状态不可开封"})
		return
	}

	now := time.Now()
	record.Status = "Exposed"
	record.LastOpenTime = &now
	username, _ := c.Get("username")
	record.Operator = username.(string)

	models.DB.Save(&record)
	c.JSON(http.StatusOK, record)
}

// CloseMSDBag 封袋操作
func CloseMSDBag(c *gin.Context) {
	id := c.Param("id")
	var record models.MSDRecord
	if err := models.DB.First(&record, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "记录不存在"})
		return
	}

	if record.Status != "Exposed" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "当前状态不可封袋"})
		return
	}

	now := time.Now()
	exposureHours := now.Sub(*record.LastOpenTime).Hours()
	
	record.RemainingLife -= exposureHours
	record.TotalExposedTime += exposureHours
	record.Status = "InBag"
	record.LastCloseTime = &now
	if record.RemainingLife <= 0 {
		record.RemainingLife = 0
		record.Status = "Expired"
	}

	username, _ := c.Get("username")
	record.Operator = username.(string)

	models.DB.Save(&record)
	c.JSON(http.StatusOK, record)
}

// StartBaking 开始烘烤
func StartBaking(c *gin.Context) {
	id := c.Param("id")
	var record models.MSDRecord
	if err := models.DB.First(&record, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "记录不存在"})
		return
	}

	now := time.Now()
	record.Status = "Baking"
	record.LastCloseTime = &now
	username, _ := c.Get("username")
	record.Operator = username.(string)

	models.DB.Save(&record)
	c.JSON(http.StatusOK, record)
}
