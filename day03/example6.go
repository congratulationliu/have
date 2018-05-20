package main

import (
	"strings"
	"fmt"
	"strconv"
)

func main() {
	str := "hello world world\n"

	result := strings.Index(str,"h") //获取index,首次出现h的位置,获取不到返回-1
	fmt.Println("Index is :",result)

	result = strings.LastIndex(str,"l") //获取str中l最后出现的位置,没有返回-1
	fmt.Println("LastIndex is :",result)

	replaceResult := strings.Replace(str,"world","myself",1) //替换当最后的int为-1时候,替换所有,等于0的时候不替换
	fmt.Println("Replace is :",replaceResult)


	count := strings.Count(str,"l") //求某个字符串出现的次数
	fmt.Println("count is :",count)

	repeatResult := strings.Repeat(str,2) //字符串重复几次
	fmt.Println("repeat_result is :",repeatResult)

	tolowerResult := strings.ToLower(str) //全部变小写
	fmt.Println("tolower_result is :",tolowerResult)

	toupperResult := strings.ToUpper(str) //全部变大写
	fmt.Println("toupper_result is :",toupperResult)

	trimspaceResult := strings.TrimSpace(str) //去掉左右两头空格,默认\n也会去掉
	fmt.Println("trimspace result is :",trimspaceResult)

	trimResult := strings.Trim(str,"\n") //自定义替换方式
	fmt.Println("trims result is :",trimResult)

	trimLeftResult := strings.TrimLeft(str,"\n") //去掉左侧的\n
	fmt.Println("trimLeftspace result is :",trimLeftResult)

	trimRightResult := strings.TrimRight(str,"\n") //去掉右侧的\n
	fmt.Println("trimRightspace result is :",trimRightResult)


	feildResult := strings.Fields(str)  //将字符串以数组形式返回,默认以空格切割
	for i := 0; i < len(feildResult) ; i++ {
		fmt.Println(feildResult[i])
	}

	spiltResult := strings.Split(str,"l") //将字符串以l切割
	for i := 0; i < len(spiltResult) ; i++ {
		fmt.Println(spiltResult[i])
	}

	joinResult := strings.Join(spiltResult,"l") //将数组join到一期
	fmt.Println("join str is ",joinResult)

	myInt := 100
	itoaResult := strconv.Itoa(myInt) //将整数转化为字符串
	fmt.Println("itoa is ",itoaResult)


	atoiResult,err := strconv.Atoi(itoaResult) //将字符串转化为整数
	if err != nil{
		fmt.Println("this is not convert :",err)
	}
	fmt.Println("atoi is :",atoiResult)

}


