package main

import "fmt"

//通过递归求阶乘
func calc(n int) int{
	if n == 1{
		return 1
	}
	return calc(n-1) *n

}

func main() {
	m := calc(5)
	fmt.Println(m)



}
