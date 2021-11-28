package routes

import (
	v1 "fast-admin-service/api/v1"
	"github.com/gin-gonic/gin"
)

func InitCasbinRoutes(r *gin.RouterGroup) {
	casbinRouter := r.Group("policy")
	casbinApi := v1.SysCasbinApi{}
	{
		casbinRouter.GET("policy", casbinApi.GetPolicys)
	}
}
