// @Time  : 2022/9/5 15:00
// @Email: jtyoui@qq.com

package gostr_test

import (
	"github.com/gounits/gostr"
	"math"
	"testing"
)

var test gostr.Str

func TestMain(m *testing.M) {
	test = gostr.New("hello world!")
	Slice = gostr.NewSlice([]string{"hello", "world", "hello", "china"})
	Set = gostr.NewSet([]string{"a", "a", "b"})
	m.Run()
}

func TestStr_Capitalize(t *testing.T) {
	capt := test.Capitalize()
	if capt != "Hello world!" {
		panic("test Str is not capitalized")
	}
}

func TestStr_Count(t *testing.T) {
	count := test.Count("l")
	if count != 3 {
		panic("test Str.Count error")
	}
}

func TestStr_EndsWith(t *testing.T) {
	ends := test.EndsWith("!")
	if ends != true {
		panic("test Str is not ends with !")
	}
}

func TestStr_Find(t *testing.T) {
	if test.Find("l") != 2 {
		panic("test Str.Find error")
	}

	if test.Find("z") != -1 {
		panic("test Str.Find error")
	}

}

func TestStr_In(t *testing.T) {
	if !test.In("l") {
		panic("test Str.In error")
	}

	if test.In("z") {
		panic("test Str.In error")
	}
}

func TestStr_IsALNum(t *testing.T) {
	if test.IsALNum() {
		panic("test Str.IsALNum error")
	}
	test1 := gostr.New("我是中国人！")
	if test1.IsALNum() {
		panic("test Str.IsALNum error")
	}
}

func TestStr_Index(t *testing.T) {
	oldTest := test.Index(0, -1)
	if oldTest != "hello world" {
		panic("test Str.Index error")
	}
	oldTest = test.Index(-6, 0)
	if oldTest != "world!" {
		panic("test Str.Index error")
	}
}

func TestStr_IsAlpha(t *testing.T) {
	if test.IsAlpha() {
		panic("test Str.IsAlpha error")
	}
}

func TestStr_Digit(t *testing.T) {
	if n, _ := test.Digit(); n != 0 {
		panic("test Str.Digit error")
	}
	oldTest := gostr.New("1234")
	if n, _ := oldTest.Digit(); n != 1234 {
		panic("test Str.Digit error")
	}

}

func TestStr_IsLower(t *testing.T) {
	if test.IsLower() {
		panic("test Str.IsLower error")
	}

	oldTest := gostr.New("hello")
	if !oldTest.IsLower() {
		panic("test Str.IsLower error")
	}
}

func TestStr_IsNumeric(t *testing.T) {
	if test.IsNumeric() {
		panic("test Str.IsNumeric error")
	}

	oldTest := gostr.New("2121")
	if !oldTest.IsNumeric() {
		panic("test Str.IsNumeric error")
	}
}

func TestStr_IsSpace(t *testing.T) {
	if test.IsSpace() {
		panic("test str.IsSpace error")
	}
}

func TestStr_IsUpper(t *testing.T) {
	if test.IsUpper() {
		panic("test Str.IsUpper error")
	}
}

func TestStr_Title(t *testing.T) {
	if test.Title() != "HELLO WORLD!" {
		panic("test Str.Title error")
	}
}

func TestStr_Join(t *testing.T) {
	old := gostr.New(" ")
	if old.Join([]string{"hello", "world"}) != "hello world" {
		panic("test Str.Join error")
	}
}

func TestStr_Len(t *testing.T) {
	if test.Len() != len(test) {
		panic("test Str.Len error")
	}
}

func TestStr_Levenshtein(t *testing.T) {
	f := test.Levenshtein("hello worlds")
	if math.Abs(f-0.916666666) > 0.01 {
		panic("test Str.Levenshtein error")
	}
}

func TestStr_Lower(t *testing.T) {
	if test.Lower() != "hello world!" {
		panic("test Str.Lower error")
	}
}

func TestStr_LStrip(t *testing.T) {
	if test.LStrip("hel") != "o world!" {
		panic("test Str.LStrip error")
	}
}

func TestStr_Replace(t *testing.T) {
	if test.Replace("!", ".") != "hello world." {
		panic("test Str.Replace error")
	}
}

func TestStr_RStrip(t *testing.T) {
	if test.RStrip("!") != "hello world" {
		panic("test Str.RStrip error")
	}
}

func TestStr_Runes(t *testing.T) {
	r2 := []rune(test)
	r1 := test.Runes()
	for i := 0; i < len(r1); i++ {
		if r1[i] != r2[i] {
			panic("test Str.Runes error")
		}
	}
}

func TestStr_Split(t *testing.T) {
	s1 := test.Split(" ")
	s2 := gostr.Slice{"hello", "world!"}
	for i := 0; i < 2; i++ {
		if s1[i] != s2[i] {
			panic("test Str.Split error")
		}
	}
}

func TestStr_StartsWith(t *testing.T) {
	if !test.StartsWith("he") {
		panic("test Str.StartsWith error")
	}
}

func TestStr_Strip(t *testing.T) {
	old := gostr.New(" hello")
	if old.Strip(" ") != "hello" {
		panic("test Str.Strip error")
	}
}

func TestStr_SwapCase(t *testing.T) {
	if test.SwapCase() != "HELLO WORLD!" {
		panic("test Str.SwapCase error")
	}
}

func TestStr_ToString(t *testing.T) {
	if test.ToString() != "hello world!" {
		panic("test Str.ToString error")
	}
}

func TestStr_TrimPrefix(t *testing.T) {
	if test.TrimPrefix("hel") != "lo world!" {
		panic("test Str.TrimPrefix error")
	}
}

func TestStr_TrimSuffix(t *testing.T) {
	if test.TrimSuffix("!") != "hello world" {
		panic("test Str.TrimSuffix error")
	}
}

func TestStr_Upper(t *testing.T) {
	if test.Upper() != "HELLO WORLD!" {
		panic("test Str.Upper error")
	}
}

func TestStr_Eq(t *testing.T) {
	if !test.Eq("hello world!") {
		panic("test Str.Eq error")
	}

	old := gostr.New(" hello\r\n\t")
	if !old.Eq("hello") {
		panic("test Str.Eq error")
	}
}

func TestStr_TrimSpace(t *testing.T) {
	old := gostr.New(" hello\r\n\t")
	if old.TrimSpace() != "hello" {
		panic("test Str.TrimSpace error")
	}
}
