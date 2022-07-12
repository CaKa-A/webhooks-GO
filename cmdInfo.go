package main

import (
	"fmt"
	"os/exec"
)

// Deploy 启动脚本方法
func Deploy() error {
	//创建脚本执行命令
	c := "./deploy.sh"
	cmd := exec.Command("sh", "-c", c)
	if _, err := cmd.Output(); err != nil {
		//打印并返回错误信息
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}
