package main

import "fmt"

//选择排序
func selectSort(a []int){
	for i := 0;i < len(a);i++{ //循环整个序列
		var min int = i
		for j:= i + 1; j < len(a);j++{ //循环从1:最后的序列
			if (a[min] > a[j]){ //a[i] > a[j] i为当前值 大于下一个值
				min = j
			}
			fmt.Println(a)
		}
	a[i],a[min] = a[min],a[i] //最小值和当前值交换
	}

}

func main() {
	slice := []int{8,7,5,4,3,10,15} //切片
	selectSort(slice)
	fmt.Println("finally: ",slice)
}

/*
[8 7 5 4 3 10 15]
[8 7 5 4 3 10 15]
[8 7 5 4 3 10 15]
[8 7 5 4 3 10 15]
[8 7 5 4 3 10 15]
[8 7 5 4 3 10 15]
[3 7 5 4 8 10 15]
[3 7 5 4 8 10 15]
[3 7 5 4 8 10 15]
[3 7 5 4 8 10 15]
[3 7 5 4 8 10 15]
[3 4 5 7 8 10 15]
[3 4 5 7 8 10 15]
[3 4 5 7 8 10 15]
[3 4 5 7 8 10 15]
[3 4 5 7 8 10 15]
[3 4 5 7 8 10 15]
[3 4 5 7 8 10 15]
[3 4 5 7 8 10 15]
[3 4 5 7 8 10 15]
[3 4 5 7 8 10 15]
finally:  [3 4 5 7 8 10 15]

*/