package insertSort

/**
 插入排序

	arr := []int{2,7,1,11,4,5,15,10,14}
	insertSort.InsertSort(arr)
	fmt.Printf("%v\n",arr)
*/

func InsertSort(arr []int)  {
	for i := 1 ; i< len(arr)-1 ; i++ {
		a,j := arr[i],i-1
		for j >= 0 && arr[j] > a  {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = a
	}
}