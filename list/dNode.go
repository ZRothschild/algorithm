//双链表
package list

/*
双链表
 */
type DNode struct {
	Data int
	Prior *DNode
	Next *DNode
}

/*
双链表生成 前插入法
 */
func (d *DNode) CreateListF(arr []int) *DNode  {
	d.Prior,d.Next = nil,nil
	for _,v := range arr {
		s := new(DNode)
		s.Data = v
		s.Next = d.Next
		if d.Next != nil {
			d.Next.Prior = s
		}
		d.Next = s
		s.Prior = d
	}
	return d
}

/*
双链表生成 尾插入法
 */
func (d *DNode) CreateListR(arr []int) *DNode  {
	r := d
	for _,v := range arr  {
		s := new(DNode)
		s.Data = v
		r.Next = s
		s.Prior = r
		r= s
	}
	r.Next = nil
	return d
}

/*
双链表 指定位置插入元素
	arr := []int{2,7,4,8,10,4,6}
	var l list.DNode   //{[] 0}
	L := l.CreateListR(arr).ListerInsert(1,10)
	fmt.Println(L.Next.Data)
 */
func (d *DNode) ListerInsert(local,val int) *DNode  {
	p := d
	for local > 1 && p != nil{
		p = p.Next
		local--
	}
	if p == nil {
		return d
	}else {
		s := new(DNode)
		s.Data = val
		s.Next = p.Next
		if p.Next != nil {
			p.Next.Prior = s
		}
		s.Prior = p
		p.Next = s
		return d
	}
}