package BTNode

/*
						A
				B	   			C
			D	  	K  		E     	F
		H	   G		  N   S

ABCDKEFHGNS
*/

import (
	"fmt"
)

/*
二叉树链
 */
type BTNode struct {
	Data rune
	LChild *BTNode
	RChild *BTNode
}

type BTNodes interface{
	CreateBTNode(s string) *BTNode
}

/**
生产二叉树
	s := []rune("A(B(D(,G)),C(E,F))")
	b := new(BTNode.BTNode)
	b.CreateBTNode(s).DispBTNode()
 */
func (b *BTNode) CreateBTNode(s []rune) *BTNode {
	var p *BTNode
	top, k,b := -1,0,nil
	st := make(map[int]*BTNode)
	for _,v := range s {

		switch v {
		case '(':
			top++
			st[top] = p
			k = 1
			break
		case ')':
			top--
			break
		case ',':
			k= 2
			break
		default:
			p = new(BTNode)
			p.Data = v
			p.LChild,p.RChild = nil,nil
			if b == nil {
				b = p
			}else {
				switch k {
				case 1 :
					st[top].LChild = p
					break
				case 2 :
					st[top].RChild = p
					break
				}
			}
		}
	}
	return b
}

/**
二叉树链 查找指定元素,找到返回其地址，否则返回nil
	s := []rune("A(B(D(,G)),C(E,F))")
	b := new(BTNode.BTNode)
	a := b.CreateBTNode(s).FindBTNode(rune('B'))
	fmt.Println(a)
 */
func (b *BTNode) FindBTNode(x rune) *BTNode {
	if b == nil {
		return b
	}else if b.Data == x {
		return b
	} else {
		p := b.RChild.FindBTNode(x)
		if p != nil {
			return p
		} else {
			return b.LChild.FindBTNode(x)
		}
	}
}

func (b *BTNode) LChildBTNode( ) *BTNode {
	return b.LChild
}

func (b *BTNode) RChildBTNode( ) *BTNode {
	return b.RChild
}

/**
树高度
	s := []rune("A(B(D(,G)),C(E,F))")
	b := new(BTNode.BTNode)
	a := b.CreateBTNode(s).HeightBTNode()
	fmt.Println(a)
*/
func (b *BTNode) HeightBTNode() int {
	if b == nil {
		return 0
	}else {
		l := b.LChild.HeightBTNode()
		r := b.RChild.HeightBTNode()
		if l > r {
			return l+1
		}
		return r+1
	}
}

/**
二叉树链展示
*/
func (b *BTNode) DispBTNode() {
	if b != nil {
		fmt.Printf(" %v ",string(b.Data))
		if b.LChild != nil || b.RChild != nil {
			fmt.Printf(" %v ","(")
			b.LChild.DispBTNode()
			if b.RChild != nil {
				fmt.Printf(" %v ",",")
			}
			b.RChild.DispBTNode()
			fmt.Printf(" %v ",")")
		}
	}
}

/**
二叉树遍历 先序遍历
	s := []rune("A(B(D(H,G),K),C(E(N,S),F))")
	b := new(BTNode.BTNode)
	b.CreateBTNode(s).PreBTNode()
*/
func (b *BTNode) PreBTNode() {
	if b != nil {
		fmt.Printf("%v ",string(b.Data))
		b.LChild.PreBTNode()
		b.RChild.PreBTNode()
	}
}

/*
二叉树遍历 中序遍历
	s := []rune("A(B(D(H,G),K),C(E(N,S),F))")
	b := new(BTNode.BTNode)
	b.CreateBTNode(s).InBTNode()
 */
func (b *BTNode) InBTNode() {
	if b != nil {
		b.LChild.InBTNode()
		fmt.Printf("%v ",string(b.Data))
		b.RChild.InBTNode()
	}
}

/*
二叉树遍历 后序遍历
	s := []rune("A(B(D(H,G),K),C(E(N,S),F))")
	b := new(BTNode.BTNode)
	b.CreateBTNode(s).PostBTNode()
*/
func (b *BTNode) PostBTNode() {
	if b != nil {
		b.LChild.PostBTNode()
		b.RChild.PostBTNode()
		fmt.Printf("%v ",string(b.Data))
	}
}

/*
二叉树遍历 输出所有叶子节点 左边开始
	s := []rune("A(B(D(H,G),K),C(E(N,S),F))")
	b := new(BTNode.BTNode)
	b.CreateBTNode(s).DispL()
*/

func (b *BTNode) DispL() {
	if b != nil {
		if b.LChild == nil && b.RChild == nil {
			fmt.Printf("%s " ,string(b.Data))
		}else {
			b.LChild.DispL()
			b.RChild.DispL()
		}
	}
}

/*
二叉树遍历 输出所有叶子节点 右边开始
	s := []rune("A(B(D(H,G),K),C(E(N,S),F))")
	b := new(BTNode.BTNode)
	b.CreateBTNode(s).DispR()
*/

func (b *BTNode) DispR() {
	if b != nil {
		if b.LChild == nil && b.RChild == nil {
			fmt.Printf("%s " ,string(b.Data))
		}else {
			b.LChild.DispR()
			b.RChild.DispR()
		}
	}
}

/*
求某一个值在二叉树的深度
*/
func (b *BTNode) Level(x rune,h int) (*BTNode,int)  {
	if b == nil {
		return b,0
	}else if b.Data == x {
		return b,h
	}else {
		l,n := b.LChild.Level(x,h+1)
		if n != 0 {
			return l,n
		}
		return b.RChild.Level(x,h+1)
	}
}

/*
判断两树相似，相似不一定相等
*/
func (b *BTNode) Common(c *BTNode) bool  {
	if b == nil && c == nil {
		return true
	}else if b == nil && c == nil {
		return false
	}else {
		n := b.LChild.Common(c.LChild)
		m := b.RChild.Common(c.RChild)
		return n && m
	}
}

/*
查找链树 指定元素的所有父节点
*/
func (b *BTNode) Ancestor(x rune) bool  {
	if b == nil {
		return false
	}else if (b.LChild != nil && b.LChild.Data == x) || (b.RChild != nil && b.RChild.Data == x)   {
		fmt.Printf("%v ",string(b.Data))
		return true
	} else if b.LChild.Ancestor(x) || b.RChild.Ancestor(x) {
		fmt.Printf("%v ",string(b.Data))
		return true
	}else {
		return false
	}
}

/****************************************二叉树非递归遍历start******************************************/

/*
先序遍历
	s := []rune("A(B(D(H,G),K),C(E(N,S),F))")
	b := new(BTNode.BTNode)
	b.CreateBTNode(s).PreOrder()
	b.CreateBTNode(s).PreBTNode()
*/
func (b *BTNode) PreOrder() {
	st := make(map[int]*BTNode)
	top := -1
	if b != nil {
		top++
		st[top] = b
		for top > -1  {
			p := st[top]
			top--
			fmt.Printf("%v ",string(p.Data))
			if p.RChild != nil {
				top++
				st[top] = p.RChild
			}
			if p.LChild != nil {
				top++
				st[top] = p.LChild
			}
		}
		fmt.Printf("%v","\n")
	}
}

/*
中序遍历
	s := []rune("A(B(D(H,G),K),C(E(N,S),F))")
	b := new(BTNode.BTNode)
	b.CreateBTNode(s).InOrder()
*/
func (b *BTNode) InOrder() {
	st := make(map[int]*BTNode)
	top := -1
	if b != nil {
		p := b
		for top > -1 || p != nil  {
			for p != nil {
				top++
				st[top] = p
				p = p.LChild
			}
			if top > -1 {
				p = st[top]
				top--
				fmt.Printf(" %v ",string(p.Data))
				p = p.RChild
			}
		}
		fmt.Printf("%v","\n")
	}
}

func (b *BTNode) PostOrder() {

}

/****************************************二叉树非递归遍历end********************************************/
//A
//    B     C
//  D   K    EF
// HG   NS
 func Trans(s []rune,i int) *BTNode {
	if i > len(s)-1 || s[i] == '#' {
		return nil
	}else {
		b := new(BTNode)
		b.Data = s[i]
		fmt.Println(2*i,"==",i,"=",2*i+1,string(s[i]))
		b.LChild = Trans(s,2*i)
		b.RChild = Trans(s,2*i+1)
		return b
	}
}