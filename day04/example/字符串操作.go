package main

import "fmt"

func testModifyString(){
	var a string = "我hello world"
	b := []rune(a) //有中文的时候需要通过字符数组更改
	b[0] = '1'
	fmt.Println(string(b))


}
func main() {
	testModifyString()
}
