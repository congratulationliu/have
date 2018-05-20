package main

import "fmt"

type slice struct { //封装成一个结构体
	ptr *[100]int
	len int
	cap int
}

func make1(s slice,cap int) slice{
	s.ptr = new([100]int)
	s.cap = cap
	s.len = 0
	return s
}

func modify(s slice){
	s.ptr[1] = 1000
}

func testSlice(){
	var s1 slice
	s1 = make1(s1,10)
	s1.ptr[0] = 100

	modify(s1)

	fmt.Println(s1.ptr)

}


func main() {

	testSlice()

}
