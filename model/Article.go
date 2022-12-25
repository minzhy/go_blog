package model

import "gorm.io/gorm"

type Article struct {
	// 这里一定要加上外健的注释，否则这里会编译错误
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	Cid   int    `gorm:"type:int;not null" json:"cid"`
	// Cid 就是Category的id
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}
