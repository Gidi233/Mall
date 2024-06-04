package manage

import "Mall/service"

type ManageGroup struct {
	ManageAdminUserApi
}

var mallAdminUserService = service.ServiceGroupApp.ManageServiceGroup.ManageAdminUserService
var mallAdminUserTokenService = service.ServiceGroupApp.ManageServiceGroup.ManageAdminUserTokenService
var mallUserService = service.ServiceGroupApp.ManageServiceGroup.ManageUserService
var fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
