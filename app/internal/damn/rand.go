package damn

import "math/rand"

func flipCoin() bool {
	return rand.Float32() < 0.5
}

func randIntMinMax(min, max int) int {
	n := max - min + 1
	if n < 0 {
		return 0
	}

	return rand.Intn(n) + min
}

func randPick[T any](s []T) T {
	var out T
	if len(s) > 0 {
		out = s[rand.Intn(len(s))]
	}

	return out
}

func randMapKey[K comparable, V any](m map[K]V) K {
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
