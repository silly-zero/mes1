package models

import (
	"gorm.io/gorm"
)

type Role string

const (
	RoleCoreAdmin Role = "core_admin" // 核心管理员
	RoleAdmin     Role = "admin"      // 管理员
)

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;not null;size:50" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Nickname string `gorm:"size:50" json:"nickname"`
	Role     Role   `gorm:"type:varchar(20);default:'admin'" json:"role"`
}
