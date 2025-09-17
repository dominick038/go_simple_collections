package queue

type node[T primitive] struct {
	value T
	next  *node[T]
}

type linkedListQueue[T primitive] struct {
	front *node[T]
	back  *node[T]
	count uint
}

func NewLinkedListQueue[T primitive]() Queue[T] {
	return &linkedListQueue[T]{
		front: nil,
		back:  nil,
		count: 0,
	}
}

func (llq *linkedListQueue[T]) EnQueue(value T) {
	newNode := &node[T]{
		value: value,
		next:  nil,
	}

	if llq.count == 0 {
		llq.front = newNode
		llq.back = newNode
	} else {
		llq.back.next = newNode
		llq.back = newNode
	}

	llq.count++
}

func (llq *linkedListQueue[T]) DeQueue() (value T, err error) {
	if llq.count == 0 {
		return value, &QueueEmptyError{}
	}

	result := llq.front
	llq.front = result.next
	llq.count--

	return result.value, nil
}

func (llq *linkedListQueue[T]) Count() uint {
	return llq.count
}
