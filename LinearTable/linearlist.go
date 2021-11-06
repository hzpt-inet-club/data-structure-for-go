package LinearTable

import "fmt"

type List struct {
	elem []interface{}
	len  int
	cap  int
}

func (listA *List) Elem() []interface{} {
	return listA.elem
}

func (listA *List) SetElem(elem []interface{}) {
	listA.elem = elem
	listA.cap = cap(elem)
	listA.len = len(elem)
}

// GetLen 求长度
func (listA List) GetLen() int {
	return len(listA.elem)
}

// ClearList 清空表
func (listA *List) ClearList() {
	listA.elem = nil
	listA.len = len(listA.elem)
	listA.cap = cap(listA.elem)
}

// IfFull 判断表是否满了
func (listA List) IfFull() {
	if len(listA.elem) == cap(listA.elem) {
		println("表满")
	} else {
		println("表未满")
	}
}

// IfNil 判断是否为空
func (listA List) IfNil() {
	if listA.elem != nil {
		println("这不是空的")
	} else {
		println("空的")
	}
}

// ListAdd 添加数值
func (listA *List) ListAdd(seat int, value interface{}) {
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

// ListChange 改变指定位置的元素
func (listA *List) ListChange(seat int, value interface{}) {
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

// ListDelete 删除指定位置的元素
func (listA *List) ListDelete(seat int) {
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

// GetElem 得到指定位置的元素
func (listA List) GetElem(seat int) {
	println(listA.elem[seat])
}

// Listergodic 遍历数组
func (listA List) Listergodic() {
	for i := 0; i < len(listA.elem); i++ {
		fmt.Println(listA.elem[i])
	}
	println()
}

// Splicing 拼接表AB
func Splicing(listA List, listB List) (listC List) {
	listC.elem = append(listA.elem, listB.elem...)
	listC.len = len(listA.elem)
	listC.cap = cap(listA.elem)
	return listC
}

// DeleteTarget 删除表中有指定内容的项
func (listA *List) DeleteTarget(value interface{}) {
	Len := 0
	var listC []interface{}
	for i, v := range listA.elem {
		if v == value {
			switch i {

			case 0:
				listA.elem = listA.elem[1:]

			case listA.len - 1:
				listA.elem = listA.elem[:i-Len]

			default:
				listB := listA.elem[:i-Len]
				listC = make([]interface{}, i-Len, listA.cap)
				copy(listC, listB)
				listB = listA.elem[i+1-Len:]
				listC = append(listC, listB...)
				listA.elem = listC
			}
			Len++
		}
	}
	listA.len = len(listA.elem)
	listA.cap = cap(listA.elem)
}
