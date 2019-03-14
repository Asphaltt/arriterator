package arriterator

import (
	"errors"
	"reflect"
)

// Iterator is the interface of `arriterator`
type Iterator interface {
	// HasNext gets whether the iterator has next element
	HasNext() bool

	// Next gets the next elements of the iterator
	Next() interface{}
}

type iterator struct {
	arr reflect.Value

	idx, size int // iterating index, `arr`'s size
}

// New creates an iterator from `val`.
// If `val` is `nil` or is neither array nor slice,
// you won't get an iterator and will get an error.
//
// The created iterator converts `val` to `arr` of `[]interface{}`,
// and keeps its own index `idx` for iterating,
// and keeps the size `size` of `arr`.
func New(val interface{}) (Iterator, error) {
	if val == nil {
		return nil, errors.New("nil pointer")
	}

	v := reflect.ValueOf(val)
	if v.Kind() != reflect.Slice && v.Kind() != reflect.Array {
		return nil, errors.New("non-array or non-slice pointer")
	}

	it := &iterator{
		arr:  v,
		idx:  0,
		size: v.Len(),
	}
	return it, nil
}

// HasNext gets whether the iterator has next element of its `arr`.
// It gets `true`, if its `idx` is less than its `size`.
func (i *iterator) HasNext() bool {
	return i != nil && i.idx < i.size
}

// Next gets the next element of the iterator.
// It gets `nil`, if it's `nil` or its `idx` is not less than its `size`.
// It gets the `idx` element of `arr`, and increments `idx`.
func (i *iterator) Next() interface{} {
	if i == nil || i.idx >= i.size {
		return nil
	}

	ret := i.arr.Index(i.idx)
	i.idx++
	return ret.Interface()
}
