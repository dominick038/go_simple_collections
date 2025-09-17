package stack

import "math"

type arrayStack[T primitive] struct {
	arr      []T
	capacity uint
	count    uint
}

func NewArrayStack[T primitive]() Stack[T] {
	arr := make([]T, 32)
	return &arrayStack[T]{
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

func (as *arrayStack[T]) Push(value T) {
	// This is bad and does not work at all, needs fixing but lazy umu
	// Should check if size == maxuint and also should check if capacity will overflow and then set to maxuint
	if as.capacity == math.MaxUint {
		panic("Stack exceeded capacity!!")
	}

	if as.capacity == as.count {
		newCap := as.capacity * 2
		as.arr = resizeArr(as.arr, newCap)
		as.capacity = newCap
	}

	as.arr[as.count] = value
	as.count++
}

func (as *arrayStack[T]) Pop() (value T, err error) {
	if as.count == 0 {
		return value, StackEmptyError{}
	}

	// Could do with a downsize check on pop?
	// Doesn't matter too much but could be usefull, would be make and then copy to index

	value = as.arr[as.count-1]
	as.count--

	return value, nil
}

func (as *arrayStack[T]) Peek() (value T, err error) {
	if as.count == 0 {
		return value, StackEmptyError{}
	}

	return as.arr[as.count-1], nil
}

func (as *arrayStack[T]) Count() (count uint) {
	return as.count
}

func (lls *arrayStack[T]) IsEmpty() bool {
	return lls.count == 0
}
