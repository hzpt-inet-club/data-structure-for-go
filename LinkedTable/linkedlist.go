package LinkedTable

import "fmt"

type LinkList struct {
	Head Head
}

type Head struct {
	Next *Node
}

type Node struct {
	Data interface{}
	Next *Node
}

// GetLen 求链表长度
func (LinkListA LinkList) GetLen() int {
	index := 1
	var head Node
	if LinkListA.Head.Next == nil {
		return 0
	}
	head = *LinkListA.Head.Next
	for head.Next != nil {
		index++
		head = *head.Next
	}
	return index
}

// AddFirst 添加数据到链表首位
func (LinkListA *LinkList) AddFirst(a interface{}) {
	var node = Node{Data: a}
	node.Next = LinkListA.Head.Next
	LinkListA.Head.Next = &node
}

// ListErgodic 遍历链表
func (LinkListA *LinkList) ListErgodic() {
	head := LinkListA.Head.Next
	for head != nil {
		fmt.Println(head.Data)
		head = head.Next
	}
}

// AddRead 添加数据到链表末位
func (LinkListA *LinkList) AddRead(a interface{}) {
	if LinkListA.Head.Next == nil {
		node := Node{Data: a}
		LinkListA.Head.Next = &node
		return
	}
	head := LinkListA.Head.Next
	for head.Next != nil {
		head = head.Next
	}
	node := Node{Data: a}
	head.Next = &node
}

// FindNode 查找第i个位置的节点
func (LinkListA LinkList) FindNode(deli int) Node {
	var head Node
	if deli < 1 || deli > LinkListA.GetLen() {
		fmt.Println("请输入正确的值")
		return head
	}
	head = *LinkListA.Head.Next
	for count := 1; count != deli; count++ {
		if head.Next != nil {
			head = *head.Next
		}
	}
	return head
}

// AddI 在第i个位置插入节点
func (LinkListA *LinkList) AddI(deli int, a interface{}) {
	nodA := Node{Data: a}
	nodeB := LinkListA.FindNode(deli)
	nodA.Next = &nodeB
	nodeC := LinkListA.FindNode(deli - 1)
	nodeC.Next = &nodA
	LinkListA.Head.Next = &nodeC
}

// DeleteI 删除第i个位置的节点
func (LinkListA *LinkList) DeleteI(index int) {
	var LinkListB LinkList
	if LinkListA.GetLen() == 1 {
		LinkListA.Head.Next = nil
		return
	} else {
		if index >= 1 && index <= LinkListA.GetLen() {
			if index == 1 {
				*LinkListA.Head.Next = *LinkListA.Head.Next.Next
				return
			} else {
				nodeA := LinkListA.FindNode(index - 1)
				nodeA.Next = nodeA.Next.Next
				head := LinkListA.Head.Next
				for head.Next != nil {
					if head.Next != nodeA.Next {
						LinkListB.AddRead(head.Data)
						head = head.Next
					} else {
						LinkListB.ConnectNode(nodeA)
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
func (LinkListA *LinkList) ConnectNode(node Node) {
	head := LinkListA.Head.Next
	for head.Next != nil {
		head = head.Next
	}
	head.Next = node.Next
}

// ConnectLinkList 把表B拼接到表A后
func (LinkListA *LinkList) ConnectLinkList(list LinkList) {
	head := LinkListA.Head.Next
	for head.Next != nil {
		head = head.Next
	}
	head.Next = list.Head.Next
}

// FindValue 查找Node中Data等于Value的项
func (LinkListA LinkList) FindValue(a interface{}) ([]int, []Node) {
	head := LinkListA.Head.Next
	intSlice := make([]int, LinkListA.GetLen())
	nodeSlice := make([]Node, LinkListA.GetLen())
	j := 0
	for i := 1; i <= LinkListA.GetLen(); i++ {
		if head.Data == a {
			intSlice[j] = i
			nodeSlice[j] = *head
			j++
		}
		head = head.Next
	}
	return intSlice, nodeSlice
}
