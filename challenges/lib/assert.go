package lib

import "fmt"

func Assertf(cond bool, message string, args ...any) {
	if !cond {
		panic(fmt.Sprintf(message, args...))
	}
}
