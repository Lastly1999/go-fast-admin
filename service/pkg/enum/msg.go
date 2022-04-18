package enum

var MsgFlags = map[int]string{
	SUCCESS:                        "ok",
	ERROR:                          "fail",
	INVALID_PARAMS:                 "JWT authentication error, please login again",
	AUTH_ERROR:                     "登录失败，请检查参数是否正确",
	CODE_ERROR:                     "图形验证码错误，请重新登录",
	BIN_JSON_ERROR:                 "JSON解析异常,请检查数据格式",
	PARAMS_ERROR:                   "url参数解析异常，请检查数据格式",
	AUTH_FAIL:                      "Casbin Authentication failed 权限认证失败，无api权限",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token verification failed timeout，请重新登录",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "jwt Token validation failed，请检查是否令牌是否有误",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
