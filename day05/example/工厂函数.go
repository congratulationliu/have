package main

import "fmt"
//go中没构造函数,可以通过工厂函数实现

type Student struct {
	Name string
	Age int
}

func NewStudent(name string,age int) *Student{
	res := new(Student)
	res.Name = name
	res.Age = age
	return res
	//return &Student{Name:name,Age:age}
}

func main() {
	S := NewStudent("alex",18)
	fmt.Println(S.Name,S.Age)
}
