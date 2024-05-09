package main

import (
	"fmt"
	"simple.collection/src/list"
)

func f1(l list.List[int]) {
	fmt.Println(l.Get())
	fmt.Println(l.GetIndex(l.Size()))
}
func equal[T comparable](a, b T) bool {
	return a == b
}

type MyStruct struct {
	Name string
	Age  int
}

func main() {
	arrayList := list.NewArrayList[int]()
	arrayList.Add(1)
	arrayList.AddLast(2)
	arrayList.AddFirst(-1)
	arrayList.AddIndex(0, 1)
	arrayList.AddIndex(-2, 0)

	println(arrayList.Get())
	println(arrayList.GetFirst())
	println(arrayList.GetLast())
	println(arrayList.GetIndex(arrayList.Size() - 1))
	println(arrayList.Contains(-2))
	array := arrayList.ToArray()
	fmt.Println(array)
	//array[0] = 10
	//f1(arrayList)

	linkedList := list.NewLinkedList[int]()
	linkedList.Add(1)
	linkedList.AddLast(2)
	linkedList.AddFirst(-1)
	linkedList.AddIndex(0, 1)
	linkedList.AddIndex(-2, 0)

	println(linkedList.Get())
	println(linkedList.GetFirst())
	println(linkedList.GetLast())
	println(linkedList.GetIndex(linkedList.Size() - 1))
	println(linkedList.Contains(-2))
	array2 := linkedList.ToArray()
	fmt.Println(array2)

	//value := reflect.ValueOf(nil)
	//fmt.Println(value)
	//typeOf := reflect.TypeOf(arrayList)
	//fmt.Println(typeOf.Kind())
	//fmt.Println(typeOf.Name())
	//var i []int
	//i := []int{1, 2, 3}
	//i[0] = 1
	//fmt.Println(i[3])
	//var j []int
	//m := MyStruct{Age: 10}
	//m2 := MyStruct{}
	//fmt.Println(m == m2)
}
