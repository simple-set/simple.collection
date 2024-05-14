package list

import (
	"errors"
	"reflect"
)

type dataNode[T any] struct {
	data T
	next *dataNode[T]
	prev *dataNode[T]
}

func newDataNode[T any](data T, next *dataNode[T], prev *dataNode[T]) *dataNode[T] {
	return &dataNode[T]{data: data, next: next, prev: prev}
}

type LinkedList[T any] struct {
	head *dataNode[T]
	tail *dataNode[T]
	size int
}

func (l *LinkedList[T]) Add(data T) {
	if l.tail == nil {
		l.head = newDataNode(data, nil, nil)
		l.tail = l.head
	} else {
		newNode := newDataNode(data, nil, l.tail)
		l.tail.next = newNode
		l.tail = newNode
	}
	l.size++
}

func (l *LinkedList[T]) AddLast(data T) {
	l.Add(data)
}

func (l *LinkedList[T]) AddFirst(data T) {
	if l.head == nil {
		l.Add(data)
	} else {
		newNode := newDataNode(data, l.head, nil)
		l.head.prev = newNode
		l.head = newNode
		l.size++
	}
}

func (l *LinkedList[T]) AddIndex(data T, index int) {
	if index < 0 {
		return
	} else if index == 0 {
		l.AddFirst(data)
	} else if index >= l.Size() {
		l.AddLast(data)
	} else {
		current := l.head
		for i := 0; i < index; i++ {
			current = current.next
		}
		newNode := newDataNode(data, current, current.prev)
		current.prev.next = newNode
		current.prev = newNode
		l.size++
	}
}

func (l *LinkedList[T]) Get() (T, error) {
	if l.tail != nil {
		return l.tail.data, nil
	}
	return *new(T), errors.New("empty List")
}

func (l *LinkedList[T]) GetLast() (T, error) {
	return l.Get()
}

func (l *LinkedList[T]) GetFirst() (T, error) {
	if l.head != nil {
		return l.head.data, nil
	}
	return *new(T), errors.New("empty List")
}

func (l *LinkedList[T]) GetIndex(index int) (T, error) {
	if l.Size() <= 0 {
		return *new(T), errors.New("empty List")
	}
	if index < 0 || index >= l.size {
		return *new(T), errors.New("index out of bounds")
	}

	current := l.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current.data, nil
}

func (l *LinkedList[T]) Remove() {
	if l.tail != nil {
		l.tail = l.tail.prev
		l.tail.next = nil
		l.size--
	}
}

func (l *LinkedList[T]) Contains(data T) bool {
	if l.Size() <= 0 {
		return false
	}

	dataValue := reflect.ValueOf(data)
	current := l.head
	for {
		nodeValue := reflect.ValueOf(current.data)
		if dataValue.Kind() == nodeValue.Kind() && dataValue.Interface() == nodeValue.Interface() {
			return true
		}
		if current.next == nil {
			return false
		}
		current = current.next
	}
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) ToArray() []T {
	if l.head == nil {
		return nil
	}

	var arr []T
	current := l.head
	for {
		arr = append(arr, current.data)
		if current.next == nil {
			break
		}
		current = current.next
	}
	return arr
}

func (l *LinkedList[T]) Clear() {
	l.size = 0
	l.head = nil
	l.tail = nil
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.size > 0
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}
