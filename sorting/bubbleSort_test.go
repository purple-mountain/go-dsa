package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	got := bubbleSort([]int{4, 2, 5, 9, 3})
	want := []int{2, 3, 4, 5, 9}
	assert(t, got, want)
}

func assert[T comparable](t *testing.T, got, want []T) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v; want %v", got, want)
	}
}
