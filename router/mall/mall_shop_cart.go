package mall

import (
	"Mall/api"
	"Mall/middleware"

	"github.com/gin-gonic/gin"
)

type MallShopCartRouter struct {//换
}

func (m *MallUserRouter) InitMallShopCartRouter(Router *gin.RouterGroup) {
	mallShopCartRouter := Router.Use(middleware.UserJWTAuth())
	var mallShopCartApi = api.ApiGroupApp.MallApiGroup.MallShopCartApi
	{
		mallShopCartRouter.GET("/shop-cart", mallShopCartApi.CartItemList)                                             //购物车列表(网页移动端不分页)
		mallShopCartRouter.POST("/shop-cart", mallShopCartApi.SaveMallShoppingCartItem)                                //添加购物车
		mallShopCartRouter.PUT("/shop-cart", mallShopCartApi.UpdateMallShoppingCartItem)                               //修改购物车
		mallShopCartRouter.DELETE("/shop-cart/:newBeeMallShoppingCartItemId", mallShopCartApi.DelMallShoppingCartItem) //删除购物车
		mallShopCartRouter.GET("/shop-cart/settle", mallShopCartApi.ToSettle)                                          //根据购物项id数组查询购物项明细

	}
}
