package main

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type myList struct {
	length int
	head   *ListItem
	tail   *ListItem
}

func NewList() List {
	return &myList{}
}

// Длина списка.
func (t *myList) Len() int {
	return t.length
}

// Первый элемент.
func (t *myList) Front() *ListItem {
	if t.length > 0 {
		return t.head
	}

	return nil
}

// Последний элемент.
func (t *myList) Back() *ListItem {
	if t.length != 0 {
		return t.tail
	}

	return nil
}

// Добавить элемент в начало.
func (t *myList) PushFront(v interface{}) *ListItem {
	var newItem = ListItem{nil, nil, nil}
	newItem.Value = v
	newItem.Prev = nil
	if t.length > 0 {
		newItem.Next = t.head
		newItem.Next.Prev = &newItem
	} else {
		t.tail = &newItem
	}

	t.length++
	t.head = &newItem

	return t.Front()
}

// Добавить элемент в конец.
func (t *myList) PushBack(v interface{}) *ListItem {
	var newItem = ListItem{nil, nil, nil}
	newItem.Value = v
	newItem.Next = nil

	if t.length > 0 {
		newItem.Prev = t.tail
		newItem.Prev.Next = &newItem
	} else {
		t.head = &newItem
	}

	t.length++
	t.tail = &newItem

	return t.Back()
}

// Удалить элемент.
func (t *myList) Remove(i *ListItem) {
	// Поменяем ссылки смежных элементов.
	prevItem := i.Prev
	nextItem := i.Next

	if prevItem != nil {
		prevItem.Next = nextItem
	}
	if nextItem != nil {
		nextItem.Prev = prevItem
	}

	t.length--
	if i.Next == nil {
		t.tail = i.Prev
	}
	if i.Prev == nil {
		t.head = i.Next
	}
}

// Переместить элемент в начало.
func (t *myList) MoveToFront(i *ListItem) {
	// Если элемент и так в начале списка, то ничего не делаем.
	if i.Prev == nil {
		return
	}

	t.Remove(i)
	t.PushFront(i.Value)
}
