package main

import "testing"

func TestPush(t *testing.T) {
	stack := newStackExample()
	assert(t, stack.length, 10)
	assert(t, stack.head.value, 9)
	assert(t, stack.head.next.value, 8)

	stack.push(10)
	assert(t, stack.length, 11)
	assert(t, stack.head.value, 10)
}

func TestPop(t *testing.T) {
	t.Run("Empty stack", func(t *testing.T) {
		stack := &Stack[int]{}
		_, ok := stack.pop()
		assert(t, ok, false)
	})
	t.Run("One node stack", func(t *testing.T) {
		stack := &Stack[int]{}
		stack.push(1)
		item, ok := stack.pop()
		assert(t, ok, true)
		assert(t, item, 1)
		assert(t, stack.length, 0)
		assert(t, stack.head, nil)
	})
	t.Run("Full Stack", func(t *testing.T) {
		stack := newStackExample()
		headAfterPop := stack.head.next
		item, ok := stack.pop()
		assert(t, ok, true)
		assert(t, item, 9)
		assert(t, stack.length, 9)
		assert(t, stack.head, headAfterPop)
	})
}

func TestPeek(t *testing.T) {
	t.Run("Empty stack", func(t *testing.T) {
		stack := &Stack[int]{}
		_, ok := stack.peek()
		assert(t, ok, false)
	})
	t.Run("Full Stack", func(t *testing.T) {
		stack := newStackExample()
		item, ok := stack.peek()
		assert(t, ok, true)
		assert(t, item, 9)
	})
}

func assert[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func newStackExample() *Stack[int] {
	stack := &Stack[int]{}
	for i := 0; i < 10; i++ {
		stack.push(i)
	}
	return stack
}
