package manage

import (
	"Mall/api"
	"Mall/middleware"

	"github.com/gin-gonic/gin"
)

type ManageAdminUserRouter struct {
}

func (r *ManageAdminUserRouter) InitManageAdminUserRouter(Router *gin.RouterGroup) {
	mallAdminUserRouter := Router.Use(middleware.AdminJWTAuth())
	mallAdminUserWithoutRouter := Router
	var mallAdminUserApi =api.ApiGroupApp.ManageApiGroup.ManageAdminUserApi
	{
		mallAdminUserRouter.POST("createMallAdminUser", mallAdminUserApi.CreateAdminUser) // 新建MallAdminUser
		mallAdminUserRouter.PUT("adminUser/name", mallAdminUserApi.UpdateAdminUserName)   // 更新MallAdminUser
		mallAdminUserRouter.PUT("adminUser/password", mallAdminUserApi.UpdateAdminUserPassword)
		mallAdminUserRouter.GET("users", mallAdminUserApi.UserList)
		mallAdminUserRouter.PUT("users/:lockStatus", mallAdminUserApi.LockUser)
		mallAdminUserRouter.GET("adminUser/profile", mallAdminUserApi.AdminUserProfile) // 根据ID获取 admin详情
		mallAdminUserRouter.DELETE("logout", mallAdminUserApi.AdminLogout)
		// mallAdminUserRouter.POST("upload/file", mallAdminUserApi.UploadFile) //上传图片

	}
	{
		mallAdminUserWithoutRouter.POST("adminUser/login", mallAdminUserApi.AdminLogin) //管理员登陆
	}
}
