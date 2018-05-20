package main

import (
	"sync"
	//"math/rand"
	"fmt"
	"sync/atomic"
	"time"
)

var lock sync.Mutex
var rwlock sync.RWMutex
//读写锁适用于读多写少的场景
//设置3秒,读写锁执行22万次,互斥锁执行2292次
//go get xx下载第三方包，默认会下载到gopath的src下

func testRwlock(){
	var a map[int]int
	a = make(map[int]int)
	var count int32
	a[0] = 10
	a[1] = 10
	a[2] = 10
	a[3] = 10
	a[4] = 10

	//写操作
	for i := 0; i < 2 ; i++{
		go func(b map[int]int){
			//lock.Lock()
			rwlock.Lock()
			//b[0] = rand.Intn(100)
			time.Sleep(3 * time.Millisecond)
			rwlock.Unlock()
			lock.Unlock()
		}(a)
	}
	//读操作
	for i := 0; i < 100 ; i++{
		go func(b map[int]int){
			//lock.Lock()
			for {
				//lock.Lock()
				rwlock.RLock()
				//fmt.Println(a)
				time.Sleep(time.Millisecond)
				rwlock.RUnlock()
				//lock.Unlock()
				atomic.AddInt32(&count,1) //atomic原子操作,相当于串行
			}
		}(a)
	}

	time.Sleep(time.Second * 3)
	fmt.Println(atomic.LoadInt32(&count)) //原子操作

}

func main() {
	testRwlock()
}