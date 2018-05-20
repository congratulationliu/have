package main

import "fmt"

type Carter interface {
	Run()
	GetName() string
	Didi()
}

type BMW struct {
	name string
}

func (p *BMW) Run(){
	fmt.Println("this bmw is running")
}

func(p *BMW) GetName() string{
	return p.name
}

func (p *BMW) Didi() {
	fmt.Println("this bmw is didi")
}




func main() {
	var car Carter
	isBmw := new(BMW)
	isBmw.name = "liangliang"
	car = isBmw
	car.Run()
	res := car.GetName()
	fmt.Println(res)


}
