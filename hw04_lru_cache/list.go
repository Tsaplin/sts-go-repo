package main

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem, index int)      // todo index
	MoveToFront(i *ListItem, index int) // todo index
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
	newSlice := make([]ListItem, 0)
	var d = myList{items: newSlice}
	return &d
}

// Длина списка
func (t *myList) Len() int {
	var listLength int = len(t.items)
	return listLength
}

// Первый элемент
func (t *myList) Front() *ListItem {
	if len(t.items) > 0 {
		return &t.items[0]
	}

	return nil
}

// Последний элемент
func (t *myList) Back() *ListItem {
	indexLast := 0
	if len(t.items) != 0 {
		indexLast = len(t.items) - 1
		return &t.items[indexLast]
	}

	return nil
}

// Добавить элемент в начало
func (t *myList) PushFront(v interface{}) *ListItem {
	var newItem = ListItem{nil, nil, nil}
	newItem.Value = v
	newItem.Prev = nil
	if len(t.items) > 0 {
		newItem.Next = &t.items[0]
		t.items[0].Prev = &newItem
	}

	t.items = append([]ListItem{newItem}, t.items...)
	return t.Front()
}

// Добавить элемент в конец
func (t *myList) PushBack(v interface{}) *ListItem {
	var newItem = ListItem{nil, nil, nil}
	newItem.Value = v
	newItem.Next = nil

	listLength := len(t.items)
	if listLength > 0 {
		newItem.Prev = &t.items[listLength-1]
		t.items[listLength-1].Next = &newItem
	}

	t.items = append(t.items, []ListItem{newItem}...)

	return t.Back()
}

// Удалить элемент
func (t *myList) Remove(i *ListItem, index int) { // todo изначально не было index
	// Поменяем ссылки смежных элементов
	prevItem := i.Prev
	nextItem := i.Next

	if prevItem != nil {
		t.items[index-1].Next = nextItem
	}
	if nextItem != nil {
		t.items[index+1].Prev = prevItem
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
func (t *myList) MoveToFront(i *ListItem, index int) {
	// Если элемент и так в начале списка, то ничего не делаем
	if i.Prev == nil {
		return
	}

	t.Remove(i, index)
	t.PushFront(i.Value)
}
