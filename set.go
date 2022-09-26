// @Time  : 2022/9/5 13:28
// @Email: jtyoui@qq.com

package gostr

type Set map[Str]struct{}

// NewSet new Set object of Str
//
// Notice: Set is disorder!
func NewSet[T ~[]E, E Stringer](s T) Set {
	m := make(Set, len(s))
	for _, e := range s {
		m[New(e)] = struct{}{}
	}
	return m
}

// Add some elements to the Set
func (s Set) Add(elems ...Str) Set {
	var s1 Set
	if s == nil {
		s1 = make(Set)
	} else {
		s1 = s.Clone()
	}

	for _, e := range elems {
		if !s1.In(e) {
			s1[e] = struct{}{}
		}
	}
	return s1
}

// ToSlice convert Slice type
func (s Set) ToSlice() Slice {
	return NewSlice(s.Keys())
}

// Keys get all elements
//
// Notice: this is disorder!
func (s Set) Keys() Slice {
	var key Slice
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

// Intersection  Get the intersection of s and s1
func (s Set) Intersection(s1 Set) Set {
	var s2 = make(Set, 0)
	for key := range s {
		if _, ok := s1[key]; ok {
			s2[key] = struct{}{}
		}
	}
	return s2
}

// Union Get the union of s and s1
func (s Set) Union(s1 Set) Set {
	s2 := make(Set, s.Len()+s1.Len())

	for key := range s {
		if _, ok := s2[key]; !ok {
			s2[key] = struct{}{}
		}
	}

	for key := range s1 {
		if _, ok := s2[key]; !ok {
			s2[key] = struct{}{}
		}
	}

	return s2
}

// Clone a set of s
func (s Set) Clone() Set {
	s1 := make(Set, s.Len())
	for key := range s {
		s1[key] = struct{}{}
	}
	return s1
}
