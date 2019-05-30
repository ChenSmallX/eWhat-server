package logger

import (
	"log"
)

// Init function: must run once and only once
func Init(toSys, toFile, isDebug bool) {
	if isInitialized == false {
		isInitialized = true
	} else {
		log.Fatalln("Fatal Error: logger initialized more than once...")
	}

	// set the logging mode
	logSetting.logToSys = toSys
	logSetting.logToFile = toFile
	logSetting.isDebug = isDebug

	// set the log file
	if logSetting.logToFile {
		getLogFile(DEBUG).Close()
	}

	logFileName = ""
}

// Fatal function
// Log a fatal to system log
// and to file.
func Fatal(content string) {
	if logSetting.logToSys {
		log.SetPrefix("[FATAL]")
		log.SetFlags(log.LstdFlags)
		log.Fatalln(content)
	}

	if logSetting.logToFile {
		logger, f := getFileLogger(FATAL) //.Fatalln(content)
		logger.Fatalln(content)
		f.Close()
	}
}

// Error function
// Log a error to system log
// and to file. Use Error would let system
// into panic status, system need a recover
// manipulation to reset system status.
func Error(content string) {
	if logSetting.logToSys {
		log.SetPrefix("[ERROR]")
		log.SetFlags(log.LstdFlags)
		log.Panicln(content)
	}

	if logSetting.logToFile {
		logger, f := getFileLogger(ERROR)
		logger.Panicln(content)
		f.Close()
	}
}

// Warn function
// Log a warn to system log
// and to file. Use Error would let system
// into panic status, system need a recover
// manipulation to reset system status.
func Warn(content string) {
	if logSetting.logToSys {
		log.SetPrefix("[WARN]")
		log.SetFlags(log.LstdFlags)
		log.Panicln(content)
	}

	if logSetting.logToFile {
		logger, f := getFileLogger(WARN) //.Panicln(content)
		logger.Panicln(content)
		f.Close()
	}
}

// Info function
// Log a Info to system log and file.
func Info(content string) {
	if logSetting.logToSys {
		log.SetPrefix("[INFO]")
		log.SetFlags(log.LstdFlags)
		log.Println(content)
	}

	if logSetting.logToFile {
		logger, f := getFileLogger(INFO) //.Println(content)
		logger.Println(content)
		f.Close()
	}
}

// Debug function
// Log a Debug to system log and file.
// It is output only when logSetting.isDebug is true.
func Debug(content string) {
	if logSetting.isDebug == false {
		return
	}

	if logSetting.logToSys {
		log.SetPrefix("[DEBUG]")
		log.SetFlags(log.LstdFlags)
		log.Println(content)
	}

	if logSetting.logToFile {
		logger, f := getFileLogger(DEBUG) //.Println(content)
		logger.Println(content)
		f.Close()
	}
}
