package stack

type ArrayStack[T primitive] struct {
	arr      []T
	capacity uint
	count    uint
}

func NewArrayStack[T primitive]() Stack[T] {
	arr := make([]T, 32)
	return &ArrayStack[T]{
		arr:      arr,
		capacity: 32,
		count:    0,
	}
}

func resizeArr[T primitive](arr []T, newCap uint) (newArr []T) {
	result := make([]T, newCap)
	copy(result, arr)
	return result
}

func (as *ArrayStack[T]) Push(value T) {
	if as.capacity == as.count {
		newCap := as.capacity * 2
		as.arr = resizeArr(as.arr, newCap)
		as.capacity = newCap
	}

	as.arr[as.count] = value
	as.count++
}

func (as *ArrayStack[T]) Pop() (value T, err error) {
	if as.count == 0 {
		return value, StackEmptyError{}
	}

	value = as.arr[as.count-1]
	as.count--

	return value, nil
}

func (as *ArrayStack[T]) Peek() (value T, err error) {
	if as.count == 0 {
		return value, StackEmptyError{}
	}

	return as.arr[as.count-1], nil
}

func (as *ArrayStack[T]) Count() (count int) {
	return int(as.count)
}
