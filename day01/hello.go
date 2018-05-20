package main

import (
	"fmt"
	//"time"
)

func test_pipe(){
	//创建一个大小为3的管道
	pipe := make(chan int,3)
	pipe <- 1
	pipe <- 2
	pipe <- 3
	var t1 int
	t1 = <- pipe //管道先进先出
	pipe <- 4
	fmt.Println(t1)
	fmt.Println(len(pipe))
}

func main() {
	test_pipe()
	//time.Sleep(1 *time.Second)
}

/*
当放入的数据大于管道大小会编译的时候会产生死锁

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.test_pipe()
	C:/Users/liujiliang/PycharmProjects/go_project/hello.go:14 +0xae
main.main()
	C:/Users/liujiliang/PycharmProjects/go_project/hello.go:19 +0x27
exit status 2
*/