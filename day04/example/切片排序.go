package main

import (
	"sort"
	"fmt"
)

func testIntSort(){
	var a = [...]int{22,33,1,3,4} //数组是不能排序的,因为他是值类型


	sort.Ints(a[:]) //将传入的值变为切片
	fmt.Println(a)

}

func testStringSort(){
	var a = [...]string{"aa","ccc","bbb","ffff"}

	sort.Strings(a[:])
	fmt.Println(a)

}
func testFloatSort(){
	var a = [...]float64{0.38,0.02,0.01,0.44}

	sort.Float64s(a[:])
	fmt.Println(a)

}

func testIntSearch(){
	var a = [...]int{22,33,1,3,4}

	sort.Ints(a[:])
	index := sort.SearchInts(a[:],22) //即便你不排序,默认也会排序,但是值会混乱
	fmt.Println(index)
}


func main() {
	testIntSort()
	testStringSort()
	testFloatSort()
	testIntSearch()


}
