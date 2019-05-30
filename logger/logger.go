package logger

var isInitialized = false

// logDate set by function chechDate
var logDate struct {
	year  int
	month int
	day   int
}

type logType int

const (
	// DEBUG output everything
	DEBUG logType = 0
	// INFO output user needed
	INFO logType = 1
	// WARN output selfFixed error
	WARN logType = 2
	// ERROR output attention error
	ERROR logType = 3
	// FATAL output error must stop
	FATAL logType = 4
)

func (f logType) String() string {
	switch f {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// LogSetting control the
// destination of log infomation
var logSetting struct {
	logToFile bool
	logToSys  bool
	// isDebug value: 0/1
	// 0: output Info leval
	// 1: output Debug and Info leval
	isDebug bool
}

var logFileName string
