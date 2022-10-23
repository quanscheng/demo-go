package main

import (
	"fmt"
	"time"
)

/*
他们按照我们期望的顺序打印出了消息，几乎都一样，可是我们明白这是模拟出来的，以并行的方式。
我们让 main() 函数暂停 10 秒从而确定它会在另外两个协程之后结束。
如果不这样（如果我们让 main() 函数停止 4 秒）， main() 会提前结束， longWait() 则无法完成。
如果我们不在 main() 中等待，协程会随着程序的 结束而消亡。
*/

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
