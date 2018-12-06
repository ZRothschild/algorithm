package main

import (
	"./BTNode"
	"fmt"
)

func main ()  {
	s := []rune("#ABCDKEFHGNS")
	b := BTNode.Trans(s,1)
	fmt.Print("%v",string(b.RChild.Data))
}