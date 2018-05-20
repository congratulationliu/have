package main

import "fmt"

type Rectangle struct {
	wight,height float64
}

func area(r Rectangle) float64{
	return r.wight*r.height
}

func main() {
	r1 := Rectangle{12,2}

	res := area(r1)
	fmt.Println(res)
}
