package main

import "fmt"
import "time"

var pipe chan int

func add_sum(a int,b int ) {
	sum := a + b
	time.Sleep(3 * time.Second)
	pipe <- sum


}

func main() {
	pipe = make(chan int ,1)
	go add_sum(1,2)
	c :=<- pipe
	fmt.Println("num=" ,c)
}
