package main

import "fmt"

func testArray(){
	var a [5]int = [5]int{1,2,3,4,5}
	var a1 = [5]int{2,3,4,5,6}
	var a2 =  [...]int{1,2,3}
	var a3 = [...]string{1:"hello",3:"world"}
	fmt.Println(a)
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}

//多维数组

func testArray1() {
	var a = [2][5]int{{1,2,3,4,5},{6,7,8,9,0}} //第一个是有几个{},第二个是每个{}有几个元素
	fmt.Println(a)
	var b = [...][5]int{{1,2,3,4,5},{6,7,8,9,0}} //一维可以加...,其他不行
	fmt.Println(b)

	//多维数组遍历
	for row,v := range a{
		for col,v1 := range v{
			fmt.Printf("row:%d,col:%d,val:%d\t",row,col,v1)
		}
		fmt.Println()
	}
}


func main() {
	testArray()
	testArray1()
}

/*
[1 2 3 4 5]
[2 3 4 5 6]
[1 2 3]
[ hello  world]
*/