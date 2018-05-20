package main

import "fmt"



func main() {
	str := "hello"
	str1 := "world"

	//result := str + " " +str1
	result := fmt.Sprintf("%s %s", str, str1) //字符串拼接
	q := result[6:] //字符串切片
	fmt.Println(result)
	fmt.Println(q)
}
