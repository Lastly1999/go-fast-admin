package v1

import (
	"fast-admin-service/model"
	"fast-admin-service/model/request"
	"fast-admin-service/pkg/app"
	"fast-admin-service/pkg/captcha"
	"fast-admin-service/pkg/enum"
	"fast-admin-service/pkg/utils"
	"fast-admin-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthApi struct {
}

type IAuthApi interface {
	LoginAction(c *gin.Context)
	GetUserInfoById(c *gin.Context)
	GetAuthCode(c *gin.Context)
	GetBaseMenus(c *gin.Context)
	GetBaseMenusIds(c *gin.Context)
	UpdateUserRole(c *gin.Context)
}

var authService services.AuthService

// LoginAction
// @Tags Auth
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /auth/login [post]
func (authApi *AuthApi) LoginAction(c *gin.Context) {
	appRes := app.Gin{C: c}
	loginParam := &request.Login{}
	err := c.ShouldBindJSON(loginParam)
	// 绑定json错误
	if err != nil {
		appRes.Response(http.StatusOK, enum.BIN_JSON_ERROR, nil)
		return
	}
	userReqBody := &model.SysUser{
		UserName: loginParam.UserName,
		PassWord: utils.EncryptSh256(loginParam.PassWord),
	}
	// 验证图形验证码
	res := captcha.Verify(loginParam.CodeAuth, loginParam.Code)
	if res {
		auth, err := authService.CheckAuth(userReqBody)
		if err != nil {
			appRes.Response(http.StatusOK, enum.AUTH_ERROR, nil)
			return
		}
		if auth.RoleId == 0 {
			appRes.Response(http.StatusOK, enum.AUTHORITY_ERROR, nil)
			return
		}
		// 成功 派发token 用户的权限 默认选择第一个作为默认角色
		token, err := utils.GenerateToken(auth.UserName, auth.PassWord, auth.RoleId, uint(int(auth.ID)))
		if err != nil {
			appRes.Response(http.StatusOK, enum.ERROR_AUTH, "token派发错误")
			return
		}
		// 返回接口
		appRes.Response(http.StatusOK, enum.SUCCESS, gin.H{
			"token": token,
		})
	} else {
		appRes.Response(http.StatusOK, enum.CODE_ERROR, nil)
	}
}

// GetUserInfoById
// @Tags Auth
// @Summary 获取系统用户信息
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /auth/user [get]
func (authApi *AuthApi) GetUserInfoById(c *gin.Context) {
	appRes := app.Gin{C: c}
	info, err := utils.ParseTokenRequest(c)
	if err != nil {
		appRes.Response(http.StatusOK, enum.INVALID_TOKEN_PARAMS_ERROR, nil)
		return
	}
	userInfo, err := authService.GetSystemUserInfoById(int(info.UserId))
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, gin.H{
		"userInfo": userInfo,
	})
}

// GetAuthCode
// @Tags Auth
// @Summary 获取图片验证码
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /auth/code [post]
func (authApi *AuthApi) GetAuthCode(c *gin.Context) {
	appRes := app.Gin{C: c}
	code, base, err := authService.GenerateVerificode()
	if err != nil {
		appRes.Response(http.StatusBadGateway, enum.ERROR, "Fail:获取参数失败")
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, gin.H{
		"code":     code,
		"codeBase": base,
	})
}

// GetBaseMenus
// @Tags Auth
// @Summary 获取基础权限菜单
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /auth/menu [get]
func (authApi *AuthApi) GetBaseMenus(c *gin.Context) {
	appRes := app.Gin{C: c}
	info, err := utils.ParseTokenRequest(c)
	if err != nil {
		appRes.Response(http.StatusOK, enum.INVALID_TOKEN_PARAMS_ERROR, nil)
		return
	}
	menus, err := authService.GetSystemPermissionsMenu(strconv.Itoa(int(info.RoleId)))
	if err != nil {
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, gin.H{
		"menus": menus,
	})
}

// GetBaseMenusIds
// @Tags Auth
// @Summary 获取用户权限菜单id组
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /auth/menuids/:id [get]
func (authApi *AuthApi) GetBaseMenusIds(c *gin.Context) {
	appRes := app.Gin{C: c}
	// 解析token内的用户参数
	id := c.Param("id")
	ids, err := authService.GetSystemPermissionsMenuIds(id)
	if err != nil {
		appRes.Response(http.StatusOK, enum.INVALID_TOKEN_PARAMS_ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, gin.H{
		"roleIds": ids,
	})
}

// UpdateUserRole
// @Tags Auth
// @Summary 更新用户默认角色
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /auth/role/:id [patch]
func (authApi *AuthApi) UpdateUserRole(c *gin.Context) {
	appRes := app.Gin{C: c}
	sysRoleDefaultParams := request.SysRoleDefaultParams{}
	err := c.ShouldBindJSON(&sysRoleDefaultParams)
	if err != nil {
		appRes.Response(http.StatusOK, enum.BIN_JSON_ERROR, nil)
		return
	}
	err = userService.UpdateUserRole(&sysRoleDefaultParams)
	if err != nil {
		appRes.Response(http.StatusOK, enum.ERROR, nil)
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, nil)
}
