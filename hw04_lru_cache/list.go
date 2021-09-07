package hw04lrucache

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

type list struct {
	FrontField *ListItem
	BackField  *ListItem
	LenField   int
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int {
	return l.LenField
}

func (l *list) Front() *ListItem {
	return l.FrontField
}

func (l *list) Back() *ListItem {
	return l.BackField
}

func (l *list) PushFront(v interface{}) *ListItem {
	frontOld := l.Front()
	frontNew := &ListItem{
		Value: v,
		Next:  frontOld,
		Prev:  nil,
	}

	if frontOld != nil {
		frontOld.Prev = frontNew
	}

	l.FrontField = frontNew
	if l.Len() == 0 {
		l.BackField = frontNew
	}
	l.LenField++

	return frontNew
}

func (l *list) PushBack(v interface{}) *ListItem {
	backOld := l.Back()
	backNew := &ListItem{
		Value: v,
		Next:  nil,
		Prev:  backOld,
	}

	if backOld != nil {
		backOld.Next = backNew
	}

	l.BackField = backNew
	if l.Len() == 0 {
		l.FrontField = backNew
	}
	l.LenField++

	return backNew
}

func (l *list) Remove(i *ListItem) {
	if l.Len() == 1 {
		l.FrontField = nil
		l.BackField = nil
		l.LenField--
		return
	}

	prev, next := i.Prev, i.Next
	if prev == nil {
		next.Prev = prev
		l.FrontField = next
		l.LenField--
		return
	}
	if next == nil {
		prev.Next = next
		l.BackField = prev
		l.LenField--
		return
	}

	prev.Next = next
	next.Prev = prev
	l.LenField--
}

func (l *list) MoveToFront(i *ListItem) {
	value := i.Value
	l.Remove(i)
	l.PushFront(value)
}
