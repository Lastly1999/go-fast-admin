package routes

import (
	v1 "fast-admin-service/api/v1"
	"github.com/gin-gonic/gin"
)

func InitAuthRouter(r *gin.RouterGroup) {
	authRouter := r.Group("auth")
	authApi := v1.AuthApi{}
	{
		// 获取用户基础权限菜单
		authRouter.GET("menu", authApi.GetBaseMenus)
		// 获取用户权限菜单id组
		authRouter.GET("menuids/:id", authApi.GetBaseMenusIds)
		// 获取系统用户信息
		authRouter.GET("user", authApi.GetUserInfoById)
		// 更新用户默认角色
		authRouter.PATCH("user", authApi.UpdateUserRole)
	}
}
