package router

import (
	"Mall/router/mall"
	"Mall/router/manage"
)

type RouterGroup struct {
	Manage manage.ManageRouterGroup
	Mall   mall.MallRouterGroup
}

var RouterGroupApp = new(RouterGroup)
