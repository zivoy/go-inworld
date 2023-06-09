package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyQueue(t *testing.T) {
	queue := NewQueue[int]()
	assert.NotNil(t, queue)

	assert.Equal(t, 0, queue.Size(), "queue should be empty")
}

func TestFillQueue(t *testing.T) {
	queue := NewQueue[int]()

	assert.Equal(t, 0, queue.Size(), "queue should be empty")
	queue.Append(NewPointer(4))
	queue.Append(NewPointer(2))
	queue.Append(NewPointer(7))
	assert.Equal(t, 3, queue.Size(), "queue should be empty")
}

func TestQueue_Pop(t *testing.T) {
	queue := NewQueue[int]()

	queue.Append(NewPointer(4))
	queue.Append(NewPointer(2))
	assert.Equal(t, 2, queue.Size(), "queue should have items")

	assert.Equal(t, 4, *queue.Pop(), "did not take out first argument")
	assert.Equal(t, 2, *queue.Pop())
	assert.Nil(t, queue.Pop(), "queue shouldnt have anything in it")
}

func TestQueue_Clear(t *testing.T) {
	queue := NewQueue[string]()

	queue.Append(NewPointer("something"))
	queue.Append(NewPointer("or other"))
	assert.Equal(t, 2, queue.Size(), "queue should have items")

	assert.Equal(t, "something", *queue.Pop(), "did not take out first argument")

	queue.Clear()
	assert.Equal(t, 0, queue.Size(), "queue should be empty")
}
