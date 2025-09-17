package queue

type primitive interface {
	~bool |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}

type Queue[T primitive] interface {
	EnQueue(value T)
	DeQueue() (value T, err error)
	Count() uint
}

type QueueEmptyError struct {
}

func (e QueueEmptyError) Error() string {
	return "Error: Queue is empty, can't Dequeue()"
}
