package merge

/*
	Merge 合并两个升序的数组

	arr := []int{2,7,9,11,1,5,8,10,14}
	merge.Merge(arr,0,3,8)
	fmt.Printf("%v\n",arr)
*/

func Merge(arr []int,p,q,r int) {
	bp := make([]int,r-p+1,r-p+1)
	i,j,k := p,q+1,0
	for i<= q && j<= r {
		if arr[i] <= arr[j] {
			bp[k] = arr[i]
			k++
			i++
		}else {
			bp[k] = arr[j]
			k++
			j++
		}
	}
	if i == q+1 {
		for ; j <= r; j++ {
			bp[k] = arr[j]
			k++
		}
	}else {
		for ; i <= q ;i++ {
			bp[k] = arr[i]
			k++
		}
	}
	k = 0
	for i=p; i <= r ;i++ {
		arr[i] = bp[k]
		k++
	}
}