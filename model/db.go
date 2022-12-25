package model

import (
	"fmt"
	"goblog/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 配置我们db配置的参数，相当于一个汇总
// 利用gorm库写

var db *gorm.DB
var err error

func InitDb() {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数:", err)
	}

	/*
		GORM 的AutoMigrate() 方法用于自动迁移ORM 的Schemas。
		所谓“迁移” 就是刷新数据库中的表格定义，使其保持最新（只增不减）。
		AutoMigrate 会创建（新的）表、缺少的外键、约束、列和索引，并且会更改现有列的类型（如果其大小、精度、是否为空可更改的话）。
	*/
	db.AutoMigrate(&User{}, &Article{}, &Category{})
	// 相当于自动创建了这几个表

	// 本质上是用一个db的连接池连接的
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("数据库配置错误")
	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	// 这里最好不要大于routers的连接时间
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	// sqlDB.Close()
}
