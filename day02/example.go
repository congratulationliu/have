package main

import "fmt"

//func f() (r int) {
//    t := 5
//    defer func() {
//        t = t + 5
//    }()
//    return
//}

//func f() (r int) {
//
//    t := 5
//
//    r = t //赋值指令
//
//    func() { //defer被插入到赋值与返回之间执行，这个例子中返回值r没被修改过
//
//        t = t + 5
//
//    }()
//    return //空的return指令
//
//}



//func f() (r int) {
//    defer func(r int) {
//        r = r + 5
//    }(r)
//    return 1
//}

func f() int{
	var i int
	defer func(a int){
		fmt.Println("i f6",a)
		i += 5
	}(i)
	i = 1
	i++
	return i
}

func main() {
	a := f()
	fmt.Println(a)
}
