package main

import (
	"slices"
)

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem, index int)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type myList struct {
	//List // Remove me after realization.
	// Place your code here.
	items []ListItem
}

func NewList() List {
	newSlice := make([]ListItem, 1)
	var d = myList{newSlice}
	return d
}

// Длина списка
func (t myList) Len() int {
	return len(t.items)
}

// Первый элемент
func (t myList) Front() *ListItem {
	return &t.items[0]
}

// Последний элемент
func (t myList) Back() *ListItem {
	indexLast := 0
	if len(t.items) != 0 {
		indexLast = len(t.items) - 1
	}

	return &t.items[indexLast]
}

// Добавить элемент в начало
func (t myList) PushFront(v interface{}) *ListItem {
	newItem := v.(ListItem)
	newItem.Next = &t.items[0]
	newItem.Prev = nil
	t.items = append([]ListItem{newItem}, t.items...)
	return t.Front()
}

// Добавить элемент в конец
func (t myList) PushBack(v interface{}) *ListItem {
	slices.Reverse(t.items)
	t.PushFront(v) // !!! а ссылки при реверсе поменяются ???
	slices.Reverse(t.items)
	return t.Back()
}

// Удалить элемент
func (t myList) Remove(i *ListItem, index int) { // todo изначально не было index
	// Поменяем ссылки смежных элементов
	prevItem := i.Prev
	nextItem := i.Next

	if prevItem != nil {
		prevItem.Next = nextItem
	}
	if nextItem != nil {
		nextItem.Prev = prevItem
	}

	// Удалим элемент из слайса
	var newSlice []ListItem
	if index == t.Len()-1 {
		newSlice = t.items[:index]
	} else {
		newSlice = append(t.items[:index], t.items[index+1:]...)
	}

	t.items = newSlice
}

// Переместить элемент в начало
func (t myList) MoveToFront(i *ListItem) {
	// Если элемент и так в начале списка, то ничего не делаем
	if i.Prev == nil {
		return
	}

	// Поменяем ссылки смежных элементов
	prevItem := i.Prev
	nextItem := i.Next

	if prevItem != nil {
		prevItem.Next = nextItem
	}
	if nextItem != nil {
		nextItem.Prev = prevItem
	}

	// Поменяем ссылки на самом элементе
	i.Prev = nil
	i.Next = &t.items[0]

	// Добавим элемент в слайс
	newSlice := make([]ListItem, 1)
	if t.Len() == 0 || t.items == nil {
		newSlice[0] = *i
	} else {

		newSlice = append([]ListItem{*i}, t.items[0:]...)
	}
	t.items = newSlice
}
