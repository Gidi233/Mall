package mall

import (
	"Mall/api"

	"github.com/gin-gonic/gin"
)

type MallGoodsInfoIndexRouter struct {
}

func (m *MallGoodsInfoIndexRouter) InitMallGoodsInfoIndexRouter(Router *gin.RouterGroup) {
	mallGoodsRouter := Router
	var mallGoodsInfoApi = api.ApiGroupApp.MallApiGroup.MallGoodsInfoApi
	{
		mallGoodsRouter.GET("/search", mallGoodsInfoApi.GoodsSearch)           // 商品搜索
		mallGoodsRouter.GET("/goods/detail/:id", mallGoodsInfoApi.GoodsDetail) //商品详情
	}
}
