package mall

import "Mall/service"

type MallGroup struct {
	MallIndexApi
	MallGoodsInfoApi
	MallUserApi
}

var mallCarouselService = service.ServiceGroupApp.MallServiceGroup.MallCarouselService
var mallGoodsInfoService = service.ServiceGroupApp.MallServiceGroup.MallGoodsInfoService
var mallIndexConfigService = service.ServiceGroupApp.MallServiceGroup.MallIndexInfoService
var mallUserTokenService = service.ServiceGroupApp.MallServiceGroup.MallUserTokenService
var mallUserService = service.ServiceGroupApp.MallServiceGroup.MallUserService
