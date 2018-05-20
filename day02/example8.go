package main
import (
	"fmt"

)


var a = "G"



func main() {
	n()
	m()
	n()

}

func n (){
	fmt.Println(a)
}

func m (){
	a = "O" //将a赋值成O,因为m中未定义a所以修改的是全局的a

	fmt.Printf(a)
}