package main

import (
	"time"
	"fmt"
)

const (
	Man = 1
	Female = 2
)



func main() {

	Second :=  time.Now().Unix()
	if (Second % Female == 0){
		fmt.Println("female",Female)
	}else {
		fmt.Println("man",Man)
	}

}