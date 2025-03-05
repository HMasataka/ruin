package ruin

import "errors"

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

func (r *Ruin[T]) Pop() (T, error) {
	var result T

	if len(r.data) == 0 {
		return result, errors.New("ruin: empty data")
	}

	result = r.data[0]
	r.data = r.data[1:]

	return result, nil
}

func (r *Ruin[T]) Insert(data T) {
	r.data = append(r.data, data)
}
