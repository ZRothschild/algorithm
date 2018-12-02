//顺序表 操作
package list

/*
顺序表
*/
type SqList struct {
	Data []int
	Length int
}

type SqLists interface {
	CreateList(arr []int) (sqList *SqList)
	InitList() *SqList
	ListInsert(local,value int) bool
	ListDeleter(local int) bool
	DElNode1(val int) *SqList
	DElNode2(val int) *SqList
}

/*
建立顺序表
	arr := []int{2,7,8,11,12,15,17}
	var l sqList.SqList   //{[] 0}
	//fmt.Println(l)
	//var x *sqList.SqList
	//fmt.Println(x)  //<nil>
	//y := new(sqList.SqList)
	//fmt.Println(y)
	res := l.CreateList(arr)
	fmt.Printf("%v\n",res)
*/
func (l *SqList) CreateList(arr []int) *SqList {
	for _,v := range arr {
		l.Data =  append(l.Data,v)
	}
	l.Length = len(arr)
	return l
}

/*
初始化顺序表
*/
func (l *SqList) InitList() *SqList {
	return l
}

/*
插入顺序表
	arr := []int{2,7,8,11,12,15,17}
	var l sqList.SqList   //{[] 0}
	L := l.CreateList(arr)
	fmt.Printf("re %#v\n",L)
	res := L.ListInsert(3,50)
	fmt.Printf("res %v\n",res)
	fmt.Printf("L %v\n",L)
*/
func (l *SqList) ListInsert(local,value int) bool {
	var j int
	if local < 1 || local > l.Length+1 {
		return false
	}
	for j=l.Length; j >= local;j -- {
		if j ==  l.Length {
			l.Data = append(l.Data,l.Data[j-1])
		}else {
			l.Data[j] = l.Data[j-1]
		}
	}
	l.Data[local-1] = value
	l.Length++
	return true
}

/*
删除顺序表(指定位置)
	arr := []int{2,7,8,11,12,15,17}
	var l sqList.SqList   //{[] 0}
	L := l.CreateList(arr)
	fmt.Printf("re %#v\n",L)
	res := L.ListDeleter(3)
	fmt.Printf("res %v\n",res)
	fmt.Printf("L %v\n",L)
*/

func (l *SqList) ListDeleter(local int) bool {
	if local < 1 || local > l.Length+1 {
		return false
	}
	for j := local; j < l.Length-1  ; j ++ {
		l.Data[j] = l.Data[j+1]
	}
	l.Length--
	l.Data = l.Data[:l.Length-1]
	return true
}

/**
删除指定元素1
	arr := []int{2,7,12,11,12,15,17}
	var l sqList.SqList   //{[] 0}
	L := l.CreateList(arr)
	fmt.Printf("re %#v\n",L)
	res := L.DElNode1(12)
	fmt.Printf("res %v\n",res)
	fmt.Printf("L %v\n",L)
*/
func (l *SqList) DElNode1(val int) *SqList {
	var k int
	for i := 0; i < l.Length; i++ {
		if l.Data[i] != val {
			l.Data[k] = l.Data[i]
			k++
		}
	}
	l.Length = k
	l.Data = l.Data[:k]
	return l
}

/**
删除指定元素2
	arr := []int{2,7,12,11,12,15,17}
	var l sqList.SqList   //{[] 0}
	L := l.CreateList(arr)
	fmt.Printf("re %#v\n",L)
	res := L.DElNode2(12)
	fmt.Printf("res %v\n",res)
	fmt.Printf("L %v\n",L)
*/
func (l *SqList) DElNode2(val int) *SqList {
	k,i :=  0,0
	for i < l.Length {
		if l.Data[i] == val {
			k++
		}else {
			l.Data[i-k] = l.Data[i]
		}
		i++
	}
	l.Length -= k
	l.Data = l.Data[:l.Length]
	return l
}

/*
以第一个元素为分界，小于他的放在左边（前），大于的放在右边（后）
思路： 左右查找替换
	arr := []int{14,7,12,11,1,15,17}
	var l sqList.SqList   //{[] 0}
	L := l.CreateList(arr)
	fmt.Printf("re %#v\n",L)
	res := L.Move1()
	fmt.Printf("res %v\n",res)
*/
func (l *SqList) Move1() *SqList {
	pivot,j,i := l.Data[0],l.Length-1,0
	for i < j {
		for i < j && l.Data[j] > pivot {
			j--
		}
		for i < j && l.Data[i] <= pivot{
			i++
		}
		if i < j {
			l.Data[i],l.Data[j] = l.Data[j],l.Data[i]
		}
	}
	l.Data[0],l.Data[j] = l.Data[j],l.Data[0]
	return l
}

/*
以第一个元素为分界，小于他的放在左边（前），大于的放在右边（后）
思路：
	arr := []int{14,7,12,11,1,15,17}
	var l sqList.SqList   //{[] 0}
	L := l.CreateList(arr)
	fmt.Printf("re %#v\n",L)
	res := L.Move2()
	fmt.Printf("res %v\n",res)
*/
func (l *SqList) Move2() *SqList {
	pivot,j,i := l.Data[0],l.Length-1,0
	for i < j {
		for j > i && l.Data[j] > pivot {
			j--
		}
		l.Data[i] = l.Data[j]
		i++
		for i < j && l.Data[i] <= pivot {
			i++
		}
		l.Data[j] = l.Data[i]
		j--
	}
	l.Data[i] = pivot
	return l
}