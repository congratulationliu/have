package main

import "fmt"

func main() {

	LABLE1:
		for i := 0; i < 5; i++{
			for j :=0 ; j < 5 ;j++{
				if j == 4 { //当j等于4的时候就跳到LABLE1下在执行,不会只跳过j循环
					goto LABLE1
				}
				fmt.Printf("i:[%d],j:[%d]\n",i,j)
			}
		}

}
