package middleware

import (
	"Mall/model/common/response"
	"Mall/service"
	"time"

	"github.com/gin-gonic/gin"
)

var manageAdminUserTokenService = service.ServiceGroupApp.ManageServiceGroup.ManageAdminUserTokenService
var mallUserTokenService = service.ServiceGroupApp.MallServiceGroup.MallUserTokenService

// TODO:换成真的JWT
func AdminJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithDetailed(nil, "未登录或非法访问", c)
			c.Abort()
			return
		}
		err, mallAdminUserToken := manageAdminUserTokenService.ExistAdminToken(token)
		if err != nil {
			response.FailWithDetailed(nil, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if time.Now().After(mallAdminUserToken.ExpireTime) {
			response.FailWithDetailed(nil, "授权已过期", c)
			err = manageAdminUserTokenService.DeleteMallAdminUserToken(token)
			if err != nil {
				return
			}
			c.Abort()
			return
		}
		c.Next()
	}

}

func UserJWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.UnLogin(nil, c)
			c.Abort()
			return
		}
		err, mallUserToken := mallUserTokenService.ExistUserToken(token)
		if err != nil {
			response.UnLogin(nil, c)
			c.Abort()
			return
		}
		if time.Now().After(mallUserToken.ExpireTime) {
			response.FailWithDetailed(nil, "授权已过期", c)
			err = mallUserTokenService.DeleteMallUserToken(token)
			c.Abort()
			if err != nil {
				return
			}
			return
		}
		c.Next()
	}

}
