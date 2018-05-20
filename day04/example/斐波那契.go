package main

import "fmt"

func fab(n int) int{
	if n <= 1{
		return 1
	}
	return fab(n-1) + fab(n-2)

}

func main() {
	for i :=0 ; i < 10 ; i++ {
		fmt.Println(fab(i))
	}


}


/*
1
1
2
3
5
8
13
21
34
55
*/