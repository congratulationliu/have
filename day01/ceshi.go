package main

import (
	"fmt"
	"time"
)

func test_goroute(a int ,b int){
	sum := a + b
	fmt.Println(sum)
}

func main() {
	for i := 0 ; i < 100 ; i ++ {
	  go test_goroute(200,100)
	}
	time.Sleep(2 *time.Second)
}
