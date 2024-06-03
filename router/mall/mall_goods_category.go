package mall

import (
	"Mall/api"

	"github.com/gin-gonic/gin"
)

type MallGoodsCategoryIndexRouter struct {
}

func (m *MallGoodsInfoIndexRouter) InitMallGoodsCategoryIndexRouter(Router *gin.RouterGroup) {
	mallGoodsRouter := Router
	var mallGoodsCategoryApi = api.ApiGroupApp.MallApiGroup.MallGoodsCategoryApi
	{
		mallGoodsRouter.GET("categories", mallGoodsCategoryApi.GetGoodsCategory) // 获取分类数据
	}
}
