package request

import (
	"Mall/model/common/request"
	"Mall/model/manage"
)

type MallUserSearch struct {
	manage.MallUser
	request.PageInfo
}
