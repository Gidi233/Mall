package manage

import (
	"Mall/global"
	"Mall/model/common/request"
	"Mall/model/common/response"
	"Mall/model/manage"
	manageReq "Mall/model/manage/request"
	"Mall/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ManageAdminUserApi struct {
}

// 创建AdminUser
func (m *ManageAdminUserApi) CreateAdminUser(c *gin.Context) {
	var params manageReq.MallAdminParam
	_ = c.ShouldBindJSON(&params)
	if err := utils.Verify(params, utils.AdminUserRegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	mallAdminUser := manage.MallAdminUser{
		LoginUserName: params.LoginUserName,
		NickName:      params.NickName,
		LoginPassword: utils.MD5V([]byte(params.LoginPassword)),
	}
	if err := mallAdminUserService.CreateMallAdminUser(mallAdminUser); err != nil {
		global.GVA_LOG.Error("创建失败:", zap.Error(err))
		response.FailWithMessage("创建失败"+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// 修改密码
func (m *ManageAdminUserApi) UpdateAdminUserPassword(c *gin.Context) {
	var req manageReq.MallUpdatePasswordParam
	_ = c.ShouldBindJSON(&req)
	//mallAdminUserName := manage.MallAdminUser{
	//	LoginPassword: utils.MD5V([]byte(req.LoginPassword)),
	//}
	userToken := c.GetHeader("token")
	if err := mallAdminUserService.UpdateMallAdminPassWord(userToken, req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}

}

// 更新用户名
func (m *ManageAdminUserApi) UpdateAdminUserName(c *gin.Context) {
	var req manageReq.MallUpdateNameParam
	_ = c.ShouldBindJSON(&req)
	userToken := c.GetHeader("token")
	if err := mallAdminUserService.UpdateMallAdminName(userToken, req); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// AdminUserProfile 用id查询AdminUser
func (m *ManageAdminUserApi) AdminUserProfile(c *gin.Context) {
	adminToken := c.GetHeader("token")
	if err, mallAdminUser := mallAdminUserService.GetMallAdminUser(adminToken); err != nil {
		global.GVA_LOG.Error("未查询到记录", zap.Error(err))
		response.FailWithMessage("未查询到记录", c)
	} else {
		mallAdminUser.LoginPassword = "******"
		response.OkWithData(mallAdminUser, c)
	}
}

// AdminLogin 管理员登陆
func (m *ManageAdminUserApi) AdminLogin(c *gin.Context) {
	var adminLoginParams manageReq.MallAdminLoginParam
	_ = c.ShouldBindJSON(&adminLoginParams)
	if err, _, adminToken := mallAdminUserService.AdminLogin(adminLoginParams); err != nil {
		response.FailWithMessage("登陆失败", c)
	} else {
		response.OkWithData(adminToken.Token, c)
	}
}

// AdminLogout 登出
func (m *ManageAdminUserApi) AdminLogout(c *gin.Context) {
	token := c.GetHeader("token")
	if err := mallAdminUserTokenService.DeleteMallAdminUserToken(token); err != nil {
		response.FailWithMessage("登出失败", c)
	} else {
		response.OkWithMessage("登出成功", c)
	}

}

// UserList 商城注册用户列表
func (m *ManageAdminUserApi) UserList(c *gin.Context) {
	var pageInfo manageReq.MallUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := mallUserService.GetMallUserInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:       list,
			TotalCount: total,
			CurrPage:   pageInfo.PageNumber,
			PageSize:   pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// LockUser 用户禁用与解除禁用(0-未锁定 1-已锁定)
func (m *ManageAdminUserApi) LockUser(c *gin.Context) {
	lockStatus, _ := strconv.Atoi(c.Param("lockStatus"))
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := mallUserService.LockUser(IDS, lockStatus); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// TODO:UploadFile 上传单图
