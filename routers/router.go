package routers

import (
	"goblog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	// 也可以用gin.New，但是，New的话，相对于Default少了几个中间件

	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}
	// 访问：http://localhost:3000/api/v1/hello，Group的意义在于把访问的前缀固定了

	r.Run(utils.HttpPort)
	// 或者返回r.Engine
}
