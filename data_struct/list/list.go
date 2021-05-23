package list

import (
	"Rider/data_struct"
	"errors"
)

const (
	LIST_EMPTY = "the length of list is zero"
)

type List struct {
	head 		*Node
	tail 		*Node
	nodeCount   int
}

type Node struct {
	Val *data_struct.RiderObject
	next *Node
	pre  *Node
}

func NewList() *List {
	head := new(Node)
	tail := new(Node)
	head.pre = nil
	head.next = tail
	tail.pre = head
	tail.next = nil
	return &List{
		head: head,
		tail: tail,
		nodeCount: 0,
	}
}

func (list *List) InsertHead(object *data_struct.RiderObject)  {
	node := &Node{
		Val: object,
	}
	head := list.head.next
	node.next = head
	node.pre = list.head
	head.pre = node
	list.head.next = node
	list.nodeCount ++
	return
}

func (list *List) InsertTail(object *data_struct.RiderObject)  {
	node := &Node{
		Val: object,
	}
	tail := list.tail.pre
	node.next = list.tail
	node.pre = tail
	tail.next = node
	list.tail.pre = node
	list.nodeCount ++
	return
}

func (list *List) DeleteHead() (*Node, error) {
	if list.nodeCount == 0 {
		return nil, errors.New(LIST_EMPTY)
	}
	next := list.head.next
	list.head.next = next.next
	next.next.pre = list.head
	list.nodeCount --
	return next, nil
}

func (list *List) DeleteTail() (*Node, error) {
	if list.nodeCount == 0 {
		return nil, errors.New(LIST_EMPTY)
	}
	pre := list.tail.pre
	list.tail.pre = pre.pre
	pre.pre.next = list.tail
	list.nodeCount --
	return pre, nil
}

func (list *List) Length() int {
	return list.nodeCount
}