package main

import "fmt"

func testCopy(){

	var a []int = []int{1,2,3,4,5}

	b := make([]int,10) //[0,0,0,0,0,0,0,0,0,0]

	copy(b,a) //内置方法copy

	fmt.Println(b)
}

func main() {
	testCopy()

}
