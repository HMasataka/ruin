package ruin

func New[T any](data []T) *Ruin[T] {
	return &Ruin[T]{data: data}
}

type Ruin[T any] struct {
	data []T
}

func (r *Ruin[T]) Data() []T {
	return r.data
}

func (r *Ruin[_]) Len() int {
	return len(r.data)
}
