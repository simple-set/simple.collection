package list

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestLinkedList(t *testing.T) {
	fmt.Println()
	log.Println("------------Test LinkedList------------")
}

func TestLinkedList_Add(t *testing.T) {
	array := NewLinkedList[int]()
	array.Add(1)
	array.Add(2)

	if array.Size() != 2 {
		t.Fail()
	}
	if value, ok := array.GetIndex(0); ok == nil && value != 1 {
		t.Fail()
	}
	if value, ok := array.GetIndex(1); ok == nil && value != 2 {
		t.Fail()
	}
}

func TestLinkedList_AddLast(t *testing.T) {
	array := NewLinkedList[int]()
	array.AddLast(1)
	array.AddLast(2)

	if array.Size() != 2 {
		t.Fail()
	}
	if value, ok := array.GetIndex(0); ok == nil && value != 1 {
		t.Fail()
	}
	if value, ok := array.GetIndex(1); ok == nil && value != 2 {
		t.Fail()
	}
}

func TestLinkedList_AddFirst(t *testing.T) {
	array := NewLinkedList[int]()
	array.AddFirst(1)
	array.AddFirst(2)
	array.AddFirst(3)

	if array.Size() != 3 {
		t.Fail()
	}
	if value, ok := array.GetIndex(0); ok == nil && value != 3 {
		t.Fail()
	}
	if value, ok := array.GetIndex(3); ok == nil && value != 1 {
		t.Fail()
	}
}

func TestLinkedList_AddIndex(t *testing.T) {
	array := NewLinkedList[string]()
	array.AddIndex("a", 0)
	array.AddIndex("b", 1)
	array.AddIndex("c", 1)
	array.AddIndex("d", 10)

	if array.Size() != 4 {
		t.Fail()
	}
	if value, ok := array.GetIndex(0); ok == nil && value != "a" {
		t.Fail()
	}
	if value, ok := array.GetIndex(1); ok == nil && value != "c" {
		t.Fail()
	}
	if value, ok := array.GetIndex(array.Size() - 1); ok == nil && value != "d" {
		t.Fail()
	}
}

func TestLinkedList_Get(t *testing.T) {
	array := NewLinkedList[int]()
	array.Add(1)
	array.Add(2)

	if value, ok := array.Get(); ok == nil && value != 2 {
		t.Fail()
	}
}

func TestLinkedList_GetFirst(t *testing.T) {
	array := NewLinkedList[int]()
	array.Add(1)
	array.Add(2)

	if value, ok := array.GetFirst(); ok == nil && value != 1 {
		t.Fail()
	}
}

func TestLinkedList_GetLast(t *testing.T) {
	array := NewLinkedList[int]()
	array.Add(1)
	array.Add(2)

	if value, ok := array.GetLast(); ok == nil && value != 2 {
		t.Fail()
	}
}

func TestLinkedList_GetIndex(t *testing.T) {
	array := NewLinkedList[int]()
	array.Add(1)
	array.Add(2)
	array.Add(3)

	if value, ok := array.GetIndex(0); ok == nil && value != 1 {
		t.Fail()
	}
	if value, ok := array.GetIndex(2); ok == nil && value != 3 {
		t.Fail()
	}
	if _, ok := array.GetIndex(10); ok == nil {
		t.Fail()
	}
}

func TestLinkedList_Remove(t *testing.T) {
	array := NewLinkedList[int]()
	array.Add(1)
	array.Add(2)
	array.Add(3)

	array.Remove()
	if array.Size() != 2 {
		t.Fail()
	}
	if value, ok := array.GetIndex(1); ok == nil && value != 2 {
		t.Fail()
	}
}

func TestLinkedList_Contains(t *testing.T) {
	array := NewLinkedList[int]()
	array.Add(1)

	if array.Contains(1) == false {
		t.Fail()
	}
	if array.Contains(2) == true {
		t.Fail()
	}
}

func TestLinkedList_ToArray(t *testing.T) {
	array := NewLinkedList[int]()
	array.Add(1)
	array.Add(2)
	array.Add(3)

	ints := array.ToArray()
	if len(ints) != 3 {
		t.Fail()
	}
	if ints[0] != 1 || ints[1] != 2 || ints[2] != 3 {
		t.Fail()
	}
}

func TestLinkedList_Clear(t *testing.T) {
	array := NewLinkedList[int]()
	array.Add(1)
	array.Add(2)
	array.Add(3)

	array.Clear()
	if array.Size() != 0 {
		t.Fail()
	}
}

// -----------------Benchmark-------------
func linkedListAdd(count int, random bool) {
	rand.Seed(time.Now().UnixNano())
	array := NewLinkedList[int]()
	for i := 0; i < count; i++ {
		if random {
			array.AddIndex(i, rand.Intn(i+1))
		} else {
			array.Add(i)
		}
	}
}

// 顺序写入
func BenchmarkLinkedList_add1000(b *testing.B) {
	linkedListAdd(1000, false)
}
func BenchmarkLinkedList_add10000(b *testing.B) {
	linkedListAdd(10000, false)
}
func BenchmarkLinkedList_add100000(b *testing.B) {
	linkedListAdd(100000, false)
}

// 随机写入
func BenchmarkLinkedList_addRandom1000(b *testing.B) {
	linkedListAdd(1000, true)
}
func BenchmarkLinedList_addRandom10000(b *testing.B) {
	linkedListAdd(10000, true)
}
func BenchmarkLinkedList_addRandom100000(b *testing.B) {
	linkedListAdd(100000, true)
}

func LinkedListGetIndex(b *testing.B, count int, random bool) {
	array := NewArrayList[int]()
	for n := 0; n < 1000; n++ {
		array.Add(n)
	}
	rand.Seed(time.Now().UnixNano())

	// 重置时间
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for n := 0; n < count; n++ {
			if random {
				_, _ = array.GetIndex(rand.Intn(1000))
			} else {
				_, _ = array.GetIndex(n % 1000)
			}
		}
	}
}

// 顺序读取
func BenchmarkLinkedList_GetIndex1000(b *testing.B) {
	LinkedListGetIndex(b, 1000, false)
}
func BenchmarkLinkedList_GetIndex10000(b *testing.B) {
	LinkedListGetIndex(b, 10000, false)
}
func BenchmarkLinkedList_GetIndex100000(b *testing.B) {
	LinkedListGetIndex(b, 100000, false)
}

// 随机读取
func BenchmarkLinkedList_GetIndexRandom1000(b *testing.B) {
	LinkedListGetIndex(b, 1000, true)
}
func BenchmarkLinkedList_GetIndexRandom10000(b *testing.B) {
	LinkedListGetIndex(b, 10000, true)
}
func BenchmarkLinkedList_GetIndexRandom100000(b *testing.B) {
	LinkedListGetIndex(b, 100000, true)
}
