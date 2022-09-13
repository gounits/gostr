// @Time  : 2022/8/19 9:40
// @Author: jtyoui@outlook.com

package gostr

import "regexp"

var (
	_AllNum  = regexp.MustCompile("^[0-9a-zA-Z]+$")
	_Alpha   = regexp.MustCompile("^[a-zA-Z\\p{Han}]+$")
	_Lower   = regexp.MustCompile("^[a-z]+$")
	_Numeric = regexp.MustCompile("^[0-9]+$")
	_Upper   = regexp.MustCompile("^[A-Z]+$")
)

// Stringer contains all characters type
type Stringer interface {
	~string | []byte | []rune
}
