package main

import (
	"fmt"
	"log"
	"os"
)

// LogInit 配置日志信息
func LogInit() {
	//配置log文件输出地址
	logFile, err := os.OpenFile("./deployLog.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		//打印错误信息
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	//配置log文件记录信息
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Ldate)
}
