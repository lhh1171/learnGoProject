package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"unsafe"
)

func MySyscall2() {
	test4()
}

func test1() {
	if os.Getppid() != 1 {
		//判断当其是否是子进程，当父进程return之后，子进程会被 系统1 号进程接管
		/*Abs 返回路径的绝对表示。如果路径不是绝对路径，它将与当前工作目录一起将其转换为绝对路径。
		不保证给定文件的绝对路径名是唯一的。 Abs 在结果上调用Clean 。*/
		filePath, _ := filepath.Abs("./run.sh") //将命令行参数中执行文件路径转换成可用路径
		cmd := exec.Command(filePath)
		//将其他命令传入生成出的进程
		cmd.Stdin = os.Stdin //给新进程设置文件描述符，可以重定向到文件中
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Start() //开始执行新进程，不等待新进程退出
		return
	}
}

func test2() {
	if os.Getppid() != 1 {
		//将命令行参数中执行文件路径转换成可用路径
		filePath, _ := filepath.Abs("./run.sh")
		args := append([]string{filePath})
		os.StartProcess(filePath, args, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
		return
	}
}

func test3() {
	if os.Getppid() != 1 {
		//将命令行参数中执行文件路径转换成可用路径
		filePath, _ := filepath.Abs("./run.sh")

		args := append([]string{filePath})

		attr := &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}}

		sysattr := &syscall.ProcAttr{
			Dir: attr.Dir,
			Env: attr.Env,
			Sys: attr.Sys,
		}
		if sysattr.Env == nil {
			sysattr.Env = syscall.Environ()
		}
		for _, f := range attr.Files {
			sysattr.Files = append(sysattr.Files, f.Fd())
		}
		syscall.StartProcess(filePath, args, sysattr)
		return
	}
}

func test4() {
	if os.Getppid() != 1 {
		argv0 := "./run.sh"

		//ProcAttr 包含将应用于由StartProcess启动的新进程的属性
		attr := &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}}

		sysattr := &syscall.ProcAttr{
			Dir: attr.Dir,
			Env: attr.Env,
			Sys: attr.Sys,
		}
		if sysattr.Env == nil {
			sysattr.Env = syscall.Environ()
		}
		for _, f := range attr.Files {
			sysattr.Files = append(sysattr.Files, f.Fd())
		}

		argv0p, err := syscall.BytePtrFromString(argv0)
		if err != nil {
			return
		}

		envvp, err := syscall.SlicePtrFromStrings(attr.Env)
		if err != nil {
			return
		}

		syscall.RawSyscall(syscall.SYS_EXECVE,
			uintptr(unsafe.Pointer(argv0p)),
			0,
			uintptr(unsafe.Pointer(&envvp[0])))
		return
	}
}
