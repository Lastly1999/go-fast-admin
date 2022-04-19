package v1

import (
	"fast-admin-service/global"
	"fast-admin-service/model"
	"fast-admin-service/model/request"
	"fast-admin-service/pkg/app"
	"fast-admin-service/pkg/enum"
	"fast-admin-service/pkg/utils"
	"fast-admin-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserApi struct {
}

type IUserApi interface {
	GetSystemUsers(c *gin.Context)
	CreateSystemUser(c *gin.Context)
	UpdateSystemUserById(c *gin.Context)
	DeleteSystemUserById(c *gin.Context)
	GetAListOfUserRoles(c *gin.Context)
	NewUserAssociationRole(c *gin.Context)
}

var userService services.UserService

// GetSystemUsers
// @Tags User
// @Summary 获取系统用户
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /user/user [post]
func (userApi *UserApi) GetSystemUsers(c *gin.Context) {
	appRes := app.Gin{C: c}
	//infoParams := request.PageInfo{}
	infoParams := request.SystemUserListParams{}
	err := c.ShouldBindJSON(&infoParams)
	if err != nil {
		global.ZAP_LOG.Error(err.Error())
		appRes.Response(http.StatusOK, enum.BIN_JSON_ERROR, nil)
		return
	}
	users, total, err := userService.GetUsers(&infoParams)
	if err != nil {
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, gin.H{
		"users": users,
		"total": total,
	})
}

// CreateSystemUser
// @Tags User
// @Summary 新增系统用户
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /user/user [put]
func (userApi *UserApi) CreateSystemUser(c *gin.Context) {
	appRes := app.Gin{C: c}
	sysUserParams := request.SystemUserParams{}
	err := c.ShouldBindJSON(&sysUserParams)
	if err != nil {
		appRes.Response(http.StatusOK, enum.BIN_JSON_ERROR, nil)
		return
	}
	user := &model.SysUser{
		UserName: sysUserParams.UserName,
		PassWord: utils.EncryptSh256(sysUserParams.PassWord),
		NikeName: sysUserParams.NikeName,
		RoleId:   sysUserParams.RoleId,
	}
	for _, v := range sysUserParams.RoleIds {
		user.Role = append(user.Role, model.SysRole{
			RoleId: v,
		})
	}
	err = userService.CreateUser(user)
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, nil)
}

// UpdateSystemUserById
// @Tags User
// @Summary 修改系统用户信息
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /user/user [patch]
func (userApi *UserApi) UpdateSystemUserById(c *gin.Context) {
	appRes := app.Gin{C: c}
	sysUserParams := request.SystemUserParams{}
	err := c.ShouldBindJSON(&sysUserParams)
	if err != nil {
		appRes.Response(http.StatusOK, enum.BIN_JSON_ERROR, nil)
		return
	}
	user := &model.SysUser{
		Model: global.Model{
			ID: sysUserParams.Id,
		},
		UserName: sysUserParams.UserName,
		PassWord: sysUserParams.PassWord,
		NikeName: sysUserParams.NikeName,
		RoleId:   sysUserParams.RoleId,
	}
	for _, v := range sysUserParams.RoleIds {
		user.Role = append(user.Role, model.SysRole{
			RoleId: v,
		})
	}
	err = userService.UpdateUserById(user)
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, nil)
}

// DeleteSystemUserById
// @Tags User
// @Summary 删除系统用户
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /user/user [DELETE]
func (userApi *UserApi) DeleteSystemUserById(c *gin.Context) {
	appRes := app.Gin{C: c}
	param := c.Param("id")
	// 参数转换
	id, err := strconv.Atoi(param)
	if err != nil {
		appRes.Response(http.StatusOK, enum.PARAMS_ERROR, nil)
		return
	}
	err = userService.DeleteUserById(id)
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, nil)
}

// GetAListOfUserRoles
// @Tags User
// @Summary 获取系统用户角色列表
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /user/role [get]
func (userApi *UserApi) GetAListOfUserRoles(c *gin.Context) {
	appRes := app.Gin{C: c}
	param := c.Param("id")
	userId, err := strconv.Atoi(param)
	if err != nil {
		appRes.Response(http.StatusOK, enum.PARAMS_ERROR, nil)
		return
	}
	user := &model.SysUser{
		Model: global.Model{
			ID: uint(userId),
		},
	}
	roles, err := userService.GetSystemUserRoles(user)
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, gin.H{
		"roles": roles,
	})
}

// NewUserAssociationRole
// @Tags User
// @Summary 新增用户关联角色
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /user/role [patch]
func (userApi *UserApi) NewUserAssociationRole(c *gin.Context) {
	appRes := app.Gin{C: c}
	userRoleParams := request.SystemUserRoleParams{}
	err := c.ShouldBindJSON(&userRoleParams)
	if err != nil {
		appRes.Response(http.StatusOK, enum.BIN_JSON_ERROR, nil)
		return
	}
	err = userService.NewUserAssociationRole(&userRoleParams)
	if err != nil {
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, nil)
}
