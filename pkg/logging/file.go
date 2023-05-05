package logging

import (
	"fmt"
	"go-gin/pkg/setting"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func Setup() {
	if setting.AppSetting.LogSavePath != "" {
		LogSavePath = setting.AppSetting.LogSavePath
	}

	if setting.AppSetting.LogSaveName != "" {
		LogSaveName = setting.AppSetting.LogSaveName
	}

	if setting.AppSetting.LogFileExt != "" {
		LogFileExt = setting.AppSetting.LogFileExt
	}

	if setting.AppSetting.TimeFormat != "" {
		TimeFormat = setting.AppSetting.TimeFormat
	}
}

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}
func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkdir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}
	return handle
}

func mkdir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
