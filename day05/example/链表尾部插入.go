package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name string
	Age int
	Score float32
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
func tailInsertChain(p *Student){

	for i := 0 ; i < 11 ; i++ {
	var stu = Student{
		Name:fmt.Sprintf("student%d",i),
		Age:rand.Intn(100),
		Score:rand.Float32() * 100,
	}
	p.next = &stu
	p = &stu
	}
}

func main() {
	//生成链表表头
	var head Student
	head.Name = "alex"
	head.Age = 18
	head.Score = 100

	//链表尾部插入法
	tailInsertChain(&head)
	trans(&head)

	//p := &head // var p *Student = &head
}
