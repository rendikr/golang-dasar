package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("John")
	data.PushBack("Doe")
	data.PushBack("Lorem")
	data.PushBack("Ipsum")

	var head = data.Front() // John
	fmt.Println(head.Value)

	next := head.Next() // Doe
	fmt.Println(next.Value)

	next = next.Next() // Lorem
	fmt.Println(next.Value)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
