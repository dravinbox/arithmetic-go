package sort

//定义排序的公共父类
type Base struct{

}
/**
交换元素
 */
func (base Base) Swap(arr []int, targetIndex int, replaceIndex int){
	var tmp  = arr[targetIndex]
	arr[targetIndex] = arr[replaceIndex]
	arr[replaceIndex] = tmp
}


//定义排序接口
type ISort interface {
	Sort(arr []int)
}




