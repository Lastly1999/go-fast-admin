package enforce

import (
	"fast-admin-service/global"
	"fast-admin-service/pkg/utils"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// ApiCheckRule Casbin的api鉴权认证中间件
func ApiCheckRule(e *casbin.Enforcer) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取api路径
		obj := context.Request.URL.RequestURI()
		// 获取api请求方法
		act := context.Request.Method
		token := context.GetHeader("Authorization")
		parseToken, err := utils.ParseToken(token)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "casbin 鉴权token解析错误",
				"data": nil,
			})
			return
		}
		// 获取角色id 为了跟数据库的policy规则对比 sub
		sub := strconv.Itoa(int(parseToken.RoleId))
		// 判断策略中是否存在放行
		enforce, err := e.Enforce(sub, obj, act)
		if err != nil {
			global.ZAP_LOG.Info("权限认证失败，您没有此api权限")
			errMsg := fmt.Sprintf("RequestURL：%s errorMsg：api鉴权错误，您没有此api权限", obj)
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  errMsg,
				"data": nil,
			})
			context.Abort()
		}
		if enforce {
			successMsg := fmt.Sprintf("RequestURL：%s msg：api鉴权通过，身份认证成功", obj)
			global.ZAP_LOG.Info(successMsg)
			context.Next()
		} else {
			global.ZAP_LOG.Info("权限认证失败，您没有此api权限")
			errMsg := fmt.Sprintf("RequestURL：%s errorMsg：api鉴权错误，您没有此api权限", obj)
			context.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  errMsg,
				"data": nil,
			})
			context.Abort()
		}
	}
}
