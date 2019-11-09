package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	// 文件存储路径
	LogSavePath  = "runtime/logs/"
	// 文件名前缀
	LogSaveName = "log"
	// 文件名后缀
	LogFileExt = "log"
	// 日志格式化
	TimeFormat = "20060102"
)




// 获取当前路径
func GetLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

// 获取绝对路径
func GetLogFileFullPath() string {
	prefixPath := GetLogFilePath()
	filename := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, filename)
}

// 打开文件
func OpenLogFile(filePath string) *os.File {
	// 返回文件信息结构描述文件
	_, err := os.Stat(filePath)
	switch  {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission: %v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile: %v", err)
	}

	return handle
}

// 创建文件夹
func mkDir()  {
	// 获取当前路径
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir + "/" + GetLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}