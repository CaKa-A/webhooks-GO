package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strings"
)

// Branch 需要监听的分支
var Branch = "refs/heads/master"

// EventType 监听分支变动的类型
var EventType = "push"

// NonExistent 不存在标志
var NonExistent = -1

func main() {
	//创建默认路由引擎
	r := gin.Default()

	//初始化log配置
	LogInit()

	//v1版本
	v1Group := r.Group("v1")
	{
		v1Group.POST("/modify", func(context *gin.Context) {
			//获取请求中event类型
			event := context.GetHeader("x-github-event")
			fmt.Println(event)

			//获取payload
			body := context.Request.Body
			bodyIO, err := ioutil.ReadAll(body)
			if err != nil {
				//打印错误信息
				fmt.Println(err)
				//记录错误日志
				log.SetPrefix("[error]")
				log.Println(err)
				return
			}

			//判断是否包含分支信息
			existence := strings.Index(string(bodyIO), Branch)

			//判断事件类型及是否存在分支信息
			if event == EventType && existence != NonExistent {
				//打印成功信息
				fmt.Println("run cmd")
				//运行脚本
				err := Deploy()
				if err != nil {
					//记录脚本运行成功的payload
					log.SetPrefix("[success]")
					log.Println(string(bodyIO))
				} else {
					//记录错误信息
					log.SetPrefix("[error]")
					log.Println(err)
				}
			} else {
				//打印错误信息
				fmt.Println("error occurred")
				//记录错误日志
				log.SetPrefix("[error]")
				log.Println("Error in event type or changed branch")
			}
		})
	}

	//启动
	err := r.Run(":5140")
	if err != nil {
		return
	}
}
