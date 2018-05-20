package main

import "fmt"

func fab1(n int){
	var sli []int64  //生成了个切片
	sli = make([]int64,n)
	sli[0] = 1
	sli[1] = 1
	for i :=2 ; i < len(sli);i++{
		sli[i] = sli[i-1] + sli[i-2]
	}

	for _,value := range sli{
		fmt.Println(value)
	}
}

func main() {
	fab1(10)
}