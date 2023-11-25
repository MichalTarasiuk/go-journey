package lib

import (
	"errors"
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

func At[T any](s [][]T, i int) ([]T, int, error) {

	if i < 0 {
		i = len(s) + i
	}

	if i < 0 || i >= len(s) {
		return nil, -1, errors.New("index out of range")
	}

	return s[i], i, nil
}
