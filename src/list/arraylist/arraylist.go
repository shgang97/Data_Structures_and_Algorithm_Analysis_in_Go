package arraylist

import (
	"errors"
	"fmt"
)

/*
@author: shg
@since: 2022/8/15
@desc: //TODO
*/

const (
	defaultCapacity int = 10
)

type ArrayList struct {
	elementData []interface{}
	size        int
}

func New(initCapacity ...int) *ArrayList {
	list := new(ArrayList)
	var capacity int
	if len(initCapacity) == 1 {
		capacity = initCapacity[0]
	} else {
		capacity = defaultCapacity
	}
	list.elementData = make([]interface{}, capacity)
	list.size = 0
	return list
}

func (list *ArrayList) Size() int {
	return list.size
}

func (list *ArrayList) IsEmpty() bool {
	return list.size == 0
}

func (list *ArrayList) Contains(value interface{}) bool {
	return list.IndexOf(value) != -1
}

func (list *ArrayList) Get(index int) (interface{}, error) {
	if index < 0 || list.size <= index {
		return nil, errors.New(fmt.Sprintf("Index out of range: %d", index))
	}
	return list.elementData[index], nil
}

func (list *ArrayList) Set(index int, value interface{}) error {
	if index < 0 || list.size < index {
		return errors.New(fmt.Sprintf("Index out of range: %d", index))
	}
	list.elementData[index] = value
	return nil
}

func (list *ArrayList) Insert(index int, value interface{}) error {
	if index < 0 || list.size < index {
		return errors.New(fmt.Sprintf("Index out of range: %d", index))
	}

	s := list.size
	elementData := list.elementData
	if s == len(elementData) {
		elementData = list.grow()
	}

	for i := s; i > index; i-- {
		elementData[i] = elementData[i-1]
	}
	elementData[index] = value
	list.elementData = elementData
	list.size++
	return nil
}

func (list *ArrayList) Append(value interface{}) error {
	return list.Insert(list.size, value)
}

func (list *ArrayList) Remove(index int) error {
	if index < 0 || list.size < index {
		return errors.New(fmt.Sprintf("Index out of range: %d", index))
	}
	for i := index + 1; i < list.size; i++ {
		list.elementData[i-1] = list.elementData[i]
	}
	list.elementData[list.size-1] = nil
	list.size--
	return nil
}

func (list *ArrayList) IndexOf(value interface{}) int {
	return list.indexOfRange(value, 0, list.size)
}

func (list *ArrayList) indexOfRange(value interface{}, start, end int) int {
	if value == nil {
		for i := start; i < end; i++ {
			if list.elementData[i] == nil {
				return i
			}
		}
	} else {
		for i := start; i < end; i++ {
			if list.elementData[i] == value {
				return i
			}
		}
	}
	return -1
}

func (list *ArrayList) Clear() {
	for i := 0; i < list.size; i++ {
		list.elementData[i] = nil
	}
	list.size = 0
}

func (list *ArrayList) String() string {
	// ArrayList{size=3, elementData=[1, 2, 3]}
	msg := fmt.Sprintf("ArrayList{size=%d, elementData=[", list.size)
	for i := 0; i < list.size; i++ {
		msg += fmt.Sprintf("%v", list.elementData[i])
		if i < list.size-1 {
			msg += fmt.Sprint(", ")
		}
	}
	msg += "]}"
	return msg
}

func (list *ArrayList) grow() []interface{} {
	oldCapacity := len(list.elementData)
	newCapacity := oldCapacity + (oldCapacity >> 1)

	newElementData := make([]interface{}, newCapacity)
	for i := 0; i < list.size; i++ {
		newElementData[i] = list.elementData[i]
	}
	return newElementData
}
