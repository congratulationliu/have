package main

import "fmt"

func main() {
	var a int = 100

	switch a {  //switch 不需要加break,跟shell中的case很像
		case 100:
			fmt.Println("this is 0")
			fallthrough //执行后继续往下走 不判断穿透
		case 10:
			fmt.Println("this is 10")

	    case 0,1,2,3,4,5: //多个条件都走一个分支
			fmt.Println("1,2,3,4,5,0 is ok")

		default:
			fmt.Println("this is default")
		
	}
}