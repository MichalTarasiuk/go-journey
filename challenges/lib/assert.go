package lib

import "fmt"

func Assert(v bool) {
	if !v {
		panic("Assertion failed")
	}
}

func Assertf(cond bool, message string, args ...any) {
	if !cond {
		panic(fmt.Sprintf(message, args...))
	}
}

func Panicf(s string, args ...any) {
	Assertf(false, s, args...)
}

func AssertNil(v any) {
	if v != nil {
		panic(v)
	}
}
