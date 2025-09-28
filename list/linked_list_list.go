package list

type node[T primitive] struct {
	value T
	next  *node[T]
}

type linkedListList[T primitive] struct {
	front *node[T]
	back  *node[T]
	count uint
}

func NewLinkedListQueue[T primitive]() List[T] {
	return &linkedListList[T]{
		front: nil,
		back:  nil,
		count: 0,
	}
}

// Append implements List.
func (l *linkedListList[T]) Append(value T) {
	newBack := &node[T]{
		value: value,
		next:  nil,
	}

	if l.count == 0 {
		l.front = newBack
	} else {
		l.back.next = newBack
	}
	l.back = newBack
	l.count++
}

// Clear implements List.
func (l *linkedListList[T]) Clear() {
	l.front = nil
	l.back = nil
}

// Count implements List.
func (l *linkedListList[T]) Count() uint {
	return l.count
}

// FromSlice implements List.
func (l *linkedListList[T]) FromSlice(arr []T) {
	panic("unimplemented")
}

// Get implements List.
func (l *linkedListList[T]) Get(index uint) (value T, err error) {
	panic("unimplemented")
}

// Insert implements List.
func (l *linkedListList[T]) Insert(index uint, value T) (err error) {
	panic("unimplemented")
}

// IsEmpty implements List.
func (l *linkedListList[T]) IsEmpty() bool {
	panic("unimplemented")
}

// Prepend implements List.
func (l *linkedListList[T]) Prepend(value T) {
	panic("unimplemented")
}

// Remove implements List.
func (l *linkedListList[T]) Remove(index uint) (value T, err error) {
	panic("unimplemented")
}

// Set implements List.
func (l *linkedListList[T]) Set(index uint, value T) (err error) {
	panic("unimplemented")
}

// ToSlice implements List.
func (l *linkedListList[T]) ToSlice() (arr []T) {
	panic("unimplemented")
}
