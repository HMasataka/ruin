package ruin

import "sync"

type Ruin[T any] interface {
	// 基本操作
	Pop() (T, error)
	Push(T)
	Peek() (T, error)
	IsEmpty() bool
	Len() int

	// データ操作
	Clear()
	Clone() Ruin[T]
	Data() []T
}

func New[T any](data []T) Ruin[T] {
	return &ruin[T]{data: data}
}

type ruin[T any] struct {
	data []T
	m    sync.Mutex
}

func (r *ruin[T]) Data() []T {
	return r.data
}

func (r *ruin[_]) Len() int {
	return len(r.data)
}

func (r *ruin[T]) Pop() (T, error) {
	r.m.Lock()
	defer r.m.Unlock()

	var result T

	if len(r.data) == 0 {
		return result, ErrEmpty
	}

	result = r.data[0]
	r.data = r.data[1:]

	return result, nil
}

func (r *ruin[T]) Push(data T) {
	r.m.Lock()
	defer r.m.Unlock()

	r.data = append(r.data, data)
}

func (r *ruin[T]) Peek() (T, error) {
	var result T

	if len(r.data) == 0 {
		return result, ErrEmpty
	}

	result = r.data[0]

	return result, nil
}

func (r *ruin[T]) IsEmpty() bool {
	return r.Len() == 0
}

// Clear はデータをクリアします
func (r *ruin[T]) Clear() {
	r.m.Lock()
	defer r.m.Unlock()
	r.data = r.data[:0]
}

// Clone は現在のインスタンスのディープコピーを返します
func (r *ruin[T]) Clone() Ruin[T] {
	r.m.Lock()
	defer r.m.Unlock()
	newData := make([]T, len(r.data))
	copy(newData, r.data)
	return New(newData)
}
