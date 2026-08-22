package tools

import (
	"math"
	"testing"
)

func TestAverage(t *testing.T) {
	cas := []struct {
		nom     string
		entree  []int
		attendu int
	}{
		{nom: "exemple standard", entree: []int{3, 4, 5, 6, 7, 8}, attendu: 6}, // 33/6 = 5.5 arrondi à 6
		{nom: "division exacte", entree: []int{2, 4, 6}, attendu: 4},
		{nom: "un seul element", entree: []int{10}, attendu: 10},
		{nom: "avec negatifs", entree: []int{-4, 0, 4}, attendu: 0},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			resultat := int(math.Round(Average(c.entree)))
			if resultat != c.attendu {
				t.Errorf("entrée: %v\nattendu: %d\nobtenu:  %d", c.entree, c.attendu, resultat)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	cas := []struct {
		nom     string
		entree  []int // suppose deja trie, comme l'exige la fonction
		attendu int
	}{
		{nom: "pair - exemple standard", entree: []int{3, 4, 5, 6, 7, 8}, attendu: 6}, // (5+6)/2 = 5.5 arrondit à 6
		{nom: "impair", entree: []int{1, 2, 3, 4, 5}, attendu: 3},
		{nom: "un seul element", entree: []int{7}, attendu: 7},
		{nom: "deux elements", entree: []int{5, 10}, attendu: 8}, // (5+10)/2 = 7.5 arrondit à 8
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			resultat := int(math.Round(Median(c.entree)))
			if resultat != c.attendu {
				t.Errorf("entrée: %v\nattendu: %d\nobtenu:  %d", c.entree, c.attendu, resultat)
			}
		})
	}
}

func TestVariance(t *testing.T) {
	cas := []struct {
		nom     string
		entree  []int
		attendu float64
	}{
		{nom: "exemple standard", entree: []int{3, 4, 5, 6, 7, 8}, attendu: 17.5 / 6},
		{nom: "tous identiques", entree: []int{4, 4, 4, 4}, attendu: 0},
		{nom: "un seul element", entree: []int{10}, attendu: 0},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			resultat := Variance(c.entree)
			// comparaison flottante avec tolerance
			diff := resultat - c.attendu
			if diff < -0.0001 || diff > 0.0001 {
				t.Errorf("entrée: %v\nattendu: %f\nobtenu:  %f", c.entree, c.attendu, resultat)
			}
		})
	}
}

func TestStandardDeviation(t *testing.T) {
	cas := []struct {
		nom     string
		entree  []int
		attendu int
	}{
		{nom: "exemple standard", entree: []int{3, 4, 5, 6, 7, 8}, attendu: 2}, // sqrt(3.1667) = 1.78 arrondit à 2
		{nom: "tous identiques", entree: []int{4, 4, 4, 4}, attendu: 0},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			resultat := int(math.Round(StandardDeviation(c.entree)))
			if resultat != c.attendu {
				t.Errorf("entrée: %v\nattendu: %d\nobtenu:  %d", c.entree, c.attendu, resultat)
			}
		})
	}
}
