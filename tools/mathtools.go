package tools

import "math"

// renvoie la moyenne d'un slice d'int
func Average(nbs []int) float64 {
	var result int
	for _, nb := range nbs {
		result += nb
	}
	return float64(result) / float64(len(nbs))
}

// renvoie la medianne d'un slice d'int
func Median(nbs []int) float64 {
	if len(nbs)%2 == 0 {
		return float64(nbs[len(nbs)/2-1]+nbs[len(nbs)/2]) / 2
	}
	return float64(nbs[len(nbs)/2])
}

// renvoie la variance d'un slice d'int
func Variance(nbs []int) float64 {
	moyenne := Average(nbs)
	var somme float64
	for _, nb := range nbs {
		diff := float64(nb) - moyenne
		somme += diff * diff
	}
	return somme / float64(len(nbs))
}

// renvoi l'ecart type d'un slice d'int
func StandardDeviation(nbs []int) float64 {
	return math.Sqrt(Variance(nbs))
}
