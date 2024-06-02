package api

import (
	"Mall/api/mall"
	"Mall/api/manage"
)

type ApiGroup struct {
	ManageApiGroup manage.ManageGroup
	MallApiGroup   mall.MallGroup
}

var ApiGroupApp = new(ApiGroup)
