package arraylist

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {

	list := New()
	list.Insert(0, 0)
	list.Insert(1, 1)
	list.Insert(2, 2)
	list.Insert(3, 3)

	fmt.Println(list)

	list.Remove(1)
	fmt.Println(list)

	list.Insert(0, nil)
	fmt.Println(list)

	list.Append("1")
	fmt.Println(list)

	println(list.IndexOf(3))

}
