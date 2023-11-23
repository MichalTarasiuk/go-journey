package lib

import "fmt"

func Assertf(cond bool, message string, args ...any) {
	if !cond {
		panic(fmt.Sprintf(message, args...))
	}
}

func Panicf(s string, args ...any) {
	Assertf(false, s, args...)
}
