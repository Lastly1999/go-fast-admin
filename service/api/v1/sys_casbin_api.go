package v1

import (
	"fast-admin-service/pkg/app"
	"fast-admin-service/pkg/enum"
	"fast-admin-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SysCasbinApi struct {
}

type ISysCasbinApi interface {
	GetPolicys(c *gin.Context)
}

var casbinService services.SysCasbinRuleService

func (sysCasbinApi *SysCasbinApi) GetPolicys(c *gin.Context) {
	appRes := app.Gin{C: c}
	policys := casbinService.GetPolicys()
	appRes.Response(http.StatusOK, enum.SUCCESS, policys)
}
