package routes

import (
	v1 "fast-admin-service/api/v1"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) {
	userRouter := r.Group("user")
	userApi := v1.UserApi{}
	{
		// 获取系统用户
		userRouter.POST("user", userApi.GetSystemUsers)
		// 创建系统用户
		userRouter.PUT("user", userApi.CreateSystemUser)
		// 更新系统用户
		userRouter.PATCH("user", userApi.UpdateSystemUserById)
		// 删除系统用户
		userRouter.DELETE("user/:id", userApi.DeleteSystemUserById)
		// 获取系统用户角色列表
		userRouter.GET("role/:id", userApi.GetAListOfUserRoles)
		// 新增用户关联角色
		userRouter.PATCH("role", userApi.NewUserAssociationRole)
	}
}
