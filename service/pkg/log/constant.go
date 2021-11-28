package log

const (
	DebugLevel = "debug"
	// LOGFORMAT_JSON logFormat
	LOGFORMAT_JSON    = "json"
	LOGFORMAT_CONSOLE = "console"
	// TIME_KEY EncoderConfig
	TIME_KEY       = "time"
	LEVLE_KEY      = "level"
	NAME_KEY       = "logger"
	CALLER_KEY     = "caller"
	MESSAGE_KEY    = "msg"
	STACKTRACE_KEY = "stacktrace"
	// MAX_SIZE 日志归档配置项
	// 每个日志文件保存的最大尺寸 单位：M
	MAX_SIZE = 1
	// MAX_BACKUPS 文件最多保存多少天
	MAX_BACKUPS = 5
	// MAX_AGE 日志文件最多保存多少个备份
	MAX_AGE = 7
)
