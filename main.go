package main

import (
	"goblog/model"
	"goblog/routers"
)

// 首先构造数据库，围绕数据库写后端
func main() {
	// 引用数据库
	model.InitDb()

	routers.InitRouter()
}
