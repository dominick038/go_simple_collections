package list

type primitive interface {
	~bool |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64 |
		~string
}

type List[T primitive] interface {
	Get(index uint) (value T, err error)
	Set(index uint, value T) (err error)
	Remove(index uint) (value T, err error)
	Insert(index uint, value T) (err error)
	Append(value T)
	Prepend(value T)
	Clear()
	ToSlice() (arr []T)
	FromSlice(arr []T)
	Count() uint
	IsEmpty() bool
}

type ListOutOfBoundsErr struct {
}

func (e ListOutOfBoundsErr) Error() string {
	return "Error: List out of bounds, can't index that value"
}
