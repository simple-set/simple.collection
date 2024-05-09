package src

type Collection[T any] interface {
	Add(data T)
	Remove()
	Contains(data T) bool
	Size() int
	ToArray() []T
	Clear()
	IsEmpty() bool
}
