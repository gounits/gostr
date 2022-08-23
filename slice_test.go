// @Time  : 2022/8/23 12:05
// @Email: jtyoui@qq.com
// @Author: 张伟

package gostr_test

import (
	"github.com/gounits/gostr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlice_ToSet(t *testing.T) {
	values := []string{"a", "a", "b", "c"}
	set := gostr.NewSlice(values).ToSet()
	assert.Equal(t, set, map[gostr.Str]struct{}{"a": {}, "b": {}, "c": {}})
}
