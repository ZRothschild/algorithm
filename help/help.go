package help
/*
 帮助函数
*/

// 交换函数
func Swap(x ,y int) (int, int)  {
	var tmp int
	tmp = x
	x = y
	y = tmp
	return x,y
}