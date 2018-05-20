package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	fmt.Scanf("%s",&a)

	number, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("this is not convert ",err)
	}
	fmt.Println(number)
}