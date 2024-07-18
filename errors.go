package pizza

import "errors"

var ErrEmptySlice = errors.New("cannot Pop from an empty slice")