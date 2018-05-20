package main

import "fmt"

func main() {
	var word string = "hello world ,中国"

	for key,value := range word {
		fmt.Printf("key[%d],value[%c],len[%d]\n",key,value,len([]byte(string(value))))
		//cannot convert value (type rune) to type []byte  --- value是rune字符型

	}
}