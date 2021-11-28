package core

import (
	"fast-admin-service/router"
	"github.com/gin-gonic/gin"
)

func BootStarp() *gin.Engine {
	return router.InitRouter()
}
