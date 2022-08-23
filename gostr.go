// @Time  : 2022/8/13 20:54
// @Author: jtyoui@outlook.com

package gostr

import (
	"fmt"
	"strings"
)

type Str string

func New[T string | []byte | []rune](str T) Str {
	return Str(str)
}

// ToString convert Str to string
func (str Str) ToString() string {
	return string(str)
}

// In  check if substr is in Str.
func (str Str) In(substr Str) bool {
	return strings.Contains(str.ToString(), substr.ToString())
}

// Capitalize Convert the first letter to uppercase
func (str Str) Capitalize() Str {
	if str.Len() == 0 {
		return ""
	}
	chars := []rune(str)
	if 'a' <= chars[0] && chars[0] <= 'z' {
		chars[0] -= 'a' - 'A'
		return New(string(chars))
	}
	return str
}

// StartsWith tests whether the string s begins with prefix.
func (str Str) StartsWith(prefix Str) bool {
	return strings.HasPrefix(str.ToString(), prefix.ToString())
}

// EndsWith tests whether the string s ends with suffix.
func (str Str) EndsWith(suffix Str) bool {
	return strings.HasSuffix(str.ToString(), suffix.ToString())
}

// Find returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func (str Str) Find(substr Str) int {
	return strings.Index(str.ToString(), substr.ToString())
}

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
func (str Str) Count(substr Str) int {
	return strings.Count(substr.ToString(), substr.ToString())
}

// Len return string length
func (str Str) Len() int {
	return len(str)
}

// Lower returns s with all Unicode letters mapped to their lower case.
func (str Str) Lower() Str {
	value := strings.ToLower(str.ToString())
	return New(value)
}

// Replace returns a copy of the string s with all
// non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string
// and after each UTF-8 sequence, yielding up to k+1 replacements
// for a k-rune string.
func (str Str) Replace(old, new Str) Str {
	value := strings.ReplaceAll(str.ToString(), old.ToString(), new.ToString())
	return New(value)
}

// IsALNum Returns True if the string has at least one character and
// all characters are letters or numbers, otherwise returns False
func (str Str) IsALNum() bool {
	return _AllNum.MatchString(str.ToString())
}

// IsAlpha Returns True if the string has at least one character and all characters
// are alphabets or Chinese characters, otherwise returns False
func (str Str) IsAlpha() bool {
	return _Alpha.MatchString(str.ToString())
}

// IsDigit Returns True if the string contains only numbers, otherwise returns False
func (str Str) IsDigit() bool {
	return _Digit.MatchString(str.ToString())
}

// IsLower Returns True if the string contains at least one case-sensitive character and all
// of those (case-sensitive) characters are lowercase, False otherwise
func (str Str) IsLower() bool {
	return _Lower.MatchString(str.ToString())
}

// IsNumeric Returns True if the string contains only numeric characters, otherwise returns False
func (str Str) IsNumeric() bool {
	return _Numeric.MatchString(str.ToString())
}

// IsSpace Returns True if the string contains only whitespace, otherwise returns False.
func (str Str) IsSpace() bool {
	return strings.TrimSpace(str.ToString()) == ""
}

// IsUpper Returns True if the string contains at least one case-sensitive character
// and all of those (case-sensitive) characters are uppercase, False otherwise
func (str Str) IsUpper() bool {
	return _Upper.MatchString(str.ToString())
}

// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
func (str Str) Join(elems []string) Str {
	return New(strings.Join(elems, str.ToString()))
}

// Strip returns a slice of the string s with all leading and
// trailing Unicode code points contained in cut removed.
func (str Str) Strip(cut Str) Str {
	return New(strings.Trim(str.ToString(), cut.ToString()))
}

// LStrip returns a slice of the string s with all leading
// Unicode code points contained in cut removed.
//
// To remove a prefix, use TrimPrefix instead.
func (str Str) LStrip(cut Str) Str {
	return New(strings.TrimLeft(str.ToString(), cut.ToString()))
}

// TrimPrefix returns s without the provided leading prefix string.
// If s doesn't start with prefix, s is returned unchanged.
func (str Str) TrimPrefix(prefix Str) Str {
	return New(strings.TrimPrefix(str.ToString(), prefix.ToString()))
}

// RStrip returns a slice of the string s, with all trailing
// Unicode code points contained in cut removed.
//
// To remove a suffix, use TrimSuffix instead.
func (str Str) RStrip(cut Str) Str {
	return New(strings.TrimRight(str.ToString(), cut.ToString()))
}

// TrimSuffix returns s without the provided trailing suffix string.
// If s doesn't end with suffix, s is returned unchanged.
func (str Str) TrimSuffix(suffix Str) Str {
	return New(strings.TrimSuffix(str.ToString(), suffix.ToString()))
}

// Split slices s into substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// The count determines the number of substrings to return:
//
//	n > 0: at most n substrings; the last substring will be the unSplit remainder.
//	n == 0: the result is nil (zero substrings)
//	n < 0: all substrings
//
// Edge cases for s and sep (for example, empty strings) are handled
// as described in the documentation for Split.
//
// To split around the first instance of a separator, see Cut.
func (str Str) Split(sep Str, num int) []Str {
	value := strings.SplitN(str.ToString(), sep.ToString(), num)
	var values = make([]Str, len(value))
	for i := 0; i < len(value); i++ {
		values[i] = New(value[i])
	}
	return values
}

// Upper returns s with all Unicode letters mapped to their upper case.
func (str Str) Upper() Str {
	return New(strings.ToUpper(str.ToString()))
}

// Index similar python string index
func (str Str) Index(start int, end int) Str {
	if end > 0 {
		if end < start {
			panic("end must > 0 and < start")
		} else {
			value := str.ToString()[start:end]
			return New(value)
		}
	}
	ends := len(str) - end
	if ends > 0 {
		return str.Index(start, ends)
	} else {
		panic(fmt.Sprintf("flashback index exceeds string length : %d < %d", len(str), -end))
	}
}

// Title returns a copy of the string s with all Unicode letters mapped to
// their Unicode title case.
func (str Str) Title() Str {
	return New(strings.ToTitle(str.ToString()))
}

// Runes Convert []rune in a string
func (str Str) Runes() []rune {
	return []rune(str)
}

// SwapCase Convert uppercase to lowercase and lowercase to uppercase in a string
func (str Str) SwapCase() Str {
	lower := str.Lower().Runes()
	upper := str.Upper().Runes()

	var values = make([]rune, len(str))

	for index, value := range str.Runes() {
		if lower[index] == value && upper[index] != value {
			values[index] = upper[index]
		} else {
			values[index] = lower[index]
		}
	}

	return New(string(values))
}
