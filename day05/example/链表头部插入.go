package main

import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name string
	Age int
	Score float32
	up *Student
}

//函数化
func trans (p *Student){
	for p != nil {
		fmt.Println(*p)
		p = p.up
	}
	fmt.Println()
}

func HeadChain(head **Student){
	for i :=0 ; i < 11; i ++ {
		var stu = Student{
			Name:fmt.Sprintf("student%d",i),
			Age:rand.Intn(100),
			Score:rand.Float32() * 100,
		}
		stu.up = *head
		*head = &stu //head相当于副本
	}

}

func DelNode(node *Student){
	//删除节点
	var prev_node *Student = node //临时变量保留上一个节点
	for node != nil{
		if (node.Name == "student6"){
			 prev_node.up = node.up //被删除节点上一个节点的的up 指向被删除节点的up
			 break
		}
		prev_node = node //prev是node的上一个节点
		node = node.up
	}
}

func AddNode(node *Student,new_node *Student){
	//插入节点
	for node != nil{
		if (node.Name == "student6"){
			 new_node.up = node.up
			 node.up = new_node
			 break
		}
		node = node.up
	}
}


func main() {
	//生成链表表头
	var head *Student = new(Student) //head是指针  //改变一个变量的地址,传变量的变量的内存地址,要是改变一个指针的地址,将指针的内存地址传入进去
	head.Name = "alex"
	head.Age = 18
	head.Score = 100
	//链表头部插入法
	HeadChain(&head) //传入指针的内存地址
	trans(head)

	//DelNode(head)
	//trans(head)

	var newNode *Student = new(Student)
	newNode.Name = "xiaogang"
	newNode.Age = 20
	newNode.Score = 300
	AddNode(head,newNode)
	trans(head)

	//p := &head // var p *Student = &head
}
