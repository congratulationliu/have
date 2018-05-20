package main

import "fmt"

func Adder() func (int) int { //类似一个类,x为类变量,修改都生效
	var x int
	return func(d int) int{
		x += d
		return x
	}
}

func main() {
	f := Adder()
	fmt.Println(f(1))
	fmt.Println(f(20))
	fmt.Println(f(100))

}