// @Time  : 2022/8/22 17:56
// @Email: jtyoui@qq.com

package gostr

type Slice []Str

// NewSlice create a slice of string
func NewSlice[T ~[]E, E string](t T) Slice {
	var values = make(Slice, len(t))
	for index, value := range t {
		values[index] = New(value)
	}
	return values
}

// ToSet convert strings slice to set
func (s Slice) ToSet() map[Str]struct{} {
	m := make(map[Str]struct{}, len(s))
	for _, value := range s {
		if _, ok := m[value]; !ok {
			m[value] = struct{}{}
		}
	}
	return m
}

// Deduplication Remove duplicate data from string slice and return
func (s Slice) Deduplication() Slice {
	var values Slice

	for key := range s.ToSet() {
		values = append(values, key)
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
	cr := make(Slice, len(s))
	for j, i := 0, len(s)-1; i >= 0; i, j = i-1, j+1 {
		cr[j] = cr[i]
	}
	return cr
}

// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func (s Slice) Join(sep Str) Str {
	return sep.Join(s.ToString())
}
