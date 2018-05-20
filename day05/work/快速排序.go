package main

import "fmt"
//快速排序
func qSort(a []int,left,right int){
	//选中一个值左面的比他小,右面的比他大
	if left >= right{ //判断传入是否合法
		return
	}

	val := a[left]
	k := left

	for i := left + 1 ; i <= right;i++{ //循环当前序列,即使被切割也是切割序列
		if a[i] < val{//val是中间值,如果比val小就要放到val的左面
			a[k] = a[i] //将小的数与中间值的位置交换
			a[i] = a[k+1]//在将小的位置,与k的下一位交换
			a[k+1] = val //将k+1的位置赋值给中间值
			k++ //最终更换了一个位置所以要+1
		}

	}
	//a[k] = val
	qSort(a,left,k-1) //在左面再次排序
	qSort(a,k+1,right)//右面再次排序递归
}

func main() {

	slice := []int{8,7,5,4,3,10,15,20,2} //切片
	qSort(slice,0,len(slice)-1)
	fmt.Println("finally: ",slice)


}
