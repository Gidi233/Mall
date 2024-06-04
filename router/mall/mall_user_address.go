package mall

import (
	 "Mall/api"
	"Mall/middleware"

	"github.com/gin-gonic/gin"
)

type MallUserAddressRouter struct {
}

func (m *MallUserRouter) InitMallUserAddressRouter(Router *gin.RouterGroup) {
	mallUserAddressRouter := Router.Use(middleware.UserJWTAuth())
	var mallUserAddressApi = api.ApiGroupApp.MallApiGroup.MallUserAddressApi
	{
		mallUserAddressRouter.GET("/address", mallUserAddressApi.AddressList)                       //用户地址
		mallUserAddressRouter.POST("/address", mallUserAddressApi.SaveUserAddress)                  //添加地址
		mallUserAddressRouter.PUT("/address", mallUserAddressApi.UpdateMallUserAddress)             //修改用户地址
		mallUserAddressRouter.GET("/address/:addressId", mallUserAddressApi.GetMallUserAddress)     //获取地址详情
		mallUserAddressRouter.GET("/address/default", mallUserAddressApi.GetMallUserDefaultAddress) //获取默认地址
		mallUserAddressRouter.DELETE("/address/:addressId", mallUserAddressApi.DeleteUserAddress)   //删除地址
	}

}
