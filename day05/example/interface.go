package main

import "fmt"

//接口是方法都集合,不能设置变量    用途类似于python的抽象类
type Test interface {
	print()
}

type Cart struct {
	name string
	speed int
}

//cart实现了print方法,可以通过接口直接调用
func (self *Cart) print() {
	fmt.Println(self.speed)
	fmt.Println(self.name)
}


func main() {
	var t Test //接口是一个地址
	var a Cart = Cart{
		name:"baoshijie",
		speed:100,
	}

	t = &a  //接口代表了具体都类型
	t.print()


}
