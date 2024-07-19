package pizza

import "testing"

type PopTestCase[T any] struct {
	name           string
	preValues      Slice[T]
	postValues     Slice[T]
	expectedPopped *T
	expectedError  error
}

func toPtr[T any](v T) *T {
	return &v
}

func TestPop(t *testing.T) {
	testCases := []PopTestCase[int]{
		{
			name:           "Pop from non-empty slice",
			preValues:      Slice[int]{1, 2, 3},
			postValues:     Slice[int]{1, 2},
			expectedPopped: toPtr(3),
			expectedError:  nil,
		},
		{
			name:           "Pop from one-element slice",
			preValues:      Slice[int]{42},
			postValues:     Slice[int]{},
			expectedPopped: toPtr(42),
			expectedError:  nil,
		},
		{
			name:           "Pop from empty slice",
			preValues:      Slice[int]{},
			postValues:     Slice[int]{},
			expectedPopped: nil,
			expectedError:  ErrEmptySlice,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			popped, err := tc.preValues.Pop()

			// Check popped
			if popped != nil && (tc.expectedPopped == nil || *popped != *tc.expectedPopped) {
				t.Errorf("Test case '%s': Expected popped value %v, got %v", tc.name, tc.expectedPopped, *popped)
			} else if popped == nil && tc.expectedPopped != nil {
				t.Errorf("Test case '%s': Expected popped value %v, got nil", tc.name, tc.expectedPopped)
			}

			// Check postValues
			pre, post := tc.preValues, tc.postValues
			if !pre.Equals(post) {
				t.Errorf("Test case '%s': Post and pre values mismatch. Expected %v, got %v", tc.name, post, pre)
			}

			// Check error
			if err == nil && tc.expectedError != nil {
				t.Errorf("Test case '%s': Expected error '%v', got nil", tc.name, tc.expectedError)
			} else if err != nil && tc.expectedError == nil {
				t.Errorf("Test case '%s': Expected no error, got '%v'", tc.name, err)
			} else if err != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("Test case '%s': Expected error '%v', got '%v'", tc.name, tc.expectedError, err)
			}
		})
	}
}

type ForEachTestCase[T any] struct {
	name     string
	input    Slice[T]
	expected []T
	indices  []T
}

func TestForEach(t *testing.T) {
	testCases := []ForEachTestCase[int]{
		{
			name:     "ForEach over a non-empty slice",
			input:    Slice[int]{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
			indices:  []int{0, 1, 2, 3, 4},
		},
		{
			name:     "ForEach over an empty slice",
			input:    Slice[int]{},
			expected: []int{},
			indices:  []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var values []int
			var indices []int
			tc.input.ForEach(func(e int, i int) {
				values = append(values, e)
				indices = append(indices, i)
			})

			if len(values) != len(tc.expected) {
				t.Fatalf("Test case '%s': Expected %d number of values, got %d", tc.name, len(tc.expected), len(values))
			}

			for i := range values {
				if values[i] != tc.expected[i] {
					t.Errorf("Test case '%s': Expected value %d at index %d, got %d", tc.name, tc.expected[i], i, values[i])
				}
			}

			for i := range indices {
				if indices[i] != tc.indices[i] {
					t.Errorf("Test case '%s': Expected index %d, got %d", tc.name, tc.indices[i], indices[i])
				}
			}
		})
	}
}
