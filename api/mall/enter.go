package mall

import "Mall/service"

type MallGroup struct {
	MallIndexApi
	MallGoodsInfoApi
	MallGoodsCategoryApi
	MallUserApi
	MallUserAddressApi
	MallShopCartApi
	MallOrderApi
}

var mallCarouselService = service.ServiceGroupApp.MallServiceGroup.MallCarouselService
var mallGoodsInfoService = service.ServiceGroupApp.MallServiceGroup.MallGoodsInfoService
var mallIndexConfigService = service.ServiceGroupApp.MallServiceGroup.MallIndexInfoService
var mallGoodsCategoryService = service.ServiceGroupApp.MallServiceGroup.MallGoodsCategoryService
var mallUserTokenService = service.ServiceGroupApp.MallServiceGroup.MallUserTokenService
var mallUserService = service.ServiceGroupApp.MallServiceGroup.MallUserService
var mallUserAddressService = service.ServiceGroupApp.MallServiceGroup.MallUserAddressService
var mallShopCartService = service.ServiceGroupApp.MallServiceGroup.MallShopCartService
var mallOrderService = service.ServiceGroupApp.MallServiceGroup.MallOrderService
