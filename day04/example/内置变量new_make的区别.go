package main

import "fmt"

func mytest(){

	a := new([]int) //创建一个容量为0的切片,a为内存地址,当前为nil不能放值

	b := make([]int,10)

	fmt.Println(a)
	fmt.Println(b)


	//将切片插入值
	*a = make([]int,5)  //通过make初始化容量为5
	(*a)[0] = 100 //panic: runtime error: index out of range 因为空切片需要初始化
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	mytest()
}