package main

import "fmt"

type op_func func(int ,int ) int //通过type定义了一个函数类型

func add (a,b int) int {  //这个传入和返回值必须和定义的一致
	return a + b
}

func sub (a,b int) int {
	return  a - b

}

func operator(a,b int , op op_func) int {
	return op(a,b)
}

func main() {
	 //c := add
	 //sum := operator(100,200,c)
	 sum := operator(100,200,sub) //直接传入函数名也可以

	 fmt.Println(sum)
}