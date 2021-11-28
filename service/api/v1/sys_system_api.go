package v1

import (
	"fast-admin-service/pkg/app"
	"fast-admin-service/pkg/enum"
	"fast-admin-service/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SystemApi struct {
}

type ISystemApi interface {
	GetAListOfSystemIcons(c *gin.Context)
}

var sysSystemService services.SysSystemService

// GetAListOfSystemIcons
// @Tags Auth
// @Summary 获取系统图标
// @Success 200 {string} string "{"success":true,"data":{},"msg":"ok"}"
// @Router /sys/icons [get]
func (systemApi *SystemApi) GetAListOfSystemIcons(c *gin.Context) {
	appRes := app.Gin{C: c}
	icons, err := sysSystemService.GetSystemIcons()
	if err != nil {
		return
	}
	appRes.Response(http.StatusOK, enum.SUCCESS, gin.H{
		"icons": icons,
	})
}
