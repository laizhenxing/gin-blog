package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

var (
	F *os.File

	DefaultPrefix = ""
	DefaultCallerDepth = 2

	logger *log.Logger
	logPrefix = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

func init()  {
	filePath := GetLogFileFullPath()
	F = OpenLogFile(filePath)

	// log.LstdFlags 日志记录的格式属性
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// 设置日志信息的前缀
func setPrefix(level Level)  {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		// [level][file:line]
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		// [level]
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}

func Debug(v ...interface{})  {
	setPrefix(DEBUG)
	logger.Println(v)
}

func Info(v ...interface{})  {
	setPrefix(INFO)
	logger.Println(v)
}

func Warn(v ...interface{})  {
	setPrefix(WARN)
	logger.Println(v)
}

func Error(v ...interface{})  {
	setPrefix(ERROR)
	logger.Println(v)
}

func Fatal(v ...interface{})  {
	setPrefix(FATAL)
	logger.Fatal(v)
}