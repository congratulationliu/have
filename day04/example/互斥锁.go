package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var lock sync.Mutex //互斥锁

func testMutx(){

	//go build --race 查看有没有竞争关系,编译后执行返回
	//互斥锁同一时间只有一个执行
	var a map[int]int
	a = make(map[int]int,100)

	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	for i := 0; i < 50 ; i++{
		go func(b map[int]int){
			lock.Lock()
			b[0] = rand.Intn(100)
			lock.Unlock()
		}(a)
	}
	lock.Lock()
	fmt.Println(a)
	lock.Unlock()

	time.Sleep(time.Second * 1)
}

func main() {

	testMutx()

}
