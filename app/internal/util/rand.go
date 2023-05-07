package util

import "math/rand"

func FlipCoin() bool {
	return rand.Float32() < 0.5
}

func RandIntMinMax(min, max int) int {
	n := max - min + 1
	if n < 0 {
		return 0
	}

	return rand.Intn(n) + min
}

func RandPick[T any](s []T) T {
	var out T
	if len(s) > 0 {
		out = s[rand.Intn(len(s))]
	}

	return out
}
