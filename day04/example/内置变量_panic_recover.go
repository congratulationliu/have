package main

import (
	"time"
	_ "fmt"
	"errors"
)

func getConfig() (err error){
	return errors.New("init config failed!") //创建异常
}

func test(){
	/*
	//通过recover捕获异常程序不退出
	defer func(){
		if err := recover(); err != nil{
			fmt.Println(err)
		}
	}()

	b := 0
	a := 100 / b
	fmt.Println(a)
	*/
	err := getConfig()
	if err != nil{
		panic(err) //通过panic抛出异常与python中的raise一样
	}
	return

}

func main() {
	for {
		test()
		time.Sleep(time.Second * 1)
	}

}