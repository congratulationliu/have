package main

import "fmt"

type interger int

func (p interger) print(){
	fmt.Println("number is ",p)
}

func (p *interger) set(b interger){
	*p = b
}


type Student struct {
	name string
	age int
}

func (self *Student) init(name string,age int){
	self.name = name
	self.age = age
	fmt.Println(self)
}

func (self Student) get() Student{
	return self
}

func main() {

	var stu1 Student
	//go中自动变成指针
	stu1.init("alex",18)
	res := stu1.get()
	fmt.Println(res)


	var myint interger
	myint = 100
	myint.print()
	myint.set(1000)
	myint.print()

}
