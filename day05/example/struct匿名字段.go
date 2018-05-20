package main

import "fmt"

type Human struct {
	name string
	age int
	weight int
}
//匿名字段类似python里的继承
type Student struct {
	Human  // 匿名字段，那么默认Student就包含了Human的所有字段
	speciality string
}

func main() {
	//初始化一个学生
	mark := Student{Human{"alex",18,180},"python"}
	fmt.Println(mark)
	//打印
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)

	//修改这个学生的爱好
	mark.speciality = "golang"
	fmt.Println(mark)

	mark.age += 2
	fmt.Println(mark)
	mark.Human = Human{"dragon",33,190} //student可以.Human 直接修改
	fmt.Println(mark)
}

/*
我们看到Student访问属性age和name的时候，就像访问自己所有用的字段一样，对，匿名字段就是这样，能够实现字段的继承。是不是很酷啊？
还有比这个更酷的呢，那就是student还能访问Human这个字段作为字段名。请看下面的代码，是不是更酷了。
*/