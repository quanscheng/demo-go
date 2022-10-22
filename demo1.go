package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 实现 Stringer 内的所有方法(String()) 即可 隐式实现 接口 Stringer
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"quan", 23}
	b := Person{Name: "sicheng", Age: 27}

	fmt.Println(a, b)
}
