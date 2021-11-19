package SingleLinkedTable

import "fmt"

// SingleLinkList 单项链表
type SingleLinkList struct {
	head head
}

type head struct {
	Next *node
}

type node struct {
	Data interface{}
	Next *node
}

// GetLen 求链表长度
func (LinkListA SingleLinkList) GetLen() int {
	index := 1
	var top node
	if LinkListA.head.Next == nil {
		return 0
	}
	top = *LinkListA.head.Next
	for top.Next != nil {
		index++
		top = *top.Next
	}
	return index
}

// AddFirst 添加数据到链表首位
func (LinkListA *SingleLinkList) AddFirst(a interface{}) {
	var getNode = node{Data: a}
	getNode.Next = LinkListA.head.Next
	LinkListA.head.Next = &getNode
}

// ListErgodic 遍历链表
func (LinkListA *SingleLinkList) ListErgodic() {
	top := LinkListA.head.Next
	for top != nil {
		fmt.Println(top.Data)
		top = top.Next
	}
}

// AddRead 添加数据到链表末位
func (LinkListA *SingleLinkList) AddRead(a interface{}) {
	if LinkListA.head.Next == nil {
		getNode := node{Data: a}
		LinkListA.head.Next = &getNode
		return
	}
	top := LinkListA.head.Next
	for top.Next != nil {
		top = top.Next
	}
	getNode := node{Data: a}
	top.Next = &getNode
}

// FindNode 查找指定位置的节点
func (LinkListA SingleLinkList) FindNode(seat int) node {
	var top node
	if seat < 1 || seat > LinkListA.GetLen() {
		fmt.Println("请输入正确的值")
		return top
	}
	top = *LinkListA.head.Next
	for count := 1; count != seat; count++ {
		top = *top.Next
	}
	return top
}

// AddI 在指定位置插入节点
func (LinkListA *SingleLinkList) AddI(seat int, a interface{}) {
	getNodeA := node{Data: a}
	getNodeB := LinkListA.FindNode(seat)
	getNodeA.Next = &getNodeB
	getNodeC := LinkListA.FindNode(seat - 1)
	getNodeC.Next = &getNodeA
	LinkListA.head.Next = &getNodeC
}

// DeleteI 删除指定位置的节点
func (LinkListA *SingleLinkList) DeleteI(seat int) {
	var LinkListB SingleLinkList
	if LinkListA.GetLen() == 1 {
		LinkListA.head.Next = nil
		return
	} else {
		if seat >= 1 && seat <= LinkListA.GetLen() {
			if seat == 1 {
				*LinkListA.head.Next = *LinkListA.head.Next.Next
				return
			} else {
				getNode := LinkListA.FindNode(seat - 1)
				getNode.Next = getNode.Next.Next
				top := LinkListA.head.Next
				for top.Next != nil {
					if top.Next != getNode.Next {
						LinkListB.AddRead(top.Data)
						top = top.Next
					} else {
						LinkListB.ConnectNode(getNode)
						*LinkListA = LinkListB
						return
					}
				}
				return
			}
		}
	}
	fmt.Println("请输入正确的值")
}

// ConnectNode 把节点拼接到表后
func (LinkListA *SingleLinkList) ConnectNode(node node) {
	top := LinkListA.head.Next
	for top.Next != nil {
		top = top.Next
	}
	top.Next = node.Next
}

// ConnectLinkList 把表B拼接到表A后
func (LinkListA *SingleLinkList) ConnectLinkList(list SingleLinkList) {
	top := LinkListA.head.Next
	for top.Next != nil {
		top = top.Next
	}
	top.Next = list.head.Next
}

// FindValue 查找Node中Data等于Value的项
func (LinkListA SingleLinkList) FindValue(a interface{}) ([]int, []node) {
	top := LinkListA.head.Next
	intSlice := make([]int, LinkListA.GetLen())
	nodeSlice := make([]node, LinkListA.GetLen())
	j := 0
	for i := 1; i <= LinkListA.GetLen(); i++ {
		if top.Data == a {
			intSlice[j] = i
			nodeSlice[j] = *top
			j++
		}
		top = top.Next
	}
	return intSlice, nodeSlice
}
