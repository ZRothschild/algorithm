package search

import "fmt"

/*
线性检索

	arr := []int{2,7,1,11,4,5,15,10,14}
	res := search.LinearSearch(arr,11)
	fmt.Printf("%v\n",res)
*/
func LinearSearch(arr []int, x int) int {
	var i, n int = 0, len(arr)

	for i < n && x != arr[i] {
		i++
	}
	if i < n {
		return i
	} else {
		return -1
	}
}

/*
前提已经排过序的数组
二叉检索(也叫做二分查找)
*/
func BinarySearch(arr []int, x int) int {
	low, high, j := 0, len(arr)-1, -1
	for low <= high && j < 0 {
		mid := (low + high) / 2
		if x == arr[mid] {
			j = mid
		} else if x < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return j
}

/*
前提已经排过序的数组 (递归)
二叉检索(也叫做二分查找)
*/
func BinarySearchRec(arr []int, min, max, x int) int {
	mid := (min + max) / 2
	fmt.Println(mid)
	if min > max {
		return -1
	}
	if arr[mid] == x {
		return mid
	} else if arr[mid] < x {
		return BinarySearchRec(arr, mid+1, max, x)
	} else {
		return BinarySearchRec(arr, min, mid-1, x)
	}
}
