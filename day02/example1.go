package main

import (
	"fmt"
)

func init(){
	fmt.Println("加载init")
}
//两个数加和,遍历一个数然后另外一个数就是和减去当前值
func list(n int){
	for i := 0; i <= n ; i++ {
		fmt.Printf("%d+%d=%d\n",i,n -i ,n)
	}
}
func main() {
	list(10)
}
