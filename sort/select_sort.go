package sort

type SelectSort struct {
	Base
}


/**
求最小下标
 */
func (sort SelectSort) MinIndex(arr []int,startIndex int,endIndex int) int{
	minIndex:=startIndex
	for i := startIndex; i <= endIndex; i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}

	return  minIndex

}

/**
执行
 */
func (sort SelectSort) Sort(arr []int)  {
	for i := 0; i< len(arr)-1; i++ {
		minIndex := sort.MinIndex(arr,i,len(arr)-1)
		if minIndex!=i {
			sort.Swap(arr,i,minIndex)
		}

	}

}


