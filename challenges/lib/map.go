package lib

func InvertMap[K, V comparable](m map[K]V) map[V]K {
	r := make(map[V]K, len(m))
	for k, v := range m {
		r[v] = k
	}
	return r
}

func Intersect[K, V comparable](m1 map[K]V, m2 map[K]V) map[K]V {
	m3 := make(map[K]V)
	for k, v := range m1 {
		if _, ok := m2[k]; ok {
			m3[k] = v
		}
	}
	return m3
}

func MapSomeKey[K, V comparable](m map[K]V) K {
	for k := range m {
		return k
	}
	panic("Can't get key from empty map")
}
