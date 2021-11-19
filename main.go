package main

import (
	"_data-structure-for-go/DoublyLinkedList"
	"fmt"
)

func main() {
	var a DoublyLinkedList.DoublyLinkList
	var d DoublyLinkedList.DoublyLinkList
	a.AddRead("4")
	d.AddRead("1")
	a.AddRead("2")
	d.AddRead("2")
	d.ConnectLinkList(a)
	d.ListErgodic()

	var c, b = d.FindValue("2")
	fmt.Println(c)
	fmt.Println(b)
}
