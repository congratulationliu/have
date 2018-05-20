package main

import (
	"fmt"
	"math/rand"
)

func main() {

	num := rand.Intn(100)


	for {
		var input int
		fmt.Scanf("%d\n",&input)
		flag := false
		switch  {
			case num == input:
				fmt.Println("you are right")
				flag = true
			case input < num:
				fmt.Println("less")
			case input > num:
				fmt.Println("bigger")

		}
		if flag{
			break
		}
	}

}