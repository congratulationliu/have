package main

import (
	"package_exp/calc"
	"fmt"
)

func main() {
	sum := calc.Add(100,200)
	sub := calc.Sub(100,200)
	fmt.Println("sum=",sum,"sub=",sub)
}


//设置export GOPATH=/Users/dragon/PycharmProjects/go_project/day01 包放在src下