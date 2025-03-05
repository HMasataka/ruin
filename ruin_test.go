package ruin_test

import (
	"testing"

	"github.com/HMasataka/ruin"
	"github.com/stretchr/testify/assert"
)

func TestNewRuin(t *testing.T) {
	dummy := []int{1, 2, 3, 4, 5}

	r := ruin.New(dummy)
	if r == nil {
		assert.Fail(t, "New() is nil")
	}
}

func TestRuinData(t *testing.T) {
	dummy := []int{1, 2, 3, 4, 5}

	r := ruin.New(dummy)
	assert.Equal(t, []int{1, 2, 3, 4, 5}, r.Data())
}

func TestRuinLen(t *testing.T) {
	dummy := []int{1, 2, 3, 4, 5}

	r := ruin.New(dummy)
	assert.Equal(t, 5, r.Len())
}

func TestRuinPop(t *testing.T) {
	dummy := []int{1, 2, 3, 4, 5}

	r := ruin.New(dummy)

	value, err := r.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 1, value)
}
