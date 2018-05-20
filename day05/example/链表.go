package main

import "fmt"

type Student struct {
	Name string
	Age int
	Score int
	next *Student
}

//函数化
func trans (p *Student){
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
	fmt.Println()
}

func main() {

	var head Student
	head.Name = "alex"
	head.Age = 18
	head.Score = 100


	var stu2 Student
	stu2.Name = "xiaogang"

	head.next = &stu2
	trans(&head)

	var stu3 Student
	stu3.Name = "xiaoming"


	stu2.next = &stu3
	trans(&head)



	//p := &head // var p *Student = &head



}
