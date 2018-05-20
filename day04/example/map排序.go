package main

import (
	"fmt"
	"sort"
)

func testMapSort(){
	//保证排序正常
	var a map[int]int
	a = make(map[int]int,5)

	a[0] = 1
	a[2] = 3
	a[3] = 4
	a[10] =100

	fmt.Println(a)

	var keys []int
	for k, _ := range a{
		keys = append(keys,k) //append会直接创建不需要make
	}
	sort.Ints(keys)
	for _,v := range keys{
		fmt.Println(a[v])
	}

}

func main() {
	testMapSort()
}
