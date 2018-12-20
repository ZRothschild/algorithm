package sort

/**
 插入排序
	arr := []int{2,7,1,11,4,5,15,10,14}
	insertSort.InsertSort(arr)
	fmt.Printf("%v\n",arr)
*/
func InsertSort(arr []int) {
	for i := 1; i < len(arr)-1; i++ {
		a, j := arr[i], i-1
		for j >= 0 && arr[j] > a {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = a
	}
}

/*
冒泡排序算法
算法思想：在第一轮，取第一个元素和第二个元素比较，若第一个元素大于第二个元素，则这个两个元素位置互换;
再取新的第二个元素与第三个元素比较，执行上述同样的动作;如此继续直到第n-1个与第n个比较和交换为止。这样最大的元素将会
交换到第n的位置。
*/
func Bubble(arr []int) []int {
	var (
		i, k int
	)
	for k = len(arr) - 1; k > 0; k-- {
		for i = 0; i < k; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
			}
		}
	}
	return arr
}

/*
改进冒泡排序
	arr := []int{2,7,8,11,12,15,17}
	res := search.BinarySearch(arr,12)
	fmt.Printf("%v\n",res)
*/
func BubbleSort(arr []int) {
	k, i, flag := len(arr)-1, 0, 1
	for flag == 1 {
		k, flag = k-1, 0
		for i = 0; i < k; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				flag = 1
			}
		}
	}
}

/**
选择排序
**/
func Selection(arr []int) []int {
	var (
		i, k int
	)
	index := 0
	for k = 0; k < len(arr); k++ {
		for i = k + 1; i < len(arr); i++ {
			if arr[i] > arr[index] {
				index = i
			}
		}
		arr[k], arr[index] = arr[index], arr[k]
	}
	return arr
}
