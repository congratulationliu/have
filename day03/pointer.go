package main

import "fmt"

func test(p *int ){
	fmt.Println(p)
	*p = 10000
	return
}

func main() {
	var a int = 10
	fmt.Println(&a) //获取内存地址

	var p *int //定义一个指针
	p = &a
	*p = 20
	fmt.Println(*p) //打印指针对应的数值
	fmt.Println(a)

	var b int = 999
	p = &b  //将指针p的内存地址改为了b,修改指针p的值对b无影响
	*p = 5

	fmt.Println(a)
	fmt.Println(b)

	test(&a) //将内存地址传如函数,函数内修改指针对应的值,影响a
	fmt.Println(a)

}
