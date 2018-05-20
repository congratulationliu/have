package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	//结构体是值类型
	//var p person
	//p.name = "alex"
	//p.age = 18
	//p := person{"alex",18}
	//p := person{name:"alex",age:18}
	//p := new(person) //声明了一个人的指针 p的类型为*person
	p := &person{} //结构体指针
	//p.name = "alex"
	//p.age = 18
	(*p).age = 19
	(*p).name = "dragon"
	fmt.Printf("The Person name is %s,age is %d",p.name,p.age)


}
