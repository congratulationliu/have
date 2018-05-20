package main

import "fmt"

func add_avg (a int,arg ...int) int{ // arg ...int 可变长参数
	sum := a
	for i := 0; i < len(arg) ; i++{
		sum += arg[i]
	}
	return sum
}

func concat(a string, arg ...string) (result string){ //定义完返回值不要在函数内在声明类型
	result = a
	for i := 0 ; i < len(arg) ; i++ {
		result += arg[i]
	}
	return //不需要加返回值
}


func main() {
	sum := add_avg(10,100,1000)
	fmt.Println(sum)

	res := concat("hello"," ","myself")
	fmt.Println(res)

}