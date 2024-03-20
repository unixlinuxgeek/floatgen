package floatgen


func genRan(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
