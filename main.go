package main

import (
	"./search"
	"fmt"
)

func main() {

	arr := []int{1, 2, 5, 8, 9}
	j := search.BinarySearchRec(arr, 0, 4, 10)
	fmt.Println("aaa", j)
}
