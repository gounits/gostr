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

// SetIntersection  Get the intersection of s and s1
func (s Slice) SetIntersection(s1 Slice) Slice {
	s1Set := s.ToSet()

	s2Set := s1.ToSet()

	var s2 Slice
	for key := range s1Set {
		if _, ok := s2Set[key]; ok {
			s2 = append(s2, key)
		}
	}
	return s2
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

// Sort sorts the Slice given the rule function.
//
// the rule function must satisfy the same requirements at the interface
// type's Rule function.
func (s Slice) Sort(rule func(pre, next int) bool) Slice {
	c := s.Clone()
	sort.Slice(c, rule)
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
