package user

import (
	"api-foodmarket/services/role"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id        int    `gorm:"not null;uniqueIndex;primary_key"`
	Name      string `gorm:"size:255;not null"`
	Email     string `gorm:"size:255;not null"`
	Password  string `gorm:"size:255;not null"`
	Phone     string `gorm:"size:255;not null"`
	Address   string `gorm:"size:255;not null"`
	City      string `gorm:"size:255;not null"`
	Photo     string `gorm:"size:255;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Roles     []UserRole
}

type UserRole struct {
	Id     int `gorm:"not null;uniqueIndex;primary_key"`
	UserId int
	RoleId int
	Role   role.Role
}
