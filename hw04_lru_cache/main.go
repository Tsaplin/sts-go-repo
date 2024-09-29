package main

import (
	"fmt"
)

func main() {
	fmt.Println("hw04_lru_cache - function main")

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
	t.MoveToFront(&elem)

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
	//Cache // Remove me after realization.

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

// Получить значение из кэша по ключу
func (t *lruCache) Get(key Key) (interface{}, bool) {
	var item *ListItem = t.items[key]
	if item == nil {
		return nil, false
	}
	t.queue.PushFront(item.Value)

	return item.Value, true
}

// Добавить значение в кэш по ключу
func (t *lruCache) Set(key Key, value interface{}) bool {
	var item *ListItem = t.items[key]

	// Добавляемый элемент отсутствует в словаре
	if item == nil {
		if t.queue.Len() > t.capacity {
			tail := t.queue.Back()
			t.queue.Remove(tail)
			//t.items[] = nil
			delete(t.items, key) // ??? Какой ключ у последнего элемента ?
		}
		t.queue.PushFront(value)
		addedItem := t.queue.Front()
		t.items[key] = addedItem
		return false
	}

	// Добавляемый элемент присутствует в словаре
	item.Value = value
	t.queue.MoveToFront(item)
	return true
}

// Очистить кэш
func (t *lruCache) Clear() {
	// ??? Нужно ли чистить связный список
	clear(t.items)
}
