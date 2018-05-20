package main

import "fmt"
type Student struct {
	Name string
	Age int
	Score float32
	left *Student
	right *Student
}


func trans(p *Student){
	if (p == nil){ //如果为空就终止
		return
	}

	//前序遍历
	//fmt.Println(p)
	//trans(p.left)
	//trans(p.right)
	//中序遍历
	//trans(p.left)
	//fmt.Println(p)
	//trans(p.right)
	//后续遍历
	trans(p.left)
    trans(p.right)
	fmt.Println(p)


}

func main() {
	var root *Student = &Student{
		Name:"alex",
		Age:18,
		Score:100,
	}

	var left *Student = &Student{
		Name:"alex_left",
		Age:19,
		Score:200,
	}

	var left1 *Student = &Student{
		Name:"alex_left1",
		Age:19,
		Score:200,
	}

	var right *Student = &Student{
		Name:"alex_right",
		Age:19,
		Score:200,
	}

	var right1 *Student = &Student{
		Name:"alex_right1",
		Age:19,
		Score:200,
	}

	root.left = left
	root.right = right

	left.left = left1
	right.right = right1

	trans(root)
}
