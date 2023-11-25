package lib

func InvertMap[K, V comparable](m map[K]V) map[V]K {
	r := make(map[V]K, len(m))
	for k, v := range m {
		r[v] = k
	}
	return r
}

func MapIntersect[K, V comparable](maps ...map[K]V) map[K]V {
	if len(maps) == 0 {
		return make(map[K]V)
	}

	result := make(map[K]V)
	for k, v := range maps[0] {
		foundInAll := true
		for _, m := range maps[1:] {
			if _, ok := m[k]; !ok {
				foundInAll = false
				break
			}
		}
		if foundInAll {
			result[k] = v
		}
	}
	return result
}

func MapSomeKey[K, V comparable](m map[K]V) K {
	for k := range m {
		return k
	}
	panic("Can't get key from empty map")
}

func SliceToMap[T comparable](slice []T) map[T]T {
	result := make(map[T]T)
	for _, v := range slice {
		result[v] = v
	}
	return result
}
