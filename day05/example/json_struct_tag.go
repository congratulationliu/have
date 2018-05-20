package main

import (
	"fmt"
	"encoding/json"
)

//结构体Student必须大写,不然json包无法使用结构体里的字段
type Student struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Score int `json:"score"`

}


func main() {
	var stu1 Student
	stu1.Name = "alex"
	stu1.Age = 18
	stu1.Score = 200

	data, err := json.Marshal(stu1) //默认json序列化为byte数组
	if err != nil{
		fmt.Println("json error",err)
		return
	}
	fmt.Println(string(data))


}
