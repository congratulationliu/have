package main

import "fmt"

type Cart struct {
	weight int
	length int
}

func (p *Cart) run(speed int) {
	fmt.Println("running speed is ",speed)
}


//struct实现String,类调用方法都时候将变量转为指针,通过接口实现不会

func (p *Cart) String() string{
	str := fmt.Sprintf("[%d]-[%d]",p.length,p.weight)
	return str
}

type Bike struct {
	Cart
	speed int

}

//组合
type Train struct {
	c Cart
}

func main() {
	var a Bike
	a.speed = 100
	a.weight = 200
	a.length = 20000
	fmt.Println(a)

	a.run(100)

	var b Train
	//带着组合都标识请求
	b.c.run(1000)

	//触发了a的String的方法
	fmt.Printf("%s",&a)
}
