package main

import (
	"fmt"
	"time"
)

/*
如果 2 个协程需要通信，你必须给他们同一个通道作为参数才行。
*/

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
	/*
			尝试一下如果注释掉 time.Sleep(1e9) 会如何。
		我们发现协程之间的同步非常重要：
		1. main() 等待了 1 秒让两个协程完成，如果不这样，sendData() 就没有机会输出。
		2. getData() 使用了无限循环：
			它随着 sendData() 的发送完成和 ch 变空也结束了。
		3. 如果我们移除一个或所有 go 关键字，程序无法运行，Go 运行时会抛出 panic
	*/
}

func sendData(ch chan string) {
	ch <- "张无忌"
	ch <- "赵敏"
	ch <- "周芷若"
	ch <- "谢逊"
	ch <- "张三丰"
}

func getData(ch chan string) {
	var input string

	for {
		input = <-ch
		fmt.Printf("%s\n", input)
	}
}
