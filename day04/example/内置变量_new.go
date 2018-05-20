package main

import "fmt"

func main() {
	var i int
	fmt.Println(i)

	j := new(int) //值类型分配内存地址
	fmt.Println(j) //打印的是内存地址
	*j = 100
	fmt.Println(*j) //打印指针对应的值
}


/*
0
0xc420014060
100
*/