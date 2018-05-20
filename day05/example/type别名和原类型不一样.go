package main

import "fmt"

type interger int

func main() {
	var i interger = 100
	var j int = 1000

	//i = j
	//cannot use j (type int) as type interger in assignment
	j = int(i)
	fmt.Println(j)

}
