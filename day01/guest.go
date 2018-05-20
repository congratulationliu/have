package main

import "fmt"

func calc(a int,b int) (int,int){
	c := a + b
	d := (a + b)/2
	return c, d
}

func main() {
	//sum,avg := calc(10,5) //多个返回值
	sum,_ := calc(100,200) //多个返回值 其中一个不用,采用下划线占位
	//fmt.Println("sum=",sum,"avg=",avg)
	fmt.Println("sum=",sum)


}
