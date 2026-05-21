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
		// 核心管理员权限
		core := auth.Group("/core")
		core.Use(middleware.RoleMiddleware(models.RoleCoreAdmin))
		{
			core.GET("/dashboard", func(c *gin.Context) {
				c.JSON(200, gin.H{"message": "欢迎进入核心管理员面板"})
			})
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
