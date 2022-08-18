package linkedlist

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {
	list := New()
	list.Insert(0, 1)
	fmt.Println(list)
	list.Insert(0, 2)
	fmt.Println(list)
	list.Insert(2, 3)
	fmt.Println(list)
	list.Insert(1, 4)
	fmt.Println(list)

	fmt.Println(list.Get(0))
	fmt.Println(list.Get(1))
	fmt.Println(list.Get(2))
	fmt.Println(list.Get(3))
	fmt.Println(list.Get(4))
	fmt.Println(list.Get(5))

	err := list.Set(4, 10)
	if err != nil {
		fmt.Println(err)
	}
	err = list.Set(3, 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)

	list.Insert(4, 11)
	fmt.Println(list)
	list.Append(12)
	fmt.Println(list)

	//list.Clear()

	fmt.Println(list)
	list.Remove(2)
	fmt.Println(list)
	list.Remove(0)
	fmt.Println(list)
	list.Remove(list.size - 1)
	fmt.Println(list)

	fmt.Println(list.IndexOf(4))
	fmt.Println(list.IndexOf(10))
	fmt.Println(list.IndexOf(11))

	fmt.Println(list.Contains(4))
	fmt.Println(list.Contains(10))
	fmt.Println(list.Contains(11))
	fmt.Println(list.Contains(12))
}
