package routers

import (
	"gin-blog/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine  {
	r := gin.New()

	// 使用 Logger 中间件
	// 将日志写到默认的写入设备
	// 每一次的请求信息都会被输出到终端上，默认输出是 os.Stdout
	r.Use(gin.Logger())
	// 使用 Recovery 中间件
	r.Use(gin.Recovery())
	// 设置运行模式
	gin.SetMode(setting.RunMode)

	// 书写路由
	// gin.Context（核心） 是 gin 中的上下文，
	r.GET("/test", func(c *gin.Context) {
		// gin.H 实质是一个 map[string]interface{}
		c.JSON(200, gin.H{
			"message": "test message!",
		})
	})

	return r
}
