package myapi

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MyLineFilter() {

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	/*Err 返回Scanner遇到的第一个非 EOF 错误*/
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
