package main

import "fmt"

func perfect(n int) bool{
	var sum int = 0
	for i := 1;i < n;i++{
		if n%i == 0{ //自己取余为0
			sum += i
		}

	}
	return sum == n
}

func process(n int){
	for i := 1; i < n+1; i++{
		if perfect(i){
			fmt.Println(i)
		}
	}
}

func main() {
	var input int
	fmt.Scanf("%d",&input)

	process(input)
}