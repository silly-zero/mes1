package models

import (
	"database/sql"
	"fmt"
	"log"
	"mes-backend/config"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dbConfig := config.AppConfig.Database

	// 1. 尝试连接 MySQL 服务器（不指定数据库）来创建数据库
	dsnRoot := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port)
	createDbSql, err := sql.Open("mysql", dsnRoot)
	if err != nil {
		log.Fatalf("Failed to connect to MySQL server: %v", err)
	}
	defer createDbSql.Close()

	_, err = createDbSql.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", dbConfig.DBName))
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	// 2. 连接具体数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 自动迁移
	err = db.AutoMigrate(&User{}, &Material{}, &InboundOrder{}, &OutboundOrder{}, &Stock{}, &StockLog{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	DB = db
	fmt.Println("Database connected and migrated successfully")

	// 3. 初始化默认账号
	SeedAdmin()
}

// SeedAdmin 初始化默认核心管理员
func SeedAdmin() {
	var count int64
	DB.Model(&User{}).Where("username = ?", "user").Count(&count)
	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
		user := User{
			Username: "user",
			Password: string(hashedPassword),
			Nickname: "系统管理员",
			Role:     RoleCoreAdmin,
		}
		if err := DB.Create(&user).Error; err != nil {
			fmt.Printf("Failed to seed admin user: %v\n", err)
		} else {
			fmt.Println("Default admin user (user/123456) created successfully")
		}
	}
}
