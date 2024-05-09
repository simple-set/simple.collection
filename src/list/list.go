package list

import "simple.collection/src"

type List[T any] interface {
	AddLast(data T)
	AddFirst(data T)
	AddIndex(data T, index int)
	Get() (T, error)
	GetLast() (T, error)
	GetFirst() (T, error)
	GetIndex(index int) (T, error)
	src.Collection[T]
}
