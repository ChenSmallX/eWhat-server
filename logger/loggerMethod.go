package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// inner init function
// called by outter Init function
// paramater:
func init() {

}

func isInit() bool {
	if isInitialized {
		return true
	}
	log.Fatalln("Fatal error: logger no initialized")
	return false
}

func checkDate() {
	var now time.Time
	now = time.Now()
	logDate.year = now.Year()
	logDate.month = int(now.Month())
	logDate.day = now.Day()
}

func getLogFile(theType logType) *os.File {
	// check the date correct
	checkDate()

	// get file name
	var fileName string
	fileName = fmt.Sprintf("./log/%s.%d.%d.%d.log", theType, logDate.year, logDate.month, logDate.day)

	// check the file is exist
	var fileExist bool
	fileExist = false
	if fileName == logFileName {
		var err error
		fileExist, err = pathExist(fileName)
		if err != nil {

		}
	} else {
		logFileName = fileName
		os.Mkdir("./log", 0777)
		file, err := os.Create(fileName)
		if err != nil {

		}
		file.Close()
		fileExist, err = pathExist(fileName)
	}

	// open file and return
	if fileExist {
		openFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {

		}
		return openFile
	}
	return nil
}

func pathExist(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func getFileLogger(theType logType) (*log.Logger, *os.File) {
	logFile := getLogFile(theType)
	logger := log.New(logFile, string(theType), log.LstdFlags)
	return logger, logFile
}

// func closeFileLogger(logger log.Logger) {
//
// }
