package main

import "fmt"

func main() {
	var str = "hello world \n"
	var str1 = `
	go go go ,
	1231 123
	`
	var b  byte = 'c'
	fmt.Println(str)
	fmt.Println(str1)
	fmt.Println(b) //默认是数字标识
	fmt.Printf("%c\n",b) //字符格式化输出

}
