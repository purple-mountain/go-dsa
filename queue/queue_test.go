package main

import (
	"testing"
)

func TestQueue(t *testing.T) {
	t.Run("enqueue", func(t *testing.T) {
		queue := newQueueExample()
		assert(t, queue.length, 10)
		assert(t, queue.head.value, 0)
		assert(t, queue.tail.value, 9)
	})
	t.Run("dequeue", func(t *testing.T) {
		queue := newQueueExample()
		lengthBefore := queue.length
		headAfterDequeue := queue.head.next
		item, ok := queue.dequeue()
		assert(t, item, 0)
		assert(t, ok, true)
		assert(t, queue.length, lengthBefore-1)
		assert(t, queue.head, headAfterDequeue)
		assert(t, queue.head.value, headAfterDequeue.value)

		emptyQueue := &Queue[int]{}
		_, ok = emptyQueue.dequeue()
		assert(t, ok, false)
	})
	t.Run("peek", func(t *testing.T) {
		queue := newQueueExample()
		item, ok := queue.peek()
		assert(t, item, 0)
		assert(t, ok, true)

		emptyQueue := &Queue[int]{}
		_, ok = emptyQueue.peek()
		assert(t, ok, false)
	})
}

func assert[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func newQueueExample() *Queue[int] {
	queue := &Queue[int]{}
	for i := 0; i < 10; i++ {
		queue.enqueue(i)
	}
	return queue
}
