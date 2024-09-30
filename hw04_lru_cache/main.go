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
	dynamicValue := item.Value.(map[Key]interface{})
	t.queue.PushFront(dynamicValue)

	return dynamicValue[key], true
}

// Добавить значение в кэш по ключу.
func (t *lruCache) Set(key Key, value interface{}) bool {
	item := t.items[key]

	// Добавляемый элемент отсутствует в словаре.
	if item == nil {
		if t.queue.Len() >= t.capacity {
			tail := t.queue.Back()
			tailValue := tail.Value
			t.queue.Remove(tail)
			// Удалим значение последнего элемента из словаря.
			for k := range tailValue.(map[Key]interface{}) {
				delete(t.items, k)
			}
		}
		tempMap := make(map[Key]interface{})
		dynamicValue := value.(int)
		tempMap[key] = dynamicValue
		t.queue.PushFront(tempMap)
		addedItem := t.queue.Front()
		t.items[key] = addedItem
		return false
	}

	// Добавляемый элемент присутствует в словаре.
	tempMap := make(map[Key]interface{})
	tempMap[key] = value
	item.Value = tempMap
	t.queue.MoveToFront(item)
	return true
}

// Очистить кэш.
func (t *lruCache) Clear() {
	clear(t.items)
	t.queue = NewList()
}
