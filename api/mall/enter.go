package mall

import "Mall/service"

type MallGroup struct {
	MallUserApi
}

var mallUserTokenService = service.ServiceGroupApp.MallServiceGroup.MallUserTokenService
var mallUserService = service.ServiceGroupApp.MallServiceGroup.MallUserService
