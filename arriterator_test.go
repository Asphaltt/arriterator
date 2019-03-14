package arriterator

import (
	"testing"
)

func TestNew(t *testing.T) {
	i, e := New(nil)
	if e == nil {
		t.Log("failed with nil pointer")
		t.Fail()
	}

	i, e = New([]int{})
	if e != nil {
		t.Log("failed with integer array")
		t.Fail()
	}

	i, e = New(([]int{1, 2, 3})[:0])
	if e != nil {
		t.Log("failed with integer slice")
		t.Fail()
	}

	i, e = New(&[]int{})
	if e == nil {
		t.Log("failed with integer array pointer")
		t.Fail()
	}

	i, e = New(map[int]int{})
	if e == nil {
		t.Log("failed with integer-integer map")
		t.Fail()
	}

	_ = i
}

func TestHasNext(t *testing.T) {
	var i Iterator
	if i != nil && i.HasNext() {
		t.Log("failed with nil Iterator")
		t.Fail()
	}

	i, _ = New(&[]int{})
	if i != nil && i.HasNext() {
		t.Log("failed with integer array pointer")
		t.Fail()
	}

	i, _ = New([]int{})
	if i.HasNext() {
		t.Log("failed with empty integer array")
		t.Fail()
	}

	i, _ = New([]int{1})
	if !i.HasNext() {
		t.Log("failed with non-empty integer array")
		t.Fail()
	}

	i, _ = New([]int{1, 2, 3})
	if v := i.Next(); v == nil {
		t.Log("failed with iterated non-empty integer array")
		t.Fail()
	}
	if v := i.Next(); v == nil {
		t.Log("failed with iterated non-empty integer array")
		t.Fail()
	}
	if v := i.Next(); v == nil {
		t.Log("failed with iterated non-empty integer array")
		t.Fail()
	}
	if i.HasNext() {
		t.Log("failed with iterated non-empty integer array")
		t.Fail()
	}
}

func TestNext(t *testing.T) {
	var i Iterator
	if i != nil && i.Next() != nil {
		t.Log("failed with nil Iterator")
		t.Fail()
	}

	i, _ = New([]int{})
	if i == nil || i.Next() != nil {
		t.Log("failed with empty integer array")
		t.Fail()
	}

	i, _ = New([]int{1})
	if i == nil || i.Next() == nil {
		t.Log("failed with non-empty integer array")
		t.Fail()
	}

	i, _ = New([]int{1, 2, 3})
	if v := i.Next(); v == nil {
		t.Log("failed with iterated non-empty integer array")
		t.Fail()
	}
	if v := i.Next(); v == nil {
		t.Log("failed with iterated non-empty integer array")
		t.Fail()
	}
	if v := i.Next(); v == nil {
		t.Log("failed with iterated non-empty integer array")
		t.Fail()
	}
	if v := i.Next(); v != nil {
		t.Log("failed with iterated non-empty integer array")
		t.Fail()
	}
}

func TestIterator(t *testing.T) {
	ints := []int{111, 222, 333}
	i, e := New(ints)
	if e != nil {
		t.Log("failed with non-empty integer array")
		t.FailNow()
	}
	idx := 0
	for i.HasNext() {
		v, ok := i.Next().(int)
		if !ok {
			t.Log("failed with iterating non-empty integer array")
			t.FailNow()
		}

		if v != ints[idx] {
			t.Log("failed with iterating non-empty integer array")
			t.FailNow()
		}
		idx++
	}

	strings := []string{"abc", "bcd", "xxx"}
	i, e = New(strings)
	if e != nil {
		t.Log("failed with non-empty string array")
		t.FailNow()
	}
	idx = 0
	for i.HasNext() {
		v, ok := i.Next().(string)
		if !ok {
			t.Log("failed with iterating non-empty string array")
			t.FailNow()
		}
		if v != strings[idx] {
			t.Log("failed with iterating non-empty string array")
			t.FailNow()
		}
		idx++
	}
}
