package main

import (
	"testing"
)

func TestGet(t *testing.T) {
	linkedlist := newLinkedListExample()
	got := linkedlist.get(5)
	want := 5
	assert(t, got, want)

	got = linkedlist.get(0)
	want = linkedlist.head.value
	assert(t, got, want)

	got = linkedlist.get(linkedlist.length - 1)
	want = linkedlist.tail.value
	assert(t, got, want)
}

func TestRemove(t *testing.T) {
	linkedlist := newLinkedListExample()
	length := linkedlist.length
	got := linkedlist.remove(4)
	want := 4
	assert(t, got, want)
	assert(t, linkedlist.length, length-1)

	got = linkedlist.get(4)
	want = 5
	assert(t, got, want)
}

func TestAppend(t *testing.T) {
	linkedlist := newLinkedListExample()
	length := linkedlist.length
	linkedlist.append(10)
	assert(t, linkedlist.length, length+1)

	got := linkedlist.tail.value
	want := 10
	assert(t, got, want)

	got = linkedlist.get(10)
	want = 10
	assert(t, got, want)
}

func TestPrepend(t *testing.T) {
	linkedlist := newLinkedListExample()
	length := linkedlist.length
	linkedlist.prepend(-1)
	assert(t, linkedlist.length, length+1)

	got := linkedlist.head.value
	want := -1
	assert(t, got, want)

	got = linkedlist.get(0)
	want = -1
	assert(t, got, want)
}

func TestInsertAt(t *testing.T) {
	linkedlist := newLinkedListExample()
	length := linkedlist.length
	linkedlist.insertAt(0, 3)

	got := linkedlist.get(3)
	want := 0
	assert(t, got, want)
	assert(t, linkedlist.length, length+1)

	linkedlist.insertAt(-1, 0)
	got = linkedlist.get(0)
	want = -1
	assert(t, got, want)
	assert(t, got, linkedlist.head.value)

	linkedlist.insertAt(10, linkedlist.length)
	got = linkedlist.get(linkedlist.length - 1)
	want = 10
	assert(t, got, want)
	assert(t, got, linkedlist.tail.value)
}

func assert[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v; want %v", got, want)
	}
}

func newLinkedListExample() *LinkedList[int] {
	linkedlist := &LinkedList[int]{}
	for i := 0; i < 10; i++ {
		linkedlist.append(i)
	}
	return linkedlist
}
