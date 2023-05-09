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

func RandMapKey[K comparable, V any](m map[K]V) K {
	var (
		key K
		i   = rand.Intn(len(m))
	)

	for key = range m {
		if i == 0 {
			break
		}
		i--
	}

	return key
}
