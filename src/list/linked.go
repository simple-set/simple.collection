package list

// Linked 链表接口
type Linked[T any] interface {
	AddHeadNode(data T)
	AddTailNode(data T)
	AddNextNode(data T)
	AddPrevNode(data T)

	GetHeadNode() Linked[T]
	GetTailNode() Linked[T]
	GetNextNode() Linked[T]
	GetPrevNode() Linked[T]

	RemoveNode()
	RemoveHeadNode()
	RemoveTailNode()
	RemoveNextNode()
	RemovePrevNode()

	Size() int
	GetValue() (T, error)
	SetValue(value T)
}
