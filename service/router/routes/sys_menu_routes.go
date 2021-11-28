package routes

import (
	v1 "fast-admin-service/api/v1"
	"github.com/gin-gonic/gin"
)

func InitBaseMenuRouter(r *gin.RouterGroup) {
	baseMenuRouter := r.Group("menu")
	baseMenuApi := v1.BaseMenuApi{}
	{
		// 获取系统菜单
		baseMenuRouter.GET("/menu", baseMenuApi.GetBaseMenu)
		// 新增系统菜单
		baseMenuRouter.PUT("/menu", baseMenuApi.PutBaseMenu)
		// 删除系统菜单
		baseMenuRouter.DELETE("/menu/:id", baseMenuApi.DeleteBaseMenu)
		// 编辑系统菜单
		baseMenuRouter.PATCH("/menu", baseMenuApi.UpdateBaseMenu)
		// 查询系统菜单详情
		baseMenuRouter.GET("/menu/:id", baseMenuApi.GetBaseMenuInfo)
	}
}
