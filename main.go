package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strings"
)

// Branch 需要监听的分支
var Branch = "refs/heads/dev"

// EventType 监听分支变动的类型
var EventType = "push"

// NonExistent 不存在标志
var NonExistent = -1

func main() {
	//创建默认路由引擎
	r := gin.Default()

	//初始化log配置
	LogInit()

	//v1版本路由组
	v1Group := r.Group("v1")
	{
		//不同仓库可设置不同的脚本,scriptName必须和服务器上该仓库所执行的脚本名字对应(不包含.sh后缀)
		v1Group.POST("/modify/:scriptName", func(context *gin.Context) {
			//获取脚本名字
			scriptName := context.Param("scriptName")

			//获取请求中event类型
			event := context.GetHeader("x-github-event")

			//获取payload
			body := context.Request.Body
			bodyIO, err := ioutil.ReadAll(body)
			if err != nil {
				//记录错误日志
				log.SetPrefix("[error:" + scriptName + "]")
				log.Println(err)
				return
			}

			//判断是否包含分支信息
			existence := strings.Index(string(bodyIO), Branch)

			//判断事件类型及是否存在分支信息
			if event == EventType && existence != NonExistent {
				//运行脚本
				err := Deploy(scriptName)
				if err != nil {
					//记录错误信息
					log.SetPrefix("[error:" + scriptName + "]")
					log.Println(err)
				} else {
					//记录脚本运行成功的payload
					log.SetPrefix("[success:" + scriptName + "]")
					log.Println(string(bodyIO))
				}
			} else {
				//记录错误日志
				log.SetPrefix("[error:" + scriptName + "]")
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
