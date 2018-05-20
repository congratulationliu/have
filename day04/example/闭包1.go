package main

import (
	"strings"
	"fmt"
)

type str func(string) string

func makeSuffix(suffix string)  str { //make的时候传入suffix

	return func(name string) string{
		if strings.HasSuffix(name,suffix) == false{
			return name + suffix
		}
		return name
	}

}


func main() {
	f1 := makeSuffix(".mp3")

	fmt.Println(f1("bgm"))

	f2 := makeSuffix(".mp4")

	fmt.Println(f2("av"))


}