package main

import "goblog/routers"

// 首先构造数据库，围绕数据库写后端
func main() {
	routers.InitRouter()
}
