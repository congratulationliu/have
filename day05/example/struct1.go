package main

import (
	"fmt"
)
//结构体所有字段的内存布局都是连续的
type Student struct {
	name string
	age int //int占用8字节
	score float32
}

func main() {
	var stu Student
	stu.age = 18
	stu.name = "alex"
	stu.score = 100
	fmt.Printf("name:%s,age:%d,score:%d\n",stu.name,stu.age,stu.score)
	//在内存中的布局
	fmt.Printf("Name:%p\n",&stu.name)
	fmt.Printf("Age:%p\n",&stu.age)
	fmt.Printf("Score:%p\n",&stu.score)


	// 初始化,可以初始化部分字段
	var stu1 *Student = &Student{
		name:"xiaogang",
		age:29,
		score:1000,
	}
	fmt.Println(stu1)

	var stu2 = Student{
		name:"test",
		age:111,
	}

	fmt.Println(stu2)
	var stu3 *Student = new(Student)//创建了个内存地址
	fmt.Println(stu3)

}
