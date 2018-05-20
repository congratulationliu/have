package main

import "fmt"

func isNumber(num int) bool{
	a := num % 10
	b := (num / 10) % 10
	c := (num / 100) % 10
	res := a*a*a + b*b*b + c*c*c
	return res == num
}

func main() {
	var n int
	var m int
	fmt.Scanf("%d%d",&n,&m)
	for i := n ; i < m ; i ++ {
		if isNumber(i) == true{
			fmt.Println("this is a shuixianhua = ",i)
		}
	}
}