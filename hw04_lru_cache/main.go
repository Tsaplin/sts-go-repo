package main

import (
	"fmt"
)

func main() {
	fmt.Println("hw04_lru_cache - function main")

	//newSlice := make([]int, 1)
	//newSlice[0] = ListItem{nil, nil, nil}
	//fmt.Println("newSlice = ", newSlice)

	//var t = myList{nil, newSlice}
	t := NewList()
	//var res ListItem = *t.PushFront(10) // [10]
	//t.PushFront(10) // [10]
	//t.PushFront(20) // [20; 10]
	//t.PushFront(30) // [30; 20; 10]

	t.PushBack(10) // [10]
	t.PushBack(20) // [10; 20]
	t.PushBack(30) // [10; 20; 30]

	//var delElem ListItem = *t.Front()
	//t.Remove(&delElem, 0)

	var elem ListItem = *t.Back()
	t.MoveToFront(&elem, 2)

	var res ListItem = *t.Front()
	res = *t.Back()

	var a int = t.Len()
	fmt.Println("a = ", a)
	fmt.Println("res = ", res)
	fmt.Println("t.Len() = ", t.Len())
}

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache // Remove me after realization.

	capacity int
	queue    List
	items    map[Key]*ListItem
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
