package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scanf("%s" ,&str)

	var result = 0
	for i := 0 ; i < len(str) ; i++ {
		sum := int(str[i] - '0') //字符串切片默认是ascii码,通过ascii与0做运算 0代表30，1代表31做差仍为1
		result += (sum * sum * sum)
	}

	number,err := strconv.Atoi(str) //将字符串转化为int
	if (err != nil) {
		fmt.Printf("cat not convert %s to int \n",number)
		return
	}

	if result == number{
		fmt.Printf("%d is shuixianhua \n",number)

	}else{
		fmt.Printf("%d is not shuixianhua\n" ,number)
	}

}
