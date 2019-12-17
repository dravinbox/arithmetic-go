package sort

type InsertionSort struct{
	Base
}

/**
执行插入排序
 */
func (sort InsertionSort) Sort(arr []int)  {
	for i := 0; i< len(arr)-1; i++ {
		sort.insert(arr,0,i,i+1)

	}

}

func (sort InsertionSort) insert(arr []int,start_index int,end_index int,target_index int){
	for i := end_index; i >= start_index; i-- {
		if arr[target_index] < arr[i]{
			sort.Swap(arr,i,target_index)
			target_index = i
		}else{
			break
		}

	}

}
