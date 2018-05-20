package main

import "fmt"

func main() {
	var a []int
	a = append(a,20,3,40)
	a = append(a,a...) //数组和数组相加可以用数组...
	fmt.Println(a)

}