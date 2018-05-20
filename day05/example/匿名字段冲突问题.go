package main

import "fmt"

type Cart1 struct {
	name string
	age int
}

type Cart2 struct {
	name string
}

type Train struct {
	Cart1
	Cart2
}


func main() {
	//cart1,cart2都包含name字段,需要精确选择
	var tra1 Train
	tra1.Cart1.name = "alex"
	fmt.Println(tra1)

}
