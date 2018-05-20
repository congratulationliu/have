package main

import (
	"math/rand"
	"fmt"
	"time"
)

func init(){
	//放到初始化里
	rand.Seed(time.Now().UnixNano()) //以时间戳做为当前随机数的种子,Nano是纳秒
}
func main() {

	for i := 0 ; i < 10 ; i++ {
		a := rand.Int() //获取整数
		fmt.Println(a)
	}
	for i := 0 ;i < 10 ; i++{
		a := rand.Intn(100) //获取100以内的随机整数
		fmt.Println(a)
	}
	for i := 0 ; i < 10 ; i++ {
		a := rand.Float32() //获取浮点数
		fmt.Println(a)
	}
}
