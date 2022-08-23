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
