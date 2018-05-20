package main
import (
	"fmt"

)


var a string  //定义了个全局变量



func main() {
	a = "G" //局部变量，因为当前未生命a变量,所以修改的是全局的a
	fmt.Println("from main",a) //找自己的局部变量
	f1()
}
func f2()  {
	fmt.Println("from f2",a) //局部变量没有,找全局的a,这是全局的a已经被更改为G
}

func f1(){
	a:="O" //重新在函数内声明var a string = "O"
	fmt.Println("from f1",a)
	f2()//执行f2函数相当于重新引用不会继承自己的局部变量
}
