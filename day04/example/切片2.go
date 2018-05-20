package main

import "fmt"

func testArray3(){
	var arr [5]int = [...]int{1,2,3,4,5}


	s := arr[1:3]
	fmt.Printf("before:len[%d],cap[%d]\n",len(s),cap(s))

	s[1] = 100
	fmt.Printf("arr=%p,slice=%p\n",&arr[1],s)
	fmt.Println("before append",arr)
	s = append(s,10)
	s = append(s,10)

	s = append(s,10)
	fmt.Printf("after:len[%d],cap[%d]\n",len(s),cap(s)) //容量翻倍增加
	s = append(s,10)
	s = append(s,10)

	s[1] = 1000
	fmt.Println(s)
	fmt.Println("after append",arr) //a的值还会变,只不过没超出之前改的还是原来的内存地址
	fmt.Printf("arr=%p,slice=%p\n",&arr[1],s) //切片是可变长的,底层会创建一个新的数组,将指针指向新的数组

}

func main() {
	testArray3()
}
