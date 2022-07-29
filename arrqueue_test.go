package ods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayQueueInit(t *testing.T) {
	q := NewArrayQueue[int]()

	assert.Equal(t, q.Size(), 0)
}

func TestArrayQueueWorksAsExpected(t *testing.T) {

	q := NewArrayQueue[int]()

	for i := 0; i < 100; i++ {
		q.Add(i * 2)
	}

	assert.Equal(t, 100, q.Size())

	for i := 0; i < 100; i++ {
		val := q.Remove()
		assert.Equal(t, val, i*2)
	}

	assert.Equal(t, 0, q.Size())

}
