// @Time  : 2022/8/22 17:56
// @Email: jtyoui@qq.com

package gostr

import (
	"errors"
	"sort"
)

type Slice []Str

// NewSlice create a slice of Str
// NewSlice create a slice of string
func NewSlice[T ~[]E, E Stringer](t T) Slice {
	var values = make(Slice, len(t))
	for index, value := range t {
		values[index] = New(value)
	}
	return values
}

// ToSet convert strings slice to set
func (s Slice) ToSet() Set {
	return NewSet(s)
}

// Deduplication Remove duplicate data from string slice and return
// Notice: order
func (s Slice) Deduplication() Slice {
	var values Slice

	flag := make(map[Str]struct{}, s.Len())

	for _, key := range s {
		if _, ok := flag[key]; !ok {
			values = append(values, key)
			flag[key] = struct{}{}
		}
	}

	return values
}

// In Check if a string is in an array
func (s Slice) In(sub Str) bool {
	for i := 0; i < len(s); i++ {
		if s[i] == sub {
			return true
		}
	}
	return false
}

// ToString convert String type
func (s Slice) ToString() []string {
	var values = make([]string, len(s))

	for index, value := range s {
		values[index] = value.ToString()
	}
	return values
}

// Reverse copy Any slice and positions reverse
func (s Slice) Reverse() Slice {
	s1 := s.Clone()
	length := s1.Len()
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		flag := s1[i]
		s1[i] = s1[j]
		s1[j] = flag
	}
	return s1
}

// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func (s Slice) Join(sep Str) Str {
	return sep.Join(s.ToString())
}

// Append at the Slice tail add some elem.
func (s Slice) Append(elems ...Str) Slice {
	return append(s, elems...)
}

// Len get slice length
//
// Returns 0 if s==nil or len(s)==0
func (s Slice) Len() int {
	return len(s)
}

// Clone get a Clone Slice data
func (s Slice) Clone() Slice {
	return append(Slice(nil), s...)
}

// CharSort character sorting,default ascending
func (s Slice) CharSort(descending bool) Slice {
	c := s.Clone()
	if descending {
		sort.Slice(c, func(i, j int) bool {
			return c[i] > c[j]
		})
	} else {
		sort.Slice(c, func(i, j int) bool {
			return c[i] < c[j]
		})
	}
	return c
}

// Delete remove an element in Slice.
//
// Returns err if the elements not exits in Slice or Slice.Len == 0.
func (s Slice) Delete(ele Str) (ns Slice, err error) {
	if s.Len() == 0 {
		err = errors.New("slice len is null")
		return
	}

	for _, key := range s {
		if key != ele {
			ns = ns.Append(key)
		}
	}
	return
}

// Eq compares two Slice to be same.
// Notice: ignoring String case from Slice.
func (s Slice) Eq(s1 Slice) bool {
	if s.Len() != s1.Len() {
		return false
	}

	for i := 0; i < s.Len(); i++ {
		if !s[i].Eq(s1[i]) {
			return false
		}
	}

	return true
}

// Counter counts the number of elements
func (s Slice) Counter() map[Str]int {
	var m = make(map[Str]int, len(s))
	for _, key := range s {
		m[key] += 1
	}
	return m
}

// Index similar python string index
func (s Slice) Index(start int, end int) Slice {
	length := s.Len()

	// return Str.Len if end == 0
	if end == 0 {
		end = length
	} else if end < 0 {
		end += length
	}

	if start < 0 {
		start += length
	}

	if start > end || end > length {
		panic("end index must > start index,and < length")
	}

	if start < 0 || end < 0 {
		panic("index format error")
	}

	value := s[start:end]
	return value
}

// POP get the last element
func (s Slice) POP() (ns Slice, ele Str) {
	return s.Index(0, -1), s[s.Len()-1]
}
