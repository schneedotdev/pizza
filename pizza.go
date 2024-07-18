package pizza

import "reflect"

type Slice[T any] []T

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
