package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

func logFile() *os.File {
	file, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Can't create log file")
		return nil
	}
	return file
}

func Logging() *logrus.Logger {
	file := logFile()
	logger := &logrus.Logger{
		Out:          file,
		Formatter:    &logrus.JSONFormatter{},
		ReportCaller: true,
		Level:        logrus.ErrorLevel,
	}
	return logger
}
