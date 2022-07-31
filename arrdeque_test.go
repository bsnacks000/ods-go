package ods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayDequeInit(t *testing.T) {
	q := NewArrayDeque[int]()

	assert.Equal(t, q.Size(), 0)
}

func TestArrayDequeWorksAsExpected(t *testing.T) {
	q := NewArrayDeque[int]()

	for i := 0; i < 10; i++ {
		q.Add(i, i*2)
	}

	assert.Equal(t, 10, q.Size())

	for i := 0; i < 10; i++ {
		val := q.Remove(0)
		assert.Equal(t, i*2, val)
	}

	assert.Equal(t, 0, q.Size())
}
