package pizza

import "reflect"

type Slice[T any] []T

// Length returns the length of the slice
func (s Slice[T]) Length() int {
	return len(s)
}

// Capacity returns the length of the slice
func (s Slice[T]) Capacity() int {
	return cap(s)
}

// Equals checks whether two Slices are deeply equal or not.
func (a Slice[T]) Equals(b Slice[T]) bool {
	return reflect.DeepEqual(a, b)
}

// Pop removes the last element from the slice and returns it.
// If the slice is empty, it returns nil and an error.
func (s *Slice[T]) Pop() (*T, error) {
	if len(*s) > 0 {
		popped := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]
		return &popped, nil
	}
	return nil, ErrEmptySlice
}

// Push appends an element to the slice and returns the updated slice.
func (s *Slice[T]) Push(v T) *Slice[T] {
	*s = append(*s, v)
	return s
}

// ForEach loops through a Slice and calls it's callback func during each iteration.
func (s Slice[T]) ForEach(callback func(element T, index int)) {
	for i, v := range s {
		callback(v, i)
	}
}

// Some loops through a Slice and returns a boolean based on whether a predicate is met.
func (s Slice[T]) Some(callback func(element T, index int) bool) bool {
	for i, v := range s {
		if callback(v, i) {
			return true
		}
	}
	return false
}