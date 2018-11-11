package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var Log *log.Logger

func init() {
	Log = log.New(os.Stdout, "godfs ", log.Ldate|log.Ltime|log.Lshortfile)
}

const (
	info    string = "[INFO]"
	debug   string = "[DEBUG]"
	error   string = "[ERROR]"
	warning string = "[WARNING]"
)

func Info(format string, vars ...interface{}) {
	print(format, info, vars)
}

func Debug(format string, vars ...interface{}) {
	print(format, debug, vars)
}

func Error(format string, vars ...interface{}) {
	print(format, error, vars)
}

func Warning(format string, vars ...interface{}) {
	print(format, warning, vars)
}

func print(format string, level string, vars []interface{}) {
	line := format
	for _, o := range vars {
		line = strings.Replace(line, "{}", fmt.Sprint(o), 1)
	}
	Log.Output(3, level+" "+line)
}
