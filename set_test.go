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
	old := Set.Add("a", "b")
	s1 := old.Keys()
	s1 = s1.CharSort(false)
	if !reflect.DeepEqual(s1, gostr.NewSlice([]string{"a", "b"})) {
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
	s1 = s1.CharSort(false)
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
	s1 := Set.ToSlice()
	s1 = s1.CharSort(false)
	if !reflect.DeepEqual(s1, gostr.Slice{"a", "b"}) {
		panic("test Set.ToSlice Error")
	}
}

func TestSet_SetIntersection(t *testing.T) {
	s1 := Set.Intersection(gostr.NewSet([]string{"a", "h"}))
	if !reflect.DeepEqual(s1, gostr.NewSet([]string{"a"})) {
		panic("test Slice.SetIntersection Error")
	}
}

func TestSet_Union(t *testing.T) {
	s1 := Set.Union(gostr.NewSet([]string{"a", "h"}))
	if !reflect.DeepEqual(s1, gostr.NewSet([]string{"a", "b", "h"})) {
		panic("test Slice.SetIntersection Error")
	}
}
