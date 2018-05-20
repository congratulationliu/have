package main

import (
	"bufio"
	"os"
	"fmt"
)

func count(str string) (wordCount,spaceCount,intCount,otherCount int){
	
	n := []rune(str)
	for _,v := range n{
		switch  {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			wordCount++

		case v >= '0' && v <= '9':
			intCount++

		case v == ' ':
			spaceCount++

		default:
			otherCount++
		}
	}
	return
	
}

func main() {
	reader := bufio.NewReader(os.Stdin) //读一行
	result,_,err := reader.ReadLine() //返回一行

	if err != nil{
		fmt.Printf("read from console err:",err)
	}
	wc,sc,ic,oc := count(string(result))

	fmt.Printf("wordCount=%d\nspaceCount=%d\nintCount=%d\notherCount=%d",wc,sc,ic,oc)

}