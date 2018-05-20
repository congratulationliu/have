package main

import "fmt"

func main() {

	switch  a :=5; { //直接赋值
 		case a >0 && a < 10:
			fmt.Println(" a > 0 & a < 10")

		case a >10 && a < 20:
			fmt.Println("a > 10 & a < 20")

		default:
			fmt.Println("default")
	}

}
