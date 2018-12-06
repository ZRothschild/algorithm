package list

import (
	"fmt"
)

/*
单链表
*/
type LNode struct {
	Data int
	Next *LNode
}

type LinkNode interface {
	CreateListF(arr []int) *LNode
	CreateListR(arr []int) *LNode
}

/*
生成链表（头插入法）
	arr := []int{14,7,4,8,10,1,6}
	var l list.LNode   //{[] 0}
	L := l.CreateListF(arr)
	fmt.Printf("re %v  %v \n",L.Next.Data,L.Next.Next.Data)
*/
func (n *LNode) CreateListF(arr []int) *LNode {
	n.Next = nil
	for _,v := range arr {
		s := new (LNode)
		s.Data = v
		s.Next = n.Next
		n.Next = s
	}
	return n
}

/*
生成链表（尾插入法）
	arr := []int{14,7,4,8,10,1,6}
	var l list.LNode   //{[] 0}
	L := l.CreateListR(arr)
	fmt.Printf("re %#v\n",L.Next.Next.Next.Data)
*/
func (n *LNode) CreateListR(arr []int) *LNode {
	r := n
	for _,v := range arr {
		s := new (LNode)
		s.Data = v
		r.Next = s
		r = s
	}
	r.Next = nil
	return n
}

/*
 插入节点
	arr := []int{14,7,4,8,10,1,6}
	var l list.LNode   //{[] 0}
	L := l.CreateListR(arr).ListInsert(3,19)
	L.DispList()
*/
func (n *LNode) ListInsert(local, val int) *LNode {
	j := 0
	p := n
	for j < local -1 && p != nil {
		p = p.Next
		j++
	}
	if p == nil {
		return n
	}else {
		s := new (LNode)
		s.Data = val
		s.Next = p.Next
		p.Next = s
		return n
	}
}

/*
删除节点
	arr := []int{14,7,4,8,10,1,6}
	var l list.LNode   //{[] 0}
	L := l.CreateListR(arr).DeleterInsert(3)
	L.DispList()
*/
func (n *LNode) DeleterInsert(local int) *LNode {
	j := 0
	p := n
	for j < local -1 && p != nil {
		p = p.Next
		j++
	}
	if p == nil {
		return n
	}else {
		q := p.Next
		if q == nil {
			return n
		}
		p.Next = q.Next
		return n
	}
}

/*
输出线性表
*/
func (n *LNode) DispList() {
	p := n.Next
	for p != nil {
		fmt.Printf("%d ",p.Data)
		p = p.Next
	}
	fmt.Print("\n")
}

/*
删除最大节点
	arr := []int{2,7,4,8,10,1,6}
	var l list.LNode   //{[] 0}
	l.CreateListR(arr).DelMaxNode().DispList()
 */
func (n *LNode) DelMaxNode() *LNode {
	p,pre := n.Next,n
	maxP,maxPre := p,pre
	for p != nil {
		if maxP.Data < p.Data {
			maxP = p
			maxPre = pre
		}
		pre = p
		p = p.Next
	}
	maxPre.Next = maxP.Next
	return n
}

/*
排序   zhi
	arr := []int{2,7,4,8,10,4,6}
	var l list.LNode   //{[] 0}
	l.CreateListR(arr).Sort().DispList()
*/
func (n *LNode) Sort() *LNode  {
	p := n.Next.Next
	n.Next.Next = nil
	for p != nil {
		q := p.Next
		pre := n
		//n.DispList()
		//fmt.Println(pre.Next.Data,p.Data)
		for pre.Next != nil && pre.Next.Data < p.Data {
			pre = pre.Next
		}
		p.Next = pre.Next
		//fmt.Println(pre.Next)
		pre.Next = p
		p = q
		q.DispList()
	}
	return n
}