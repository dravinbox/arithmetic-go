package sort

type BubbleSort struct{
	Base
}

/**
	执行冒泡排序
 */
func (sort BubbleSort) Sort(arr []int)  {
	for i := len(arr)-1; i > 0; i-- {
		sort.bubble(arr,0,i)
	}

}


/**
	执行一次冒泡
 */
func (sort BubbleSort) bubble(arr []int,start_index int,end_index int){
	for i := start_index; i < end_index; i++ {
		if arr[i] > arr[i+1] {
			sort.Swap(arr,i,i+1)
		}

	}

}





