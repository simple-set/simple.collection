package list

import "errors"

type NormalLinked[T any] struct {
	list    *LinkedList[T]
	current *dataNode[T]
}

func (n *NormalLinked[T]) AddHeadNode(data T) {
	n.list.AddFirst(data)

}

func (n *NormalLinked[T]) AddTailNode(data T) {
	n.list.AddLast(data)
}

func (n *NormalLinked[T]) AddNextNode(data T) {
	if n.current != nil {
		newNode := newDataNode(data, n.current.next, n.current)
		if n.current.next != nil {
			n.current.next.prev = newNode
		} else {
			n.list.tail = newNode
		}
		n.current.next = newNode
		n.list.size++
	}
}

func (n *NormalLinked[T]) AddPrevNode(data T) {
	if n.current == nil {
		n.list.AddFirst(data)
		return
	}

	newNode := newDataNode(data, n.current, n.current.prev)
	if n.current.prev != nil {
		n.current.prev.next = newNode
	} else {
		n.list.head = newNode
	}
	n.current.prev = newNode
	n.list.size++
}

func (n *NormalLinked[T]) GetHeadNode() Linked[T] {
	if n.list.head != nil {
		return NewNodeLinked(n.list.head, n.list)
	}
	return nil
}

func (n *NormalLinked[T]) GetTailNode() Linked[T] {
	if n.list.tail != nil {
		return NewNodeLinked(n.list.tail, n.list)
	}
	return nil
}

func (n *NormalLinked[T]) GetNextNode() Linked[T] {
	if n.current == nil || n.current.next == nil {
		return nil
	} else {
		return NewNodeLinked(n.current.next, n.list)
	}
}

func (n *NormalLinked[T]) GetPrevNode() Linked[T] {
	if n.current == nil || n.current.prev == nil {
		return nil
	} else {
		return NewNodeLinked(n.current.prev, n.list)
	}
}

func (n *NormalLinked[T]) RemoveNode() {
	if n.current != nil {
		if n.current.prev == nil {
			n.list.head = n.current.next
		} else {
			n.current.prev.next = n.current.next
		}

		if n.current.next == nil {
			n.list.tail = n.current.prev
		} else {
			n.current.next.prev = n.current.prev
		}
		n.list.size--
	}
}

func (n *NormalLinked[T]) RemoveHeadNode() {
	head := n.list.head
	if head != nil {
		if head.next != nil {
			head.next.prev = nil
			n.list.head = head.next
		} else {
			n.list.head = nil
		}

		if n.list.tail == head {
			n.list.tail = nil
		}
		n.list.size--
	}
}

func (n *NormalLinked[T]) RemoveTailNode() {
	tail := n.list.tail
	if tail != nil {
		if tail.prev != nil {
			tail.prev.next = nil
			n.list.tail = tail.prev
		} else {
			n.list.tail = nil
		}

		if n.list.head == tail {
			n.list.head = nil
		}
		n.list.size--
	}
}

func (n *NormalLinked[T]) RemoveNextNode() {
	if n.current != nil && n.current.next != nil {
		if n.current.next.next != nil {
			n.current.next.next.prev = n.current
		}
		if n.current.next.prev != nil {
			n.current.next.prev.next = n.current.next.next
		}
		if n.current.next == nil {
			n.list.tail = n.current
		}
		n.list.size--
	}
}

func (n *NormalLinked[T]) RemovePrevNode() {
	if n.current != nil && n.current.prev != nil {
		if n.current.prev.prev != nil {
			n.current.prev.prev.next = n.current
			n.current.prev = n.current.prev.prev
		} else {
			n.current.prev = nil
			n.list.head = n.current
		}
		n.list.size--
	}
}

func (n *NormalLinked[T]) Size() int {
	return n.list.Size()
}

func (n *NormalLinked[T]) GetValue() (T, error) {
	if n.current != nil {
		return n.current.data, nil
	}
	return *new(T), errors.New("current is nil")
}

func (n *NormalLinked[T]) SetValue(value T) {
	if n.current != nil {
		n.current.data = value
	}
}

func NewNodeLinked[T any](node *dataNode[T], list *LinkedList[T]) *NormalLinked[T] {
	return &NormalLinked[T]{current: node, list: list}
}

func NewNormalLinked[T any]() *NormalLinked[T] {
	return &NormalLinked[T]{list: NewLinkedList[T]()}
}
