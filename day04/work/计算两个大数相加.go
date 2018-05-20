package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//思路 将两个数字变成字符串 跟字符0去做差,然后求和
func multi(str1, str2 string) (result string) {

	if len(str1) == 0 && len(str2) == 0 {
		result = "0"
		return
	}

	var index1 = len(str1) - 1 //第一个字符的长度
	var index2 = len(str2) - 1 //第二个字符的长度
	var left int

	for index1 >= 0 && index2 >= 0 {
		c1 := str1[index1] - '0'  //返回的是str
		c2 := str2[index2] - '0'  //返回的是str

		sum := int(c1) + int(c2) + left //数字相加进位
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0' //取余,在与字符0运算
		result = fmt.Sprintf("%c%s", c3, result) //格式化输出字符串
		index1--
		index2--
	}

	for index1 >= 0 { //当第一个字符长于第二个字符
		c1 := str1[index1] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'

		result = fmt.Sprintf("%c%s", c3, result)
		index1--
	}

	for index2 >= 0 { //当第二个字符长于第一个字符
		c1 := str2[index2] - '0'
		sum := int(c1) + left
		if sum >= 10 {
			left = 1
		} else {
			left = 0
		}
		c3 := (sum % 10) + '0'
		result = fmt.Sprintf("%c%s", c3, result)
		index2--
	}

	if left == 1 { //判断最后的进位是否仍然是1
		result = fmt.Sprintf("1%s", result)
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:", err)
		return
	}

	strSlice := strings.Split(string(result), "+") //通过+进行切片
	if len(strSlice) != 2 {
		fmt.Println("please input a+b")
		return
	}

	strNumber1 := strings.TrimSpace(strSlice[0]) //退去空格
	strNumber2 := strings.TrimSpace(strSlice[1])
	fmt.Println(multi(strNumber1, strNumber2))
}
