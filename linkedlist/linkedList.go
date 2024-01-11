package main

import "fmt"

type Node[T comparable] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

type LinkedList[T comparable] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func (l *LinkedList[T]) get(idx int) (item T) {
	if idx > l.length-1 {
		return
	}
	if l.length == 0 {
		return
	}
	node := l.head
	for i := 0; i < idx; i++ {
		node = node.next
	}
	return node.value
}

func (l *LinkedList[T]) append(value T) {
	newNode := &Node[T]{value: value}
	if l.length == 0 {
		l.head = newNode
		l.tail = newNode
		l.length++
		return
	}
	tmp := l.tail
	tmp.next = newNode
	newNode.prev = tmp
	l.tail = newNode
	l.length++
}

func (l *LinkedList[T]) prepend(value T) {
	newNode := &Node[T]{value: value}
	if l.length == 0 {
		l.head = newNode
		l.tail = newNode
		l.length++
		return
	}
	temp := l.head
	temp.prev = newNode
	newNode.next = temp
	l.head = newNode
	l.length++
}

func (l *LinkedList[T]) printAllValues() {
	node := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%v ", node.value)
		node = node.next
	}
	fmt.Printf("\n")
}

func (l *LinkedList[T]) remove(item T) (removedItem T) {
	if l.length == 0 {
		return
	}
	currentNode := l.head
	if l.head.value == item && l.length == 1 {
		l.head = nil
		l.tail = nil
		l.length--
		return
	}
	if l.head.value == item {
		l.head = currentNode.next
		currentNode.next = nil
		l.length--
		return currentNode.value
	}
	if l.tail.value == item {
		currentNode = l.tail
		l.tail = currentNode.prev
		currentNode.prev = nil
		l.length--
		return currentNode.value
	}
	for currentNode.next != nil {
		if currentNode.value == item {
			currentNode.prev.next = currentNode.next
			currentNode.next.prev = currentNode.prev
			currentNode.next = nil
			currentNode.prev = nil
			l.length--
			return currentNode.value
		}
		currentNode = currentNode.next
	}
	return
}

func (l *LinkedList[T]) insertAt(item T, idx int) {
	if idx > l.length {
		return
	}
	newNode := &Node[T]{value: item}
	if l.length == 0 {
		l.head = newNode
		l.tail = newNode
		l.length++
		return
	}

	if idx == 0 {
		l.prepend(item)
		return
	} else if idx == l.length {
		l.append(item)
		return
	}

	node := l.head
	for i := 0; i < idx; i++ {
		node = node.next
	}
	prevNode := node.prev
	prevNode.next = newNode
	node.prev = newNode
	newNode.prev = prevNode
	newNode.next = node
	l.length++
}

// Panics
// func (l *LinkedList[T]) removeAll() {
// 	node := l.head
// 	for i := 0; i < l.length; i++ {
// 		l.remove(node.value)
// 		node = node.next
// 	}
// }
