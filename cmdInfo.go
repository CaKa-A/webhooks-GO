package main

import (
	"os/exec"
)

// Deploy 启动脚本方法
func Deploy(scriptName string) error {
	//创建脚本执行命令
	c := "./" + scriptName + ".sh"
	cmd := exec.Command("sh", "-c", c)
	if _, err := cmd.Output(); err != nil {
		//返回错误信息
		return err
	} else {
		return nil
	}
}
