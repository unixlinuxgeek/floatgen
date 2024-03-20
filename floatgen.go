package floatgen

import "math/rand"

func GenRan(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
