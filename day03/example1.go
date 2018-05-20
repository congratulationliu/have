package main

import (
	"fmt"
	"math"
)


func isPrime(n int ) bool{

	for i := 2 ; i < int(math.Sqrt(float64(n))) ; i++ { //素数是大于1的自然数除了1和他本身不能有其他整除,类似于奇数
		if n % i == 0 {
			return false
		}

	}
	return true
}

func main() {
	var n int
	var m int
	fmt.Scanf("%d%d",&n,&m) //类似于python的input输入,传入Scanf函数是指针,这样就可以确保函数外定义的变量可以修改值

	for i := n ; i < m ; i++ {
		if isPrime(i) == true{
			fmt.Printf("%d\n",i)
			continue
		}
	}


}
