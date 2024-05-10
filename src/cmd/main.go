package main

import (
	"fmt"
	"simple.collection/src/list"
)

func main() {
	arrayList := list.NewArrayList[int]()
	arrayList.Add(1)
	arrayList.AddLast(2)
	arrayList.AddFirst(-1)
	arrayList.AddIndex(0, 1)
	arrayList.AddIndex(-2, 0)

	fmt.Println(arrayList.Get())
	fmt.Println(arrayList.GetFirst())
	fmt.Println(arrayList.GetLast())
	fmt.Println(arrayList.GetIndex(arrayList.Size() / 2))
	fmt.Println(arrayList.Contains(-2))
	fmt.Println(arrayList.ToArray())

	linkedList := list.NewLinkedList[string]()
	linkedList.Add("d")
	linkedList.AddLast("e")
	linkedList.AddFirst("b")
	linkedList.AddIndex("c", 1)
	linkedList.AddIndex("a", 0)

	fmt.Println(linkedList.Get())
	fmt.Println(linkedList.GetFirst())
	fmt.Println(linkedList.GetLast())
	fmt.Println(linkedList.GetIndex(linkedList.Size() / 2))
	fmt.Println(linkedList.Contains("a"))
	fmt.Println(linkedList.ToArray())

}
