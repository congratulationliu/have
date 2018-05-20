package main

import "fmt"

func calc(index string, a, b int) int { // "1" 1  3
    ret := a + b
    fmt.Println(index, a, b, ret) // "1"
    return ret
}

func testHave(){
    a  := [...]int{1,2,3,4,5,6,7,8}
    b := a[2:5:7]
    fmt.Println(cap(b))
}

func main() {
    a := 1
    b := 2
    defer calc("1", a, calc("2", a, b))  //a = 1 b =  2
    a = 0
    defer calc("3", a, calc("4", a, b)) // a = 0 b = 2
    b = 1
    testHave()
}
