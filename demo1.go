package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")
	go longWait()
	go shortWait()
	fmt.Println("main() 函数开始等待")
	time.Sleep(10 * 1e9)
	fmt.Println()
	fmt.Println("main() 函数结束等待")
}

func longWait() {
	fmt.Println("开始 longWait()")
	time.Sleep(5 * 1e9) // 单位是 纳秒 ns 此处是5秒
	fmt.Println("结束 longWait()")
}

func shortWait() {
	fmt.Println("开始 shortWait()")
	time.Sleep(2 * 1e9)
	fmt.Println("结束 shortWait()")
}
