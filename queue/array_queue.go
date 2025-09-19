package queue

import "math"

type arrayQueue[T primitive] struct {
	arr      []T
	front    uint
	back     uint
	capacity uint
	count    uint
}

func NewArrayQueue[T primitive]() Queue[T] {
	arr := make([]T, 32)
	return &arrayQueue[T]{
		arr:      arr,
		front:    0,
		back:     0,
		count:    0,
		capacity: 32,
	}
}

func (llq *arrayQueue[T]) increaseCapacity() {
	// Check if we overflow the capacity for some weird reason
	newCap := llq.capacity * 2
	if newCap < llq.capacity {
		newCap = math.MaxUint
	}
	newArr := make([]T, newCap)

	// When full, always copy in two parts since buffer is wrapped
	firstPartSize := llq.capacity - llq.front
	copy(newArr, llq.arr[llq.front:])                 // From front to end
	copy(newArr[firstPartSize:], llq.arr[:llq.front]) // From start to front

	// Update the queue structure
	llq.arr = newArr
	llq.front = 0
	llq.back = llq.count
	llq.capacity = newCap
}

func (llq *arrayQueue[T]) EnQueue(value T) {
	if llq.count == math.MaxUint {
		panic("Error! Queue has reached max size, this very likely means there is an issue in your code!")
	}

	if llq.capacity == llq.count {
		llq.increaseCapacity()
	}

	llq.arr[llq.back] = value
	llq.back = (llq.back + 1) % llq.capacity
	llq.count++
}

func (llq *arrayQueue[T]) DeQueue() (value T, err error) {
	if llq.count == 0 {
		return value, &QueueEmptyError{}
	}

	value = llq.arr[llq.front]
	llq.front = (llq.front + 1) % llq.capacity
	llq.count--
	return value, nil
}

func (llq *arrayQueue[T]) Count() uint {
	return llq.count
}
