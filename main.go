package main

import (
	"./list"
)

func main ()  {
	arr := []int{2,7,4,8,10,4,6}
	var l list.LNode   //{[] 0}
	l.CreateListR(arr).Sort().DispList()
}