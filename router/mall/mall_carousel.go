package mall

import (
	"Mall/api"

	"github.com/gin-gonic/gin"
)

type MallCarouselIndexRouter struct {
}

func (m *MallCarouselIndexRouter) InitMallCarouselIndexRouter(Router *gin.RouterGroup) {
	mallCarouselRouter := Router
	var mallCarouselApi = api.ApiGroupApp.MallApiGroup.MallIndexApi
	{
		mallCarouselRouter.GET("index-infos", mallCarouselApi.MallIndexInfo) // 获取首页数据
	}
}
