package jwt

import (
	"fast-admin-service/pkg/enum"
	"fast-admin-service/pkg/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
// 中间件token验证
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = enum.SUCCESS
		//token := c.Query("token")
		token := c.GetHeader("authorization")
		if token == "" {
			code = enum.INVALID_PARAMS
			data = "jwt Token validation failed, null，令牌不能为空"
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = enum.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
					data = "Token verification failed timeout，请重新登录"
				default:
					code = enum.ERROR_AUTH_CHECK_TOKEN_FAIL
					data = "jwt Token validation failed，请检查是否令牌是否有误"
				}
			}
		}

		if code != enum.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  enum.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
