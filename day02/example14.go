package main

import "fmt"

func reverse(str string) (string){
	var result string
	strLen := len(str)
	for i := 0 ; i < strLen; i++ {
		result = result + fmt.Sprintf("%c",str[strLen -i -1])
	}
	return result
}

func reverse_arry(str string) (string){
	var result []byte //定义一个空数组
	length := len(str)
	for i := 0 ; i < length ; i++ {
		result = append(result,str[length -i -1]) //通过append往数组里追加值
	}
	return string(result) //最后在string化
}

func main() {
	myStr:= "hello world"
	result := reverse(myStr)
	fmt.Println(result)
	result = reverse_arry(result)
	fmt.Println(result)
}
