package main

type List struct {
	elem []interface{}
	len  int
	cap  int
}

// getLen 求长度
func getLen(list List) {
	println(len(list.elem))
}

// clearList 清空表
func clearList(list *List) {
	list.elem = nil
	list.len = len(list.elem)
	list.cap = cap(list.elem)
}

// full 判断表是否满了
func full(list List) {
	if len(list.elem) == cap(list.elem) {
		println("表满")
	} else {
		println("表未满")
	}
}

// ifNil 判断是否为空
func ifNil(list List) {
	if list.elem != nil {
		println("这不是空的")
	} else {
		println("空的")
	}
}

// listAdd 添加数值
func listAdd(listA *List, seat int, value interface{}) {
	if seat >= 1 && seat <= listA.len {
		listB := listA.elem[:seat-1]

		listC := make([]interface{}, seat-1, listA.cap+1)
		copy(listC, listB)

		listC = append(listC, value)
		listB = listA.elem[seat-1:]

		listC = append(listC, listB...)
		listA.elem = listC

	} else {
		println("请填写正确的位置")
	}
	listA.len = len(listA.elem)
	listA.cap = cap(listA.elem)
}

// listChange 改变数值
func listChange(listA *List, seat int, value interface{}) {
	if seat >= 1 && seat <= listA.len {
		listB := listA.elem[:seat-1]

		listC := make([]interface{}, seat-1, listA.cap)
		copy(listC, listB)

		listC = append(listC, value)
		listB = listA.elem[seat:]
		listC = append(listC, listB...)
		listA.elem = listC

	} else {
		println("请填写正确的位置")
	}
	listA.len = len(listA.elem)
	listA.cap = cap(listA.elem)
}

// listDelete 删除指定元素
func listDelete(listA *List, seat int) {
	if seat >= 1 && seat <= listA.len {
		switch seat {

		case 0:
			listB := listA.elem[1:]
			listC := make([]interface{}, len(listA.elem), cap(listA.elem))
			copy(listC, listB)
			listA.elem = listC

		case len(listA.elem):
			listB := listA.elem[:len(listA.elem)-1]
			listC := make([]interface{}, len(listA.elem), cap(listA.elem))
			copy(listC, listB)
			listA.elem = listC

		default:
			listB := listA.elem[:seat-1]
			listC := make([]interface{}, seat-1, listA.cap)
			copy(listC, listB)
			listB = listA.elem[seat:]
			listC = append(listC, listB...)
			listA.elem = listC
		}
	} else {
		println("请填写正确的位置")
	}
	listA.len = len(listA.elem)
	listA.cap = cap(listA.elem)
}

// getElem 得到指定元素
func getElem(list List, seat int) {
	println(list.elem[seat])
}

// listergodic 遍历数组
func listergodic(list List) {
	for i := 0; i < len(list.elem); i++ {
		println(list.elem[i])
	}
	println()
}

// addList 拼接表AB
func addList(listA *List, listB *List) {
	listA.elem = append(listA.elem, listB.elem...)
	listA.len = len(listA.elem)
	listA.cap = cap(listA.elem)
}

// ergodic 遍历切片
func ergodic(list []int) {
	for i := 0; i < len(list); i++ {
		println(list[i])
	}
	println()
}

// deleteTarget 删除表中有指定内容的象
func deleteTarget(list *List, value interface{}) {
	Len := 0
	for i, v := range list.elem {
		if v == value {
			Len++
			switch i {

			case 0:
				listB := list.elem[1:]
				listC := make([]interface{}, len(list.elem)-Len, cap(list.elem))
				copy(listC, listB)
				list.elem = listC

			case list.len - 1:
				listB := list.elem[:len(list.elem)-1]
				listC := make([]interface{}, list.len-Len, cap(list.elem))
				copy(listC, listB)
				list.elem = listC

			default:
				listB := list.elem[:i-1]
				listC := make([]interface{}, i-1, list.cap)
				copy(listC, listB)
				listB = list.elem[i:]
				listC = append(listC, listB...)
				list.elem = listC
			}
		}
	}
	list.len = len(list.elem)
	list.cap = cap(list.elem)
}
