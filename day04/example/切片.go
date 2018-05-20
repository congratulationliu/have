package main

import "fmt"

func testSlice1() {
	//切片是引用类型,传入函数中会改变原有值
	var slice []int
	var arr [5]int = [5]int{1, 2, 3, 4, 5}

	slice = arr[:]
	slice = slice[1:]
	slice = slice[:len(slice)-1] //刨除最后一个
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = slice[0:1] //切片可以在切片,这样cap容量不变,len会变小
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func main() {

	testSlice1()
}

/*
[2 3 4]
3
4
1
4
*/