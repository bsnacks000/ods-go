package ods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayStackSmoke(t *testing.T) {
	stack := NewArrayStack[int]()

	assert.Equal(t, stack.Size(), 0)
}

func TestArrayStackWorksAsExpected(t *testing.T) {
	stack := NewArrayStack[int]()

	for i := 0; i < 10; i++ {
		stack.Add(i, i*2)
	}

	assert.Equal(t, 10, stack.Size())

	check := []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 0, 0, 0, 0, 0, 0}
	assert.Equal(t, check, stack.getBackingArr())

	for i := 0; i < 5; i++ {
		stack.Remove(0) // pop the top five	 values
	}

	assert.Equal(t, 10, stack.Get(0)) // first should be 10

	stack.Add(4, 42) // add something in the middle
	assert.Equal(t, 42, stack.Get(4))

	val := stack.Remove(4)
	assert.Equal(t, 42, val)

	check = []int{10, 12, 14, 16, 18, 18, 0, 0, 0, 0}
	assert.Equal(t, check, stack.getBackingArr())
	assert.Equal(t, 5, stack.Size())

	assert.Panics(t, func() { stack.Get(42) }) // idx out of bounds
	assert.Panics(t, func() { stack.Add(42, 42) })
	assert.Panics(t, func() { stack.Remove(42) })
}
