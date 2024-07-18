package pizza

import "testing"

type TestCase[T any] struct {
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
	testCases := []TestCase[int]{
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
