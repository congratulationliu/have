package main

import "fmt"

func main() {
	var a int = 100
	var b bool
	c := 'a'

	fmt.Printf("%v\n",a) //相应值的默认格式。在打印结构体时
	fmt.Printf("%#v\n",b)
	fmt.Printf("%T\n",c) //相应值的类型的Go语法表示
	fmt.Printf("90%%\n") //百分比
	fmt.Printf("%b\n",a) //二进制显示
	fmt.Printf("%f\n",199.22) //浮点数显示
	fmt.Printf("%q\n","this is a str") //单引号围绕的字符字面值，由Go语法安全地转义
	fmt.Printf("%p\n",&a) //打印内存地址

	//将一个int类型转化成str 不能用强制转换

	str := fmt.Sprintf("%d",a)
	fmt.Printf("%q\n",str)

}

