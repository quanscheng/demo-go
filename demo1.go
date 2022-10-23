package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// 知识点: 接受者 -> 由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
func (m *MyError) Error() string {
	return fmt.Sprintf("at %v , %s", m.When, m.What)
}

func run() error {
	return &MyError{
		time.Now(), "it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
