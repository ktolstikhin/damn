package util

import "math/rand"

func MergeMaps[K comparable, V any](m1 map[K]V, m2 map[K]V) map[K]V {
	merged := make(map[K]V, len(m1))
	for k, v := range m1 {
		merged[k] = v
	}
	for k, v := range m2 {
		merged[k] = v
	}
	return merged
}

func RandMapEntry[K comparable, V any](m map[K]V) (K, V) {
	var (
		k K
		v V
		i = rand.Intn(len(m))
	)

	if len(m) == 0 {
		return k, v
	}

	for k, v = range m {
		if i == 0 {
			break
		}
		i--
	}

	return k, v
}
