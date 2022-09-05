// @Time  : 2022/9/5 13:28
// @Email: jtyoui@qq.com

package gostr

type Set map[Str]struct{}

// NewSet new Set object of Str
//
// Notice: Set is disorder!
func NewSet(s ...Str) Set {
	m := make(Set, len(s))
	m.Add(s...)
	return m
}

// Add some elements to the Set
func (s Set) Add(elems ...Str) {
	for _, e := range elems {
		if !s.In(e) {
			s[e] = struct{}{}
		}
	}
}

// ToSlice convert Slice type
func (s Set) ToSlice() Slice {
	return NewSlice(s.Keys()...)
}

// Keys get all elements
//
// Notice: this is disorder!
func (s Set) Keys() []Str {
	var key []Str
	for k := range s {
		key = append(key, k)
	}
	return key
}

// Len get elements number
func (s Set) Len() int {
	return len(s)
}

// In Check if the element is in a Set?
func (s Set) In(ele Str) (ok bool) {
	_, ok = s[ele]
	return
}
