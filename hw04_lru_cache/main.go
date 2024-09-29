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
	t.queue.PushFront(item.Value)

	return item.Value, true
}

// Добавить значение в кэш по ключу.
func (t *lruCache) Set(key Key, value interface{}) bool {
	item := t.items[key]

	// Добавляемый элемент отсутствует в словаре.
	if item == nil {
		if t.queue.Len() > t.capacity {
			tail := t.queue.Back()
			tailValue := tail.Value
			t.queue.Remove(tail)
			// Удалим значение последнего элемента из словаря.
			for k, v := range t.items {
				if v == tailValue {
					t.items[k] = nil
				}
			}
		}
		t.queue.PushFront(value)
		addedItem := t.queue.Front()
		t.items[key] = addedItem
		return false
	}

	// Добавляемый элемент присутствует в словаре.
	item.Value = value
	t.queue.MoveToFront(item)
	return true
}

// Очистить кэш.
func (t *lruCache) Clear() {
	clear(t.items)
}
