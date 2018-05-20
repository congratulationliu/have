package main

import "fmt"

func reverse(n string) (result bool){

	t := []rune(n)
	length := len(t)
	for i,_ := range t{
		if i == length/2{ //等于中间的下角标就跳出 例如1221 -->到12就停止
			break
		}

		last := length -i - 1
		if t[i] != t[last]{
			result = false
		}else{
			result = true
		}
	}
	return
}


func main() {
	var input string
	fmt.Scanf("%s",&input)

	res := reverse(input)
	if res{
		fmt.Println("ok")
	}else{
		fmt.Println("no")
	}
}
