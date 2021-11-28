package routes

import (
	v1 "fast-admin-service/api/v1"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(r *gin.RouterGroup) {
	systemRouter := r.Group("sys")
	systemApi := v1.SystemApi{}
	{
		// 获取系统图标列表
		systemRouter.GET("/icons", systemApi.GetAListOfSystemIcons)
	}
}
