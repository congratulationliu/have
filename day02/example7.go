package main

import (
	"fmt"
)

func swap(a *int,b *int){
	tmp := *a
        *a = *b
        *b = tmp
	return

}
func swap1(a int,b int) (int,int) {
	return b,a
}

func test(){
	var a = 100
	fmt.Println(a)
	for i := 0 ; i < 100 ; i++ {  //大括号括起来是语句块
		var b = i * 2
		fmt.Println(b)
	}//括号结尾就被回收了
	if (a > 0) {
		var c int = 100
		fmt.Println(c)
	}


}

func main() {
	first :=100
	second :=200
	//swap(&first,&second)
	first,second = swap1(first,second)
	fmt.Println("first= ",first)
	fmt.Println("second= ",second)

}



