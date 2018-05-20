package main

import "fmt"

//def read(){
//	conn,err := openConn()
//
//	defer func(){   //匿名函数与defer结合
//		if conn != nil {
//			conn.Close()
//		}
//	}()
//
//}

var (
	result = func(a,b int) int{  //定义一个全局的匿名函数,前提要有变量接收这个匿名函数
		return a + b
	}
)

func dtest(a,b int) int{
	result := func(a1,b1 int) int{  //定义一个匿名函数
		return a1 + b1
	}(a,b)

	return result

}

func main() {

	res_1 := result(400,800)
	fmt.Println(res_1)

	res := dtest(100,300)
	fmt.Println(res)

	i := 0
	//首先定义defer不会执行,这个时候i=0,之后无论i怎么修改都不会改变(因为fmt的函数传入的是值类型),
	// 然后将defer压入栈中,在整个程序执行完毕后,从栈中依次执行
	defer fmt.Println(i)
	defer fmt.Println("second")


	i = 10
	fmt.Println(10)

}
