package main

type Node[T comparable] struct {
	next  *Node[T]
	value T
}

type Queue[T comparable] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func (q *Queue[T]) enqueue(item T) {
	newNode := &Node[T]{value: item}
	if q.length == 0 {
		q.head = newNode
		q.tail = newNode
		q.length++
		return
	}
	q.length++
	q.tail.next = newNode
	q.tail = newNode
}

func (q *Queue[T]) dequeue() (item T, ok bool) {
	if q.length == 0 {
		return
	}
	h := q.head
	q.head = q.head.next
	h.next = nil
	q.length--
	return h.value, true
}

func (q *Queue[T]) peek() (item T, ok bool) {
	if q.head == nil {
		return
	}
	return q.head.value, true
}
