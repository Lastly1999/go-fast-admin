package routes

import (
	v1 "fast-admin-service/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRoleRouter(r *gin.RouterGroup) {
	roleRouter := r.Group("role")
	roleApi := v1.RoleApi{}
	{
		// 获取角色列表
		roleRouter.GET("/role", roleApi.GetRoles)
		// 添加角色
		roleRouter.PUT("/role", roleApi.PutRole)
		// 修改角色信息
		roleRouter.PATCH("/role", roleApi.UpdateRoleById)
		// 删除角色信息
		roleRouter.DELETE("/role/:id", roleApi.DeleteRoleById)
		// 修改角色菜单
		roleRouter.PATCH("/menu", roleApi.UpdateRoleBaseMenu)
	}
}
