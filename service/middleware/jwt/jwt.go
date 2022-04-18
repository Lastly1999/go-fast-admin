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

		code = enum.SUCCESS
		token := c.GetHeader("authorization")
		if token == "" {
			code = enum.INVALID_PARAMS
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = enum.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = enum.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}

		if code != enum.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  enum.GetMsg(code),
				"data": nil,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
