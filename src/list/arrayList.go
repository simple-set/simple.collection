package list

import (
	"errors"
	"fmt"
	"reflect"
)

var defaultCapacity = 16

type nodeSlice[T any] []T

type ArrayList[T any] struct {
	elementData nodeSlice[T]
}

func (a *ArrayList[T]) Add(data T) {
	if a.elementData == nil {
		a.elementData = newNodeSlice[T](defaultCapacity)
	}
	a.elementData = append(a.elementData, data)
}

func (a *ArrayList[T]) AddLast(data T) {
	a.Add(data)
}

func (a *ArrayList[T]) AddFirst(data T) {
	a.AddIndex(data, 0)
}

func (a *ArrayList[T]) AddIndex(data T, index int) {
	if index < 0 {
		return
	} else if index == 0 {
		a.elementData = append([]T{data}, a.elementData...)
	} else if index > a.Size() {
		a.Add(data)
	} else {
		a.elementData = append(a.elementData[:index], append([]T{data}, a.elementData[index:]...)...)
	}
}

func (a *ArrayList[T]) Get() (T, error) {
	if a.Size() > 0 {
		return a.elementData[a.Size()-1], nil
	}
	return *new(T), errors.New("empty List")
}

func (a *ArrayList[T]) GetFirst() (T, error) {
	if a.Size() > 0 {
		return a.elementData[0], nil
	}
	return *new(T), errors.New("empty List")
}

func (a *ArrayList[T]) GetLast() (T, error) {
	return a.Get()
}

func (a *ArrayList[T]) GetIndex(index int) (T, error) {
	if index >= 0 && index < a.Size() {
		return a.elementData[index], nil
	}
	return *new(T), errors.New(fmt.Sprintf("index out of range [%d] with length %d", index, a.Size()))
}

func (a *ArrayList[T]) Remove() {
	if a.Size() > 0 {
		a.elementData = a.elementData[:a.Size()-1]
	}
}

func (a *ArrayList[T]) Contains(data T) bool {
	dataValue := reflect.ValueOf(data)
	for _, node := range a.elementData {
		nodeValue := reflect.ValueOf(node)
		if nodeValue.Kind() == dataValue.Kind() && nodeValue.Interface() == dataValue.Interface() {
			return true
		}
	}
	return false
}

func (a *ArrayList[T]) Size() int {
	return len(a.elementData)
}

func (a *ArrayList[T]) ToArray() []T {
	return append([]T{}, a.elementData...)
}

func (a *ArrayList[T]) Clear() {
	a.elementData = newNodeSlice[T](defaultCapacity)
}

func (a *ArrayList[T]) IsEmpty() bool {
	return len(a.elementData) > 0
}

func newNodeSlice[T any](capacity int) nodeSlice[T] {
	return make(nodeSlice[T], 0, capacity)
}

func NewArrayList[T any]() *ArrayList[T] {
	l := &ArrayList[T]{}
	l.elementData = newNodeSlice[T](defaultCapacity)
	return l
}
