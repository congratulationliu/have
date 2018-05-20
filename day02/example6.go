package main

import (
	"fmt"
)

func modify(a int){
	a = 10 //函数属于值类型
	return
}
func modify1(p *int) {
     *p = 100
	//指针是引用类型 所有会改变
}
func main() {
	a := 5
	b := make(chan int,1)
	fmt.Println("a= ",a)
	fmt.Println("b= ",b)
	modify(a)
	fmt.Println("a= ",a)
	modify1(&a)
	fmt.Println("a= ",a)
}
