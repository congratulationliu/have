package main

import "fmt"

func test1(){
	var a [5]int //定义一个数组

	fmt.Println(a)

	//循环数组两种方式
	//方式一
	for i := 0; i < len(a) ; i++{
		fmt.Println(a[i])
	}
	//方式二
	for key,value := range a{
		fmt.Printf("%d=[%d]\n",key,value)
	}


}

func test2(){

	var a [10]int //数组是值变量,赋值之后新的改变,老的不会改变
	b := a
	b[0] = 100
	fmt.Println(a) //[0 0 0 0 0]
}

func test3(a *[5]int){ //变成指针就成为引用类型
	(*a)[0] = 10
}

func main() {
	test1()
	test2()

	var a [5]int
	test3(&a)
	fmt.Println(a)


}