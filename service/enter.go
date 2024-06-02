package service

import (
	"Mall/service/mall"
	"Mall/service/manage"
)

type ServiceGroup struct {
	ManageServiceGroup  manage.ManageServiceGroup
	MallServiceGroup    mall.MallServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
