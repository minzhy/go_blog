package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// 这个引用了gorm提供的结构体，里面包括了ID，CreatedAt……。一共四个参数。
	Username string `gorm:"type: varchar(20);not null" json:"username"`
	Password string `gorm:"type: varchar(20);not null" json:"password"`
	Role     int    `gorm:"type: int" json:"role"`
	// avatar 头像
}
