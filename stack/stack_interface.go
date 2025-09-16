package stack

type primitive interface {
	~bool | ~int | ~int64 | ~float64 | ~string | ~rune
}

type Stack[T primitive] interface {
	Push(value T)
	Pop() (value T, err error)
	Peek() (value T, err error)
	Count() int
}

type StackEmptyError struct {
}

func (e StackEmptyError) Error() string {
	return "Error: Stack is empty, can't Pop()"
}
