package cors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//	CorsMiddleWare 配置跨域中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//	获取请求方式
		method := c.Request.Method
		//	放行请求头规则
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Max-Age", "20000")
		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
