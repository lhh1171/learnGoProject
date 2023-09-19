package main

import (
	"os"
	"os/exec"
	"syscall"
)

func MyExec2() {
	/*LookPath 在 PATH 环境变量命名的目录中搜索名为文件的可执行文件。
	如果文件包含斜线，则直接尝试，不查询 PATH。否则，成功时，结果是绝对路径*/
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	/*Environ 以“key=value”的形式返回代表环境的字符串副本*/
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
