package list

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestArrayList(t *testing.T) {
	fmt.Println()
	log.Println("------------Test ArrayList------------")
}

func TestArrayList_Add(t *testing.T) {
	array := NewArrayList[int]()
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

func TestArrayList_AddLast(t *testing.T) {
	array := NewArrayList[int]()
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

func TestArrayList_AddFirst(t *testing.T) {
	array := NewArrayList[int]()
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

func TestArrayList_AddIndex(t *testing.T) {
	array := NewArrayList[string]()
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

func TestArrayList_Get(t *testing.T) {
	array := NewArrayList[int]()
	array.Add(1)
	array.Add(2)

	if value, ok := array.Get(); ok == nil && value != 2 {
		t.Fail()
	}
}

func TestArrayList_GetFirst(t *testing.T) {
	array := NewArrayList[int]()
	array.Add(1)
	array.Add(2)

	if value, ok := array.GetFirst(); ok == nil && value != 1 {
		t.Fail()
	}
}

func TestArrayList_GetLast(t *testing.T) {
	array := NewArrayList[int]()
	array.Add(1)
	array.Add(2)

	if value, ok := array.GetLast(); ok == nil && value != 2 {
		t.Fail()
	}
}

func TestArrayList_GetIndex(t *testing.T) {
	array := NewArrayList[int]()
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

func TestArrayList_Remove(t *testing.T) {
	array := NewArrayList[int]()
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

func TestArrayList_Contains(t *testing.T) {
	array := NewArrayList[int]()
	array.Add(1)

	if array.Contains(1) == false {
		t.Fail()
	}
	if array.Contains(2) == true {
		t.Fail()
	}
}

func TestArrayList_ToArray(t *testing.T) {
	array := NewArrayList[int]()
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

func TestArrayList_Clear(t *testing.T) {
	array := NewArrayList[int]()
	array.Add(1)
	array.Add(2)
	array.Add(3)

	array.Clear()
	if array.Size() != 0 {
		t.Fail()
	}
}

// -----------------Benchmark 写入-------------
func arraylistAdd(count int, random bool) {
	rand.Seed(time.Now().UnixNano())
	array := NewArrayList[int]()
	for i := 0; i < count; i++ {
		if random {
			array.AddIndex(i, rand.Intn(i+1))
		} else {
			array.Add(i)
		}
	}
}

// 顺序写入
func BenchmarkArrayList_add1000(b *testing.B) {
	arraylistAdd(1000, false)
}
func BenchmarkArrayList_add10000(b *testing.B) {
	arraylistAdd(10000, false)
}

// 随机写入
func BenchmarkArrayList_addRandom1000(b *testing.B) {
	arraylistAdd(1000, true)
}
func BenchmarkArrayList_addRandom10000(b *testing.B) {
	arraylistAdd(10000, true)
}

// -----------------Benchmark 读取-------------
func ArrayListGetIndex(b *testing.B, count int, random bool) {
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
func BenchmarkArrayList_GetIndex1000(b *testing.B) {
	ArrayListGetIndex(b, 1000, false)
}
func BenchmarkArrayList_GetIndex10000(b *testing.B) {
	ArrayListGetIndex(b, 10000, false)
}

// 随机读取
func BenchmarkArrayList_GetIndexRandom1000(b *testing.B) {
	ArrayListGetIndex(b, 1000, true)
}
func BenchmarkArrayList_GetIndexRandom10000(b *testing.B) {
	ArrayListGetIndex(b, 10000, true)
}
