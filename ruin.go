package ruin

func New[T any](data []T) *Ruin[T] {
	return &Ruin[T]{Data: data}
}

type Ruin[T any] struct {
	Data []T
}

func (r *Ruin[_]) Len() int {
	return len(r.Data)
}
