package bubble

/*
冒泡排序算法
算法思想：在第一轮，取第一个元素和第二个元素比较，若第一个元素大于第二个元素，则这个两个元素位置互换;
再取新的第二个元素与第三个元素比较，执行上述同样的动作;如此继续直到第n-1个与第n个比较和交换为止。这样最大的元素将会
交换到第n的位置。
*/
func Bubble (arr []int) []int {
	var (
		i,k int
	)
	for k= len(arr)-1;k > 0 ;k--  {
		for i= 0; i < k ; i++ {
			if arr[i] > arr[i+1] {
				arr[i],arr[i+1] = arr[i+1],arr[i]
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
func BubbleSort (arr []int) {
	k,i,flag := len(arr)-1,0,1
	for flag == 1 {
		k,flag = k-1,0
		for i = 0; i < k; i++ {
			if arr[i] > arr[i+1] {
				arr[i],arr[i+1] = arr[i+1],arr[i]
				flag = 1
			}
		}
	}
}