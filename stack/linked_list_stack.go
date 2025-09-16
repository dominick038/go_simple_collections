package stack

type node[T primitive] struct {
	value T
	next  *node[T]
}

type linkedListStack[T primitive] struct {
	top   *node[T]
	count uint
}

func NewLinkedListStack[T primitive]() Stack[T] {
	return &linkedListStack[T]{
		top:   nil,
		count: 0,
	}
}

func (lls *linkedListStack[T]) Push(value T) {
	newNode := &node[T]{
		value: value,
		next:  lls.top,
	}
	lls.top = newNode
	lls.count++
}

func (lls *linkedListStack[T]) Pop() (value T, err error) {
	if lls.top == nil {
		return value, StackEmptyError{}
	}

	result := lls.top
	lls.top = result.next

	lls.count--

	return result.value, nil
}

func (lls *linkedListStack[T]) Peek() (value T, err error) {
	if lls.top == nil {
		return value, StackEmptyError{}
	}
	return lls.top.value, nil
}

func (lls *linkedListStack[T]) Count() uint {
	return lls.count
}
