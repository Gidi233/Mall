package service

import (
	"Mall/service/mall"
	"Mall/service/manage"
	"Mall/service/file"
)

type ServiceGroup struct {
	ExampleServiceGroup file.ServiceGroup
	ManageServiceGroup  manage.ManageServiceGroup
	MallServiceGroup    mall.MallServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
