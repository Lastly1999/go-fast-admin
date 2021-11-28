package req

import (
	"github.com/gin-gonic/gin"
)

func RequestParams() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.Method == "POST" {
			//params := make(map[string]interface{})
			//data, _ := context.GetRawData()

			//if err != nil {
			//	return
			//}
			//marshal, err := json.Marshal(params)
			//global.ZAP_LOG.Info(string(data))
			return
		}
		context.Next()
	}
}
