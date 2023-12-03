package lib

import (
	"errors"
	"slices"
)

func ChunkEvery[T any](s []T, count int) [][]T {
	r := [][]T{}
	for _, v := range s {
		l, i, e := At(r, -1)

		if e != nil {
			r = append(r, []T{v})
		} else {
			if len(l) < count {
				r[i] = append(l, v)
			} else {
				r = append(r, []T{v})
			}
		}

	}
	return r
}

func At[T any](s []T, i int) (T, int, error) {
	if i < 0 {
		i = len(s) + i
	}

	if i < 0 || i >= len(s) {
		var n T
		return n, -1, errors.New("index out of range")
	}

	return s[i], i, nil
}

func MapOverSlice[T comparable, R any](s []T, f func(T) R) []R {
	r := []R{}
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func DropFirst[T any](s []T) []T {
	if len(s) > 0 {
		return s[1:]
	} else {
		panic("slice is empty")
	}
}

func TakeEvery[T any](s []T, nth int) []T {
	r := []T{}
	for i, v := range s {
		if i%nth == 0 {
			r = append(r, v)
		}
	}
	return r
}

func CopyAndReverseSlice[T any](s []T) []T {
	s = CopySlice(s)
	slices.Reverse(s)
	return s
}

func CopySlice[T any](s []T) []T {
	r := make([]T, len(s))
	copy(r, s)
	return r
}
