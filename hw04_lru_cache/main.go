package main

import (
	"fmt"
)

func main() {
	fmt.Println("hw04_lru_cache - function main")
}

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type dataStruct struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

// Получить значение из кэша по ключу.
func (t *lruCache) Get(key Key) (interface{}, bool) {
	item := t.items[key]
	if item == nil {
		return nil, false
	}
	dynamicValue := item.Value.(dataStruct)
	t.queue.PushFront(dynamicValue)

	return dynamicValue.value, true
}

// Добавить значение в кэш по ключу.
func (t *lruCache) Set(key Key, value interface{}) bool {
	item := t.items[key]
	data := dataStruct{key: key, value: value.(int)}

	// Добавляемый элемент отсутствует в словаре.
	if item == nil {
		if t.queue.Len() >= t.capacity {
			tail := t.queue.Back()
			tailValue := tail.Value
			t.queue.Remove(tail)
			// Удалим значение последнего элемента из словаря.
			delete(t.items, tailValue.(dataStruct).key)
		}
		t.queue.PushFront(data)
		addedItem := t.queue.Front()
		t.items[key] = addedItem
		return false
	}

	// Добавляемый элемент присутствует в словаре.
	item.Value = data
	t.queue.MoveToFront(item)
	return true
}

// Очистить кэш.
func (t *lruCache) Clear() {
	clear(t.items)
	t.queue = NewList()
}
