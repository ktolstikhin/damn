package util

import "math/rand"

func FlipCoin() bool {
	return rand.Float32() < 0.5
}

func RandMinMaxInt(min, max int) int {
	if min > max {
		return 0
	}

	n := max - min + 1
	if n < 0 {
		return 0
	}

	return rand.Intn(n) + min
}

func RandStr(ss []string) string {
	if len(ss) > 0 {
		return ss[rand.Intn(len(ss))]
	}

	return ""
}
