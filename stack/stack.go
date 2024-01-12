package main

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type Stack[T comparable] struct {
	length int
	head   *Node[T]
}

func (s *Stack[T]) push(item T) {
	newNode := &Node[T]{value: item}
	if s.length == 0 {
		s.head = newNode
		s.length++
		return
	}
	newNode.next = s.head
	s.head = newNode
	s.length++
}

func (s *Stack[T]) pop() (item T, ok bool) {
	if s.length == 0 {
		return
	}
	temp := s.head
	s.head = s.head.next
	temp.next = nil
	s.length--
	return temp.value, true
}

func (s *Stack[T]) peek() (item T, ok bool) {
	if s.length == 0 {
		return
	}
	return s.head.value, true
}
