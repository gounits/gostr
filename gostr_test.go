// @Time  : 2022/8/23 12:05
// @Email: jtyoui@qq.com
// @Author: 张伟

package gostr_test

import (
	"fmt"
	"github.com/gounits/gostr"
	"testing"
)

func TestGoStr_ToString(t *testing.T) {
	value := gostr.New("hello").ToString()
	if value != "hello" {
		t.Error(`GoStr("hello").ToString()!="hello"`)
	}
}

func TestGoStr_In(t *testing.T) {
	value := gostr.New("hello").In("ll")
	if value != true {
		t.Error(`GoStr("hello").In("ll")!=true`)
	}
}

func TestGoStr_Capitalize(t *testing.T) {
	b := gostr.New("hello").Capitalize()
	if b != "Hello" {
		fmt.Println(b)
		t.Error(`GoStr("hello").Capitalize()!="Hello"`)
	}
}
