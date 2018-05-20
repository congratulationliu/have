package main

import "fmt"

//插入排序 从一个有序的列表里把一个元素插入进去
func insertSort(a []int){ //一个数组是有序的,另外一个数组是无序的,无序的要去有序里比大小
	for i := 1;i < len(a);i++{ //这是遍历无序的列表
		for j := i; j > 0 ;j--{ //这是遍历有序的列表 进行比较插入 ,有序的列表是插入形式,依次增加
			if (a[j] > a[j-1]){ //当前值 大于上一个值就终止
				break
			}
			a[j],a[j-1] = a[j-1],a[j]
			fmt.Println(a)
		}
	}

}

func main() {
	slice := []int{8,7,5,4,3,10,15} //切片
	insertSort(slice)
	fmt.Println("finally: ",slice)
}

/*
7 8 5 4 3 10 15]
[7 5 8 4 3 10 15]
[5 7 8 4 3 10 15]
[5 7 4 8 3 10 15]
[5 4 7 8 3 10 15]
[4 5 7 8 3 10 15]
[4 5 7 3 8 10 15]
[4 5 3 7 8 10 15]
[4 3 5 7 8 10 15]
[3 4 5 7 8 10 15]
finally:  [3 4 5 7 8 10 15]

*/
