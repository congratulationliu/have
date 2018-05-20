package main

import "fmt"

func testMap(){
	//声明初始化
	var a map[string]string = map[string]string{
		"asdas":"asdsdasd",
	}
	//make初始化
	//a := make(map[string]string,10)

	a["abc"] = "cde"

	fmt.Println(a)

}

func testMap2(){
	//因为有2层数组,需要初始化两次,才能使用,map,chan,slice都需要初始化,声明为nil
	a := make(map[string]map[string]string,100)
	a["key1"] = make(map[string]string,10)
	a["key1"]["key2"] = "val1"
	a["key1"]["key3"] = "val2"
	a["key1"]["key4"] = "val3"
	a["key1"]["key5"] = "val4"

	fmt.Println(a)

}

func modify2 (a map[string]map[string]string){

	//查找
	_, ok := a["zhangsan"] //固定写法,第一个为值,第二个为bool
	if !ok {
		a["zhangsan"] = make(map[string]string) //判断zhangsan是否存在,不存在就创建zhangsan与对应都map
	}

	/* 等价写法
	if a["zhangsan"] == nil{
		a["zhangsan"] = make(map[string]string)
	}
	*/
	a["zhangsan"]["nickname"] = "pangpang"
	a["zhangsan"]["age"] = "18"

	return

}

func testMap3(){
	a := make(map[string]map[string]string,100)
	modify2(a)
	fmt.Println(a)
}

func trans(a map[string]map[string]string){
	//遍历map
	for k,v := range a{
		fmt.Println(k)
		for k1,v1 := range v{
			fmt.Println("\t",k1,v1)
		}
	}
}

func testMap4(){
	//遍历map
	a := make(map[string]map[string]string,100)
	a["key1"] = make(map[string]string,10)
	a["key1"]["key2"] = "val1"
	a["key1"]["key3"] = "val2"
	a["key1"]["key4"] = "val3"
	a["key1"]["key5"] = "val4"


	a["key2"] = make(map[string]string,10)
	a["key2"]["key2"] = "val1"
	a["key2"]["key3"] = "val2"
	trans(a)
	//删除map中key,map要全部清空for循环,或者重新make下
	delete(a,"key1")
	trans(a)
	//求长度
	fmt.Println(len(a))



}

func testMap5(){
	var a []map[int]int
	a = make([]map[int]int,5)

	if a[0] == nil{  //默认只声明的map,或者slice都是nil
		a[0] = make(map[int]int)
	}

	a[0][10] = 100
	fmt.Println(a)
}

func main() {
	//testMap()
	//testMap2()
	//testMap3()
	//testMap4()
	testMap5()
}
