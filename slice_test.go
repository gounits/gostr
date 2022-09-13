// @Time  : 2022/8/23 12:05
// @Email: jtyoui@qq.com
// @Author: 张伟

package gostr_test

import (
	"github.com/gounits/gostr"
	"reflect"
	"testing"
)

var Slice gostr.Slice

func TestSlice_ToSet(t *testing.T) {
	s := Slice.ToSet()
	keys := []gostr.Str{"hello", "world", "china"}
	for _, key := range keys {
		if _, ok := s[key]; !ok {
			panic("test Slice.ToSet error")
		}
	}
}

func TestSlice_Append(t *testing.T) {
	if Slice.Append("!").Join(" ") != "hello world hello china !" {
		panic("test Slice.Append error")
	}
}

func TestSlice_Clone(t *testing.T) {
	if !reflect.DeepEqual(Slice, Slice.Clone()) {
		panic("test Slice.Clone error")
	}
}

func TestSlice_Deduplication(t *testing.T) {
	if !reflect.DeepEqual(Slice.Deduplication(), gostr.Slice{"hello", "world", "china"}) {
		panic("test Slice.Deduplication error")
	}
}

func TestSlice_Delete(t *testing.T) {
	old, _ := Slice.Delete("hello")
	if !reflect.DeepEqual(gostr.Slice{"world", "china"}, old) {
		panic("test Slice.Delete error")
	}
}

func TestSlice_In(t *testing.T) {
	if !Slice.In("hello") {
		panic("test Slice.In error")
	}
}

func TestSlice_Join(t *testing.T) {
	if Slice.Join(" ") != "hello world hello china" {
		panic("test Slice.Join error")
	}
}

func TestSlice_Len(t *testing.T) {
	if Slice.Len() != len(Slice) {
		panic("test Slice.Len error")
	}
}

func TestSlice_Reverse(t *testing.T) {
	r := Slice.Reverse()
	if !reflect.DeepEqual(r, gostr.NewSlice([]string{"china", "hello", "world", "hello"})) {
		panic("test Slice.Reverse Error")
	}
	old := gostr.NewSlice([]string{"a", "b", "c", "d", "e"})
	if !reflect.DeepEqual(old.Reverse(), gostr.NewSlice([]string{"e", "d", "c", "b", "a"})) {
		panic("test Slice.Reverse error")
	}
}

func TestSlice_Sort(t *testing.T) {
	s1 := Slice.Sort(func(pre, next int) bool {
		return Slice[pre] >= Slice[next]
	})

	if !reflect.DeepEqual(s1, gostr.NewSlice([]string{"world", "hello", "hello", "china"})) {
		panic("test Slice.Sort Error")
	}
}

func TestSlice_SetIntersection(t *testing.T) {
	s1 := Slice.SetIntersection(gostr.NewSlice([]string{"hello", "you"}))
	if !reflect.DeepEqual(s1, gostr.NewSlice([]string{"hello"})) {
		panic("test Slice.SetIntersection Error")
	}
}

func TestSlice_ToString(t *testing.T) {
	if !reflect.DeepEqual(Slice.ToString(), []string{"hello", "world", "hello", "china"}) {
		panic("test Slice.ToString error")
	}
}

func TestSlice_Eq(t *testing.T) {
	if !Slice.Eq(gostr.Slice{"hello", "world", "hello", "china"}) {
		panic("test Slice.Eq error")
	}
}
