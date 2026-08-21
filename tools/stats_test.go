package tools

import "testing"

func TestAverage(t *testing.T) {
	cas := []struct {
		nom     string
		entree  []int
		attendu int
	}{
		{nom: "exemple standard", entree: []int{3, 4, 5, 6, 7, 8}, attendu: 5}, // 33/6 = 5.5 tronqué à 5
		{nom: "division exacte", entree: []int{2, 4, 6}, attendu: 4},
		{nom: "un seul element", entree: []int{10}, attendu: 10},
		{nom: "avec negatifs", entree: []int{-4, 0, 4}, attendu: 0},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			resultat := int(Average(c.entree))
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
		{nom: "pair - exemple standard", entree: []int{3, 4, 5, 6, 7, 8}, attendu: 5}, // (5+6)/2 = 5.5 tronqué à 5
		{nom: "impair", entree: []int{1, 2, 3, 4, 5}, attendu: 3},
		{nom: "un seul element", entree: []int{7}, attendu: 7},
		{nom: "deux elements", entree: []int{5, 10}, attendu: 7}, // (5+10)/2 = 7.5 tronqué à 7
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			resultat := int(Median(c.entree))
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
		// moyenne tronquee a 5 (pas 5.5), donc variance = 19/6, pas 17.5/6
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
		{nom: "exemple standard", entree: []int{3, 4, 5, 6, 7, 8}, attendu: 1}, // sqrt(3.1667) = 1.78 tronqué à 1
		{nom: "tous identiques", entree: []int{4, 4, 4, 4}, attendu: 0},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			resultat := int(StandardDeviation(c.entree))
			if resultat != c.attendu {
				t.Errorf("entrée: %v\nattendu: %d\nobtenu:  %d", c.entree, c.attendu, resultat)
			}
		})
	}
}
