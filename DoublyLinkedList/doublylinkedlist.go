package DoublyLinkedList

import "fmt"

// DoublyLinkList 循环双向链表
type DoublyLinkList struct {
	head Node
}

type Node struct {
	Data interface{}
	pre  *Node
	next *Node
}

// GetLen 求链表长度
func (LinkListA DoublyLinkList) GetLen() int {
	index := 0
	var top Node
	if LinkListA.head.next == nil {
		return 0
	}
	top = *LinkListA.head.next
	for top.Data != nil {
		index++
		top = *top.next
	}
	return index
}

// AddFirst 添加数据到链表首位
func (LinkListA *DoublyLinkList) AddFirst(a interface{}) {
	var getNode = Node{Data: a}

	getNode.pre = &LinkListA.head

	if LinkListA.GetLen() == 0 {
		getNode.next = &LinkListA.head
		LinkListA.head.next = &getNode
		LinkListA.head.pre = &getNode
	} else {
		getNode.next = LinkListA.head.next
		LinkListA.head.next = &getNode
		bottom := LinkListA.head.pre
		for bottom.pre.Data != nil {
			bottom = bottom.pre
		}
		bottom.pre = &getNode
	}
}

// ListErgodic 遍历链表
func (LinkListA *DoublyLinkList) ListErgodic() {
	top := LinkListA.head.next
	for top.Data != nil {
		fmt.Println(top.Data)
		top = top.next
	}
}

// AddRead 添加数据到链表末位
func (LinkListA *DoublyLinkList) AddRead(a interface{}) {
	var getNode = Node{Data: a}

	getNode.next = &LinkListA.head

	if LinkListA.GetLen() == 0 {
		getNode.pre = &LinkListA.head
		LinkListA.head.pre = &getNode
		LinkListA.head.next = &getNode
	} else {
		getNode.pre = LinkListA.head.pre
		LinkListA.head.pre = &getNode
		bottom := LinkListA.head.next
		for bottom.next.Data != nil {
			bottom = bottom.next
		}
		bottom.next = &getNode
	}
}

// FindNode 查找指定位置的节点
func (LinkListA DoublyLinkList) FindNode(seat int) Node {
	var top Node
	if seat < 1 || seat > LinkListA.GetLen() {
		fmt.Println("请输入正确的值")
		return top
	}
	top = *LinkListA.head.next
	for count := 1; count != seat; count++ {
		top = *top.next
	}
	return top
}

// AddI 在指定位置插入节点
func (LinkListA *DoublyLinkList) AddI(seat int, a interface{}) {
	getNodeA := Node{Data: a}

	var indexA *Node
	if seat < 1 || seat > LinkListA.GetLen()+1 {
		fmt.Println("请输入正确的值")
	}
	indexA = &LinkListA.head
	for count := 1; count != seat; count++ {
		indexA = indexA.next
	}

	indexB := indexA.next
	getNodeA.pre = indexA
	getNodeA.next = indexB
	indexA.next = &getNodeA
	indexB.pre = &getNodeA
}

// DeleteI 删除第i个位置的节点
func (LinkListA *DoublyLinkList) DeleteI(seat int) {

	var indexA *Node
	if seat < 1 || seat > LinkListA.GetLen() {
		fmt.Println("请输入正确的值")
		return
	}
	indexA = &LinkListA.head
	for count := 1; count != seat; count++ {
		indexA = indexA.next
	}
	indexB := *indexA.next.next
	indexA.next = &indexB
	indexB.pre = indexA
}

// ConnectLinkList 把双向环形链表B连接在双向环形链表A后
func (LinkListA *DoublyLinkList) ConnectLinkList(LinkListB DoublyLinkList) {
	top := LinkListB.head.next
	for top.Data != nil {
		LinkListA.AddRead(top.Data)
		top = top.next
	}
}

// FindValue 查找Node中Data等于Value的项
func (LinkListA DoublyLinkList) FindValue(a interface{}) ([]int, []Node) {
	top := LinkListA.head.next
	bottom := LinkListA.head.pre
	intSlice := make([]int, 0, LinkListA.GetLen())
	nodeSlice := make([]Node, 0, LinkListA.GetLen())
	j := 0
	for i := 1; i <= LinkListA.GetLen(); i++ {
		if top.Data == a {
			intSlice = append(intSlice, i)
			nodeSlice = append(nodeSlice, *top)
			j++
		}
		if bottom.Data == a {
			intSlice = append(intSlice, LinkListA.GetLen()-i+1)
			nodeSlice = append(nodeSlice, *bottom)
			j++
		}
		top = top.next
		if top == bottom {
			break
		}
		bottom = bottom.pre
	}
	for i := 0; i < len(intSlice)-1; i++ {
		for j = i; j < len(intSlice); j++ {
			if intSlice[i] > intSlice[j] {
				tempInt := intSlice[i]
				intSlice[i] = intSlice[j]
				intSlice[j] = tempInt

				tempNode := nodeSlice[i]
				nodeSlice[i] = nodeSlice[j]
				nodeSlice[j] = tempNode
			}
		}
	}
	return intSlice, nodeSlice
}
