package router

import (
	v1 "fast-admin-service/api/v1"
	"fast-admin-service/global"
	"fast-admin-service/middleware/enforce"
	"fast-admin-service/middleware/jwt"
	"fast-admin-service/pkg/adapter"
	"fast-admin-service/router/routes"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"time"
)

func initMiddleware(app *gin.Engine) {
	// 将gin的自带日志转为zap日志
	app.Use(ginzap.Ginzap(global.ZAP_LOG, time.RFC3339, true))
	// 堆栈错误log
	app.Use(ginzap.RecoveryWithZap(global.ZAP_LOG, true))
	//app.Use(req.RequestParams())
}

func InitRouter() (app *gin.Engine) {
	gin.ForceConsoleColor()
	app = gin.New()
	// 初始化中间件
	initMiddleware(app)
	api := app.Group("v1")
	authApi := v1.AuthApi{}
	// 授权登录
	api.POST("/auth/login", authApi.LoginAction)
	// 获取图形验证码
	api.GET("/auth/code", authApi.GetAuthCode)
	api.Use(jwt.JWT())                                     // jwt 中间件
	api.Use(enforce.ApiCheckRule(adapter.EnforcerTools())) // Casbin api鉴权
	{
		// 授权模块
		routes.InitAuthRouter(api)
		// 用户模块
		routes.InitUserRouter(api)
		// 角色模块
		routes.InitRoleRouter(api)
		// 系统模块
		routes.InitSystemRouter(api)
		// 系统菜单模块
		routes.InitBaseMenuRouter(api)
		// Casbin policys模块
		routes.InitCasbinRoutes(api)
	}
	return app
}
