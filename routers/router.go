package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles"
	//_ "github.com/swaggo/swag/example/basic/docs"
	_ "gin-blog/docs"

	"gin-blog/middleware/jwt"
	"gin-blog/pkg/setting"
	"gin-blog/routers/api"
	"gin-blog/routers/api/v1"
)

func InitRouter() *gin.Engine {

	r := gin.New()

	// 使用 Logger 中间件
	// 将日志写到默认的写入设备
	// 每一次的请求信息都会被输出到终端上，默认输出是 os.Stdout
	r.Use(gin.Logger())
	// 使用 Recovery 中间件
	r.Use(gin.Recovery())
	// 设置运行模式
	gin.SetMode(setting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/auth", api.GetAuth)

	// 注册路由
	// gin.Context（核心） 是 gin 中的上下文，
	apiv1 := r.Group("/api/v1")
	// 使用 jwt 中间件
	apiv1.Use(jwt.JWT())
	{
		// 获取标签列表页
		apiv1.GET("/tags", v1.GetTags)
		// 新建标签
		apiv1.POST("/tags", v1.AddTag)
		// 更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 获取一篇文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		// 新增文章
		apiv1.POST("/articles", v1.AddArticle)
		// 修改文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		// 删除文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}