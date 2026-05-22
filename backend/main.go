package main

import (
	"fmt"
	"mes-backend/config"
	"mes-backend/handlers"
	"mes-backend/middleware"
	"mes-backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 初始化配置
	config.InitConfig()

	// 2. 初始化数据库
	models.InitDB()

	// 3. 设置路由
	r := gin.Default()

	// 公共路由
	api := r.Group("/api")
	{
		api.POST("/login", handlers.Login)
		api.POST("/register", handlers.Register) // 临时用于初始化账号
	}

	// 权限路由示例
	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		// 库存管理相关
		inventory := auth.Group("/inventory")
		{
			inventory.GET("/stock", handlers.GetStockList)      // 库存列表
			inventory.GET("/inbound", handlers.GetInboundOrders) // 入库单列表
			inventory.POST("/receipt", handlers.CreateReceipt)   // 1. 收料
			inventory.POST("/iqc/:id", handlers.ConfirmIQC)     // 2. 检验
			inventory.POST("/inbound/:id", handlers.ConfirmInbound) // 3. 入库
			inventory.POST("/outbound", handlers.CreateOutboundOrder) // 出库
			inventory.GET("/outbound", handlers.GetOutboundOrders)    // 出库列表
		}

		// 核心管理员权限
		core := auth.Group("/core")
		core.Use(middleware.RoleMiddleware(models.RoleCoreAdmin))
		{
			core.GET("/dashboard", handlers.GetDashboardStats)
		}

		// 普通管理员权限
		admin := auth.Group("/admin")
		admin.Use(middleware.RoleMiddleware(models.RoleAdmin, models.RoleCoreAdmin))
		{
			admin.GET("/info", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "这是管理员信息页面"})
			})
		}
	}

	// 4. 启动服务
	port := fmt.Sprintf(":%d", config.AppConfig.Server.Port)
	fmt.Printf("Server is running on port %s\n", port)
	r.Run(port)
}
