package main

import "fmt"

/**
串的基本操作
*/

var (
// nums = []int64{6, 4, -3, 5, -2, -1, 0, 1, -9}
// nums := []int64{6, 4, 5, 0, 1, -3, -2, -1, -9}
// i, j, tmp = 0, 0, 0
)

// for k, v := range nums {
// 	if v >= 0 {
// 		// nums[i],nums[j] = nums[j],nums[i]
// 		i++
// 		tmp = 0
// 	} else {
// 		j = k
// 		tmp++
// 	}
// }

func main() {
	//var t  = SqString{
	//	Data: []rune{'e','f','h','m'},
	//	Length: 4,
	//}
	var s = SqString{
		Data:   []rune{'a', 'b', 'c', 'd'},
		Length: 4,
	}
	fmt.Printf("%v\n", string(DelStr(s, 3, 1).Data))
}

type SqString struct {
	Data   []rune
	Length int
}

/**
将一个字符串常量赋给串，即生一个rune的串s
var s SqString
StrAssign(&s,"abcdefg")
fmt.Println(string(s.Data))
*/
func StrAssign(s *SqString, str string) {
	for _, v := range str {
		s.Data = append(s.Data, v)
		s.Length++
	}
}

/**
将t串复制给s串
var s SqString
var t  = SqString{
	Data: []rune{'a','b','c','d'},
	Length: 4,
}
StrCopy(&s,t)
fmt.Printf("%s\n",string(s.Data))
*/
func StrCopy(s *SqString, t SqString) {
	for _, v := range t.Data {
		s.Data = append(s.Data, v)
	}
	s.Length = t.Length
}

/**
判断t串与s串相等
var t  = SqString{
	Data: []rune{'a','j','c','d'},
	Length: 4,
}
var s  = SqString{
	Data: []rune{'a','b','c','d'},
	Length: 4,
}
fmt.Printf("%t\n",StrEqual(t,s))
*/
func StrEqual(s SqString, t SqString) bool {
	var b = true
	if s.Length != t.Length {
		b = false
	} else {
		for k, v := range t.Data {
			if s.Data[k] != v {
				b = false
				break
			}
		}
	}
	return b
}

/**
s串的长度
*/
func StrLength(s SqString) int {
	return s.Length
}

/**
s串与t串拼接
var t  = SqString{
	Data: []rune{'a','j','c','d'},
	Length: 4,
}
var s  = SqString{
	Data: []rune{'a','b','c','d'},
	Length: 4,
}
fmt.Printf("%v\n",Concat(s,t))
*/
func Concat(s, t SqString) SqString {
	var z SqString
	z.Length = s.Length + t.Length
	for _, v := range s.Data {
		z.Data = append(z.Data, v)
	}
	for _, v := range t.Data {
		z.Data = append(z.Data, v)
	}
	return z
}

/**
求字串 从i位置开始截取j位
var s  = SqString{
	Data: []rune{'a','b','c','d'},
	Length: 4,
}
fmt.Printf("%v\n",string(SubStr(s,2,1).Data))
*/
func SubStr(s SqString, i, j int) SqString {
	var str SqString
	if i <= 0 || i > s.Length || j < 0 || i+j-1 > s.Length {
		return str
	}
	for k := i - 1; k < i+j-1; k++ {
		str.Data = append(str.Data, s.Data[k])
	}
	str.Length = j
	return str
}

/**
将s1串插入s串的第i(1<= i <= strLength(s))个位置
var t  = SqString{
	Data: []rune{'e','f','h','m'},
	Length: 4,
}
var s  = SqString{
	Data: []rune{'a','b','c','d'},
	Length: 4,
}
fmt.Printf("%v\n",string(InsStr(s,t,2).Data))
*/
func InsStr(s, s1 SqString, i int) SqString {
	var str SqString
	if s.Length < i {
		return str
	}
	var j int
	for j = 0; j < i; j++ {
		str.Data = append(str.Data, s.Data[j])
	}
	for _, v := range s1.Data {
		str.Data = append(str.Data, v)
	}
	for j < s.Length {
		str.Data = append(str.Data, s.Data[j])
		j++
	}
	str.Length = s.Length + s1.Length
	return str
}

/**
将s串从第i位置开始删除j个
var s  = SqString{
	Data: []rune{'a','b','c','d'},
	Length: 4,
}
fmt.Printf("%v\n",string(DelStr(s,2,2).Data))
*/
func DelStr(s SqString, i, j int) SqString {
	var str SqString
	if s.Length > i+j {
		return str
	}
	for k, v := range s.Data {
		if k < i || k > i+j-1 {
			str.Data = append(str.Data, v)
		}
	}
	return str
}

/**
将s串从第i位置开始删除j个
*/
func RepStr(s SqString, i, j int) SqString {
	var str SqString
	for k, v := range s.Data {
		if k < i || k > i+j-1 {
			str.Data = append(str.Data, v)
		}
	}
	return str
}
