package queue

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

func resizeArr[T primitive](arr []T, newCap uint) (newArr []T) {
	result := make([]T, newCap)
	copy(result, arr)
	return result
}

func (llq *arrayQueue[T]) EnQueue(value T) {
	if llq.capacity == llq.count {
		newCap := llq.capacity * 2
		llq.arr = resizeArr(llq.arr, newCap)
		llq.capacity = newCap
	}

	// Babababooy this will explode the array bcs rly rly large like the other array implementation UwU
	// But then again if your queue is anywhere close to this: 18446744073709551615 you have other problems

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
