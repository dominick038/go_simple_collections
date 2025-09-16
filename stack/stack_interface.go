package stack

type primitive interface {
	~bool |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
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
