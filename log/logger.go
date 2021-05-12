package log

import (
	"io"
	"log"
	"os"
)

var (
	INFO *log.Logger
	WARNING *log.Logger
	ERROR *log.Logger
	DEBUG *log.Logger
)

func LogInit() (error) {
	logFile, err := os.OpenFile("api.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	errFile, err := os.OpenFile("error.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	INFO = log.New(io.MultiWriter(logFile, os.Stdout), "[INFO]", log.Ldate|log.Ltime|log.Lshortfile)
	WARNING = log.New(logFile, "[WARNING]", log.Ldate|log.Ltime|log.Lshortfile)
	ERROR = log.New(io.MultiWriter(logFile, errFile), "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)
	DEBUG = log.New(logFile, "[DEBUG]", log.Ldate|log.Ltime|log.Lshortfile)
	return nil
}

func Info(format string, v ... interface{}) {
	INFO.Printf(format, v...)
}

func Warning(format string, v... interface{}) {
	WARNING.Printf(format, v...)
}

func Debug(format string, v... interface{}) {
	DEBUG.Printf(format, v...)
}

func Error(format string, v... interface{}) {
	ERROR.Printf(format, v...)
}
