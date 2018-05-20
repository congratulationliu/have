package main

import "fmt"

func testArray2(){
	var a = [10]int{1,2,3,4}

	b := a[1:3]
	fmt.Printf("%p\n",b)  //验证b切片和数组内存地址一样,证明切片就是指针

	fmt.Printf("%p\n",&a[1])//因为从第一个元素切的所以指针指向第一个内存地址

}

func main() {
	testArray2()
}

/*
0xc42001a0f8
0xc42001a0f8
*/