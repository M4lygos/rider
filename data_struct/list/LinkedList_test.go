package list

import (
	"Rider/server/core"
	"fmt"
	"testing"
)

func TestList_InsertHead(t *testing.T) {
	list := NewList()
	list.InsertHead(core.CreateObject(1, 1))
	list.InsertTail(core.CreateObject(1, 2))
	list.InsertTail(core.CreateObject(1, 3))
	list.InsertHead(core.CreateObject(1, 4))
	fmt.Println(list.Length())
	for list.head.next != list.tail {
		val, _ := list.DeleteTail()
		fmt.Println(*val.val)
		fmt.Println(list.nodeCount)
	}
}
