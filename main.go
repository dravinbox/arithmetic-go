package main

import (
	. "github.com/dravinbox/arithmetic/sort"
)

func main(){
	println("hello,world")
	var arr = []int{1,5,3,2,4}
	var iSort ISort

	//选择排序
	//iSort = sort.SelectSort{}
	//iSort.Sort(arr)

	//冒泡排序
	iSort = BubbleSort{}
	iSort.Sort(arr)



	for i := 0; i< len(arr);i++  {
		print(arr[i])

	}
}
