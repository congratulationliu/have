package main

import "fmt"
var a string = "xxxxx"  //先初始化变量,在初始化init,在执行main
var age int = 100

func init(){
	a  = "xx"
	age =  10
}
func main() {
	fmt.Println("a",a,"age",age)
}