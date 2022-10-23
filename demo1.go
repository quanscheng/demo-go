package main

/*

 */

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	// 两个go程分别计算一半, 然后计算和就是结果
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 求和结果送入管道
}
