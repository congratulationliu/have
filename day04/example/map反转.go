package main

import (
	"fmt"
)

func testMapReverse(){
	//反转要新生成一个map,需要注意字符类型
	var a map[string]int
	a = make(map[string]int,5)

	var b map[int]string
	b = make(map[int]string,5)

	a["aaa"] = 1
	a["bbb"] = 3

	for k,v := range a{
		b[v] = k
	}
	fmt.Println(b)

}

func main() {
	testMapReverse()
}
