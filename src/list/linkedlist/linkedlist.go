package linkedlist

import (
	"errors"
	"fmt"
)

/*
@author: shg
@since: 2022/8/16
@desc: //TODO
*/

type LinkedList struct {
	size  int
	first *node
	last  *node
}

type node struct {
	value interface{}
	pre   *node
	next  *node
}

func newNode(pre *node, value interface{}, next *node) *node {
	node := new(node)
	node.pre = pre
	node.value = value
	node.next = next
	return node
}

func New() *LinkedList {
	list := new(LinkedList)
	return list
}
func (list *LinkedList) Size() int {
	return list.size
}

func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *LinkedList) Contains(value interface{}) bool {
	return list.IndexOf(value) != -1
}

func (list *LinkedList) Get(index int) (interface{}, error) {
	if index < 0 || list.size <= index {
		return nil, errors.New(fmt.Sprintf("Index out of range: %d", index))
	}
	return list.node(index).value, nil
}

func (list *LinkedList) Set(index int, value interface{}) error {
	if index < 0 || list.size <= index {
		return errors.New(fmt.Sprintf("Index out of range: %d", index))
	}
	x := list.node(index)
	x.value = value
	return nil
}

func (list *LinkedList) Insert(index int, value interface{}) error {
	if index < 0 || list.size < index {
		return errors.New(fmt.Sprintf("Index out of range: %d", index))
	}

	if index == list.size {
		list.linkLast(value)
	} else {
		list.linkBefore(value, list.node(index))
	}

	return nil
}

func (list *LinkedList) Append(value interface{}) error {
	return list.Insert(list.size, value)
}

func (list *LinkedList) Remove(index int) error {
	if index < 0 || list.size <= index {
		return errors.New(fmt.Sprintf("Index out of range: %d", index))
	}

	node := list.node(index)
	if node.pre != nil {
		node.pre.next = node.next
	} else {
		list.first = node.next
	}

	if node.next != nil {
		node.next.pre = node.pre
	} else {
		list.last = node.pre
	}
	return nil
}

func (list *LinkedList) IndexOf(value interface{}) int {
	i := 0
	for x := list.first; x != nil; x = x.next {
		if x.value == value {
			return i
		}
		i++
	}
	return -1
}

func (list *LinkedList) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

func (list *LinkedList) String() string {
	// LinkedList{size=3, [1, 2, 3]}
	msg := fmt.Sprintf("LinkedList{size=%d, [", list.size)
	for x := list.first; x != nil; x = x.next {
		msg += fmt.Sprintf("%v", x.value)
		if x.next != nil {
			msg += ", "
		}
	}
	msg += "]}"
	return msg
}

func (list *LinkedList) linkLast(value interface{}) {
	l := list.last
	newNode := newNode(l, value, nil)
	list.last = newNode
	if l == nil {
		list.first = newNode
	} else {
		l.next = newNode
	}
	list.size++
}

func (list *LinkedList) node(index int) *node {
	if index < (list.size >> 1) {
		x := list.first
		for i := 0; i < index; i++ {
			x = x.next
		}
		return x
	} else {
		x := list.last
		for i := list.size - 1; i > index; i-- {
			x = x.pre
		}
		return x
	}
}

func (list *LinkedList) linkBefore(value interface{}, succ *node) {
	pred := succ.pre
	newNode := newNode(pred, value, succ)
	succ.pre = newNode
	if pred == nil {
		list.first = newNode
	} else {
		pred.next = newNode
	}
	list.size++
}
