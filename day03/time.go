package main

import (
	"time"
	"fmt"
)

func test(){
	time.Sleep(time.Millisecond * 100)
}

func main() {
	nowTime := time.Now()
	fmt.Println(nowTime.Format("2006/01/02 15:04:05"))

	startTime := time.Now().UnixNano()
	test()
	endTime := time.Now().UnixNano()
	fmt.Printf("cost time is %d us \n",(endTime - startTime) / 1000)


}
