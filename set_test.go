// @Time  : 2022/9/5 15:01
// @Email: jtyoui@qq.com

package gostr_test

import (
	"github.com/gounits/gostr"
	"reflect"
	"testing"
)

var Set gostr.Set

func TestSet_Add(t *testing.T) {
	old := gostr.NewSet([]string{"a", "a"})
	old.Add("a")
	if !reflect.DeepEqual(old, gostr.NewSet([]string{"a"})) {
		panic("test Set.Add Error")
	}
}

func TestSet_In(t *testing.T) {
	if !Set.In("a") {
		panic("test Set.In Error")
	}

	if Set.In("z") {
		panic("test Set.In Error")
	}
}

func TestSet_Keys(t *testing.T) {
	s1 := Set.Keys()
	if !reflect.DeepEqual(s1, gostr.Slice{"a", "b"}) {
		panic("test Set.Keys Error")
	}
}

func TestSet_Len(t *testing.T) {
	if Set.Len() != len(Set) {
		panic("test Set.Len Error")
	}
}

func TestSet_ToSlice(t *testing.T) {
	if !reflect.DeepEqual(Set.ToSlice(), gostr.Slice{"a", "b"}) {
		panic("test Set.ToSlice Error")
	}
}
