package initialize

import (
	"Mall/global"
	"Mall/middleware"
	"Mall/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// https
	// Router.Use(middleware.LoadTls())
	// global.GVA_LOG.Info("use middleware secure")
	
	// 跨域、设置http协议
	Router.Use(middleware.Cors())
	global.GVA_LOG.Info("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	//商城后管路由
	manageRouter := router.RouterGroupApp.Manage
	ManageGroup := Router.Group("manage-api")
	manageRouter.InitManageAdminUserRouter(ManageGroup)


	


	// 健康监测
	PublicGroup := Router.Group("")
	PublicGroup.GET("/health", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	//商城前端路由
	mallRouter := router.RouterGroupApp.Mall
	MallGroup := Router.Group("api")
	

	global.GVA_LOG.Info("router register success")
	return Router
}
