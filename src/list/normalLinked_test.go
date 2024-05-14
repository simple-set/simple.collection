package list

import (
	"fmt"
	"log"
	"testing"
)

func TestNormalLinked(t *testing.T) {
	fmt.Println()
	log.Println("------------Test NormalLinked------------")
}

func TestNormalLinked_AddHeadNode(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddHeadNode(2)

	// 断言
	if linked.Size() != 2 {
		t.Fail()
	}
	if linked.list.head.data != 2 {
		t.Fail()
	}
	if linked.list.tail.data != 1 {
		t.Fail()
	}
}

func TestNormalLinked_AddTailNode(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddTailNode(2)

	// 断言
	if linked.Size() != 2 {
		t.Fail()
	}
	if linked.list.head.data != 1 {
		t.Fail()
	}
	if linked.list.tail.data != 2 {
		t.Fail()
	}
}

func TestNormalLinked_AddNextNode_nil(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddNextNode(1)

	// 断言
	if linked.Size() != 0 {
		t.Fail()
	}
}

func TestNormalLinked_AddNextNode(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)

	headNode := linked.GetHeadNode()
	headNode.AddNextNode(3)
	headNode.AddNextNode(2)

	// 断言
	if linked.Size() != 3 {
		t.Fail()
	}
	current := linked.GetHeadNode()
	var ints = [...]int{1, 2, 3}
	for _, i := range ints {
		if value, err := current.GetValue(); err != nil || value != i {
			t.Fail()
		}
		current = current.GetNextNode()
	}
}

func TestNormalLinked_AddNextNode_after_tail(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddTailNode(2)

	linked.GetTailNode().AddNextNode(3)
	linked.GetTailNode().AddNextNode(4)

	// 断言
	if linked.Size() != 4 {
		t.Fail()
	}
	current := linked.GetHeadNode()
	var ints = [...]int{1, 2, 3, 4}
	for _, i := range ints {
		if value, err := current.GetValue(); err != nil || value != i {
			t.Fail()
		}
		current = current.GetNextNode()
	}
}

func TestNormalLinked_AddPervNode_after_head(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddTailNode(2)

	linked.GetHeadNode().AddPrevNode(0)
	linked.GetHeadNode().AddPrevNode(-1)

	if linked.Size() != 4 {
		t.Fail()
	}
	current := linked.GetHeadNode()
	var ints = [...]int{-1, 0, 1, 2}
	for _, i := range ints {
		if value, err := current.GetValue(); err != nil || value != i {
			t.Fail()
		}
		current = current.GetNextNode()
	}
}

func TestNormalLinked_AddPervNode_after_tail(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddTailNode(4)

	linked.GetTailNode().AddPrevNode(2)
	linked.GetTailNode().AddPrevNode(3)

	if linked.Size() != 4 {
		t.Fail()
	}
	current := linked.GetHeadNode()
	var ints = [...]int{1, 2, 3, 4}
	for _, i := range ints {
		if value, err := current.GetValue(); err != nil || value != i {
			t.Fail()
		}
		current = current.GetNextNode()
	}
}

func TestNormalLinked_GetHeadNode(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddHeadNode(0)

	// 断言
	if value, ok := linked.GetHeadNode().GetValue(); ok != nil || value != 0 {
		t.Fail()
	}
}

func TestNormalLinked_GetHeadNode_nil(t *testing.T) {
	linked := NewNormalLinked[int]()
	node := linked.GetHeadNode()
	if node != nil {
		t.Fail()
	}
}

func TestNormalLinked_GetTailNode(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddHeadNode(0)

	// 断言
	if value, ok := linked.GetTailNode().GetValue(); ok != nil || value != 1 {
		t.Fail()
	}
}

func TestNormalLinked_GetTailNode_nil(t *testing.T) {
	linked := NewNormalLinked[int]()
	node := linked.GetTailNode()
	if node != nil {
		t.Fail()
	}
}

func TestNormalLinked_GetNextNode(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddHeadNode(0)

	headNode := linked.GetHeadNode()
	nextNode := headNode.GetNextNode()
	if value, err := nextNode.GetValue(); err != nil || value != 1 {
		t.Fail()
	}
}

func TestNormalLinked_GetNextNode_after_tail(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddHeadNode(0)

	tailNode := linked.GetTailNode()
	nextNode := tailNode.GetNextNode()
	if nextNode != nil {
		t.Fail()
	}
}

func TestNormalLinked_GetNextNode_nil(t *testing.T) {
	linked := NewNormalLinked[int]()
	if linked.GetNextNode() != nil {
		t.Fail()
	}
}

func TestNormalLinked_GetPrevNode(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddHeadNode(0)

	tailNode := linked.GetTailNode()
	prevNode := tailNode.GetPrevNode()
	if value, err := prevNode.GetValue(); err != nil || value != 0 {
		t.Fail()
	}
}

func TestNormalLinked_GetPrevNode_after_head(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddHeadNode(0)

	headNode := linked.GetHeadNode()
	prevNode := headNode.GetPrevNode()
	if prevNode != nil {
		t.Fail()
	}
}

func TestNormalLinked_GetPrevNode_nil(t *testing.T) {
	linked := NewNormalLinked[int]()
	if linked.GetPrevNode() != nil {
		t.Fail()
	}
}

func TestNormalLinked_RemoveNode(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddTailNode(2)
	linked.AddTailNode(3)

	nextNode := linked.GetHeadNode().GetNextNode()
	nextNode.RemoveNode()

	if linked.Size() != 2 {
		t.Fail()
	}

	headNode := linked.GetHeadNode()
	if value, err := headNode.GetValue(); err != nil || value != 1 {
		t.Fail()
	}
	if headNode.GetPrevNode() != nil {
		t.Fail()
	}

	tailNode := headNode.GetNextNode()
	if value, err := tailNode.GetValue(); err != nil || value != 3 {
		t.Fail()
	}
	if tailNode.GetNextNode() != nil {
		t.Fail()
	}
	if value, err := tailNode.GetPrevNode().GetValue(); err != nil || value != 1 {
		t.Fail()
	}
}

func TestNormalLinked_RemoveNode_head(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddTailNode(2)
	linked.AddTailNode(3)

	node := linked.GetHeadNode()
	node.RemoveNode()

	if linked.Size() != 2 {
		t.Fail()
	}

	headNode := linked.GetHeadNode()
	if value, err := headNode.GetValue(); err != nil || value != 2 {
		t.Fail()
	}
	if headNode.GetPrevNode() != nil {
		t.Fail()
	}

	tailNode := headNode.GetNextNode()
	if value, err := tailNode.GetValue(); err != nil || value != 3 {
		t.Fail()
	}
	if tailNode.GetNextNode() != nil {
		t.Fail()
	}
	if value, err := tailNode.GetPrevNode().GetValue(); err != nil || value != 2 {
		t.Fail()
	}
}

func TestNormalLinked_RemoveNode_tail(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddTailNode(2)
	linked.AddTailNode(3)

	node := linked.GetTailNode()
	node.RemoveNode()

	if linked.Size() != 2 {
		t.Fail()
	}

	headNode := linked.GetHeadNode()
	if value, err := headNode.GetValue(); err != nil || value != 1 {
		t.Fail()
	}
	if headNode.GetPrevNode() != nil {
		t.Fail()
	}

	tailNode := headNode.GetNextNode()
	if value, err := tailNode.GetValue(); err != nil || value != 2 {
		t.Fail()
	}
	if tailNode.GetNextNode() != nil {
		t.Fail()
	}
	if value, err := tailNode.GetPrevNode().GetValue(); err != nil || value != 1 {
		t.Fail()
	}
}

func TestNormalLinked_RemoveHeadNode(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddTailNode(2)
	linked.AddTailNode(3)

	linked.RemoveHeadNode()

	if linked.Size() != 2 {
		t.Fail()
	}

	headNode := linked.GetHeadNode()
	if value, err := headNode.GetValue(); err != nil || value != 2 {
		t.Fail()
	}
	if linked.list.head.prev != nil {
		t.Fail()
	}
}

func TestNormalLinked_RemoveHeadNode_nil(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.RemoveHeadNode()
}

func TestNormalLinked_RemoveTailNode(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddTailNode(2)
	linked.AddTailNode(3)

	linked.RemoveTailNode()

	if linked.Size() != 2 {
		t.Fail()
	}

	tailNode := linked.GetTailNode()
	if value, err := tailNode.GetValue(); err != nil || value != 2 {
		t.Fail()
	}
	if linked.list.tail.next != nil || tailNode.GetNextNode() != nil {
		t.Fail()
	}
}

func TestNormalLinked_RemoveTailNode_nil(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.RemoveTailNode()
}

func TestNormalLinked_RemoveNextNode(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddTailNode(2)
	linked.AddTailNode(3)

	node := linked.GetHeadNode().GetNextNode()
	node.RemoveNextNode()

	if linked.Size() != 2 {
		t.Fail()
	}

	tailNode := linked.GetTailNode()
	if value, err := tailNode.GetValue(); err != nil || value != 2 {
		t.Fail()
	}
	if tailNode.GetNextNode() != nil || linked.list.tail.next != nil {
		t.Fail()
	}
}

func TestNormalLinked_RemoveNextNode_after_head(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddTailNode(2)
	linked.AddTailNode(3)

	node := linked.GetHeadNode()
	node.RemoveNextNode()

	if linked.Size() != 2 {
		t.Fail()
	}

	headNode := linked.GetHeadNode()
	if value, err := headNode.GetValue(); err != nil || value != 1 {
		t.Fail()
	}
	if value, err := headNode.GetNextNode().GetValue(); err != nil || value != 3 {
		t.Fail()
	}

	tailNode := linked.GetTailNode()
	if value, err := tailNode.GetValue(); err != nil || value != 3 {
		t.Fail()
	}
	if value, err := tailNode.GetPrevNode().GetValue(); err != nil || value != 1 {
		t.Fail()
	}

}

func TestNormalLinked_RemoveNextNode_after_tail(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddTailNode(2)
	linked.AddTailNode(3)

	node := linked.GetTailNode()
	node.RemoveNextNode()

	if linked.Size() != 3 {
		t.Fail()
	}

	headNode := linked.GetHeadNode()
	if value, err := headNode.GetValue(); err != nil || value != 1 {
		t.Fail()
	}

}

func TestNormalLinked_RemovePrevNode(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddTailNode(2)
	linked.AddTailNode(3)

	node := linked.GetHeadNode().GetNextNode()
	node.RemovePrevNode()

	if linked.Size() != 2 {
		t.Fail()
	}
	headNode := linked.GetHeadNode()
	if value, err := headNode.GetValue(); err != nil || value != 2 {
		t.Fail()
	}
	if headNode.GetPrevNode() != nil || linked.list.head.prev != nil {
		t.Fail()
	}
	if value, err := headNode.GetNextNode().GetValue(); err != nil || value != 3 {
		t.Fail()
	}
	if linked.list.head.data != 2 {
		t.Fail()
	}
	tailNode := linked.GetTailNode()
	if value, err := tailNode.GetValue(); err != nil || value != 3 {
		t.Fail()
	}
}

func TestNormalLinked_RemovePrevNode_after_head(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddTailNode(2)
	linked.AddTailNode(3)

	linked.GetHeadNode().RemovePrevNode()

	if linked.Size() != 3 {
		t.Fail()
	}
	headNode := linked.GetHeadNode()
	if value, err := headNode.GetValue(); err != nil || value != 1 {
		t.Fail()
	}
	if value, err := headNode.GetNextNode().GetValue(); err != nil || value != 2 {
		t.Fail()
	}
}

func TestNormalLinked_RemovePrevNode_after_tail(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)
	linked.AddTailNode(2)
	linked.AddTailNode(3)

	node := linked.GetTailNode()
	node.RemovePrevNode()

	if linked.Size() != 2 {
		t.Fail()
	}

	headNode := linked.GetHeadNode()
	if value, err := headNode.GetValue(); err != nil || value != 1 {
		t.Fail()
	}
	if value, err := headNode.GetNextNode().GetValue(); err != nil || value != 3 {
		t.Fail()
	}

	tailNode := linked.GetTailNode()
	if value, err := tailNode.GetValue(); err != nil || value != 3 {
		t.Fail()
	}
	if value, err := tailNode.GetPrevNode().GetValue(); err != nil || value != 1 {
		t.Fail()
	}
}

func TestNormalLinked_Size(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)

	if linked.Size() != 1 {
		t.Fail()
	}
	linked.GetHeadNode().RemoveHeadNode()
	if linked.Size() != 0 {
		t.Fail()
	}
}

func TestNormalLinked_SetValue(t *testing.T) {
	linked := NewNormalLinked[int]()
	linked.AddHeadNode(1)

	node := linked.GetHeadNode()
	node.SetValue(0)
	if value, err := node.GetValue(); err != nil || value != 0 {
		t.Fail()
	}
}
