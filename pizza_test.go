package pizza

import "testing"

type PopTestCase[T any] struct {
	name           string
	valuesBefore   Slice[T]
	valuesAfter    Slice[T]
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
			valuesBefore:   Slice[int]{1, 2, 3},
			valuesAfter:    Slice[int]{1, 2},
			expectedPopped: toPtr(3),
			expectedError:  nil,
		},
		{
			name:           "Pop from one-element slice",
			valuesBefore:   Slice[int]{42},
			valuesAfter:    Slice[int]{},
			expectedPopped: toPtr(42),
			expectedError:  nil,
		},
		{
			name:           "Pop from empty slice",
			valuesBefore:   Slice[int]{},
			valuesAfter:    Slice[int]{},
			expectedPopped: nil,
			expectedError:  ErrEmptySlice,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			popped, err := tc.valuesBefore.Pop()

			// Check popped
			if popped != nil && (tc.expectedPopped == nil || *popped != *tc.expectedPopped) {
				t.Errorf("Test case '%s': Expected popped value %v, got %v", tc.name, tc.expectedPopped, *popped)
			} else if popped == nil && tc.expectedPopped != nil {
				t.Errorf("Test case '%s': Expected popped value %v, got nil", tc.name, tc.expectedPopped)
			}

			// Check valuesAfter
			before, after := tc.valuesBefore, tc.valuesAfter
			if !before.Equals(after) {
				t.Errorf("Test case '%s': Before and after values mismatch. Expected %v, got %v", tc.name, after, before)
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
}

func TestForEach(t *testing.T) {
	testCases := []ForEachTestCase[int]{
		{
			name:     "ForEach over a non-empty slice",
			input:    Slice[int]{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "ForEach over an empty slice",
			input:    Slice[int]{},
			expected: []int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var values []int
			tc.input.ForEach(func(e int, _ int) {
				values = append(values, e)
			})

			// Check lengths of Slices
			if len(values) != len(tc.expected) {
				t.Fatalf("Test case '%s': Expected %d number of values, got %d", tc.name, len(tc.expected), len(values))
			}

			// Check values of Slices
			for i := range values {
				if values[i] != tc.expected[i] {
					t.Errorf("Test case '%s': Expected value %d at index %d, got %d", tc.name, tc.expected[i], i, values[i])
				}
			}
		})
	}
}
