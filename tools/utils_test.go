package tools

import (
	"reflect"
	"testing"
)

func TestSplitnb(t *testing.T) {
	cas := []struct {
		nom            string
		entree         string
		attendu        []int
		erreurAttendue bool
	}{
		{
			nom:     "plusieurs nombres",
			entree:  "4\n8\n6\n5\n3\n7",
			attendu: []int{4, 8, 6, 5, 3, 7},
		},
		{
			nom:     "avec ligne vide finale",
			entree:  "1\n2\n3\n",
			attendu: []int{1, 2, 3},
		},
		{
			nom:     "un seul nombre",
			entree:  "42",
			attendu: []int{42},
		},
		{
			nom:     "chaine vide",
			entree:  "",
			attendu: nil,
		},
		{
			nom:            "contient du texte non numerique",
			entree:         "1\nabc\n3",
			erreurAttendue: true,
		},
		{
			nom:     "nombres negatifs",
			entree:  "-5\n3\n-1",
			attendu: []int{-5, 3, -1},
		},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			resultat, err := Splitnb([]byte(c.entree))
			if c.erreurAttendue {
				if err == nil {
					t.Errorf("attendu une erreur, obtenu nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("erreur inattendue: %v", err)
			}
			if !reflect.DeepEqual(resultat, c.attendu) {
				t.Errorf("entrée: %q\nattendu: %v\nobtenu:  %v", c.entree, c.attendu, resultat)
			}
		})
	}
}

func TestTricroissant(t *testing.T) {
	cas := []struct {
		nom     string
		entree  []int
		attendu []int
	}{
		{
			nom:     "desordre simple",
			entree:  []int{4, 8, 6, 5, 3, 7},
			attendu: []int{3, 4, 5, 6, 7, 8},
		},
		{
			nom:     "deja trie",
			entree:  []int{1, 2, 3},
			attendu: []int{1, 2, 3},
		},
		{
			nom:     "ordre inverse",
			entree:  []int{5, 4, 3, 2, 1},
			attendu: []int{1, 2, 3, 4, 5},
		},
		{
			nom:     "avec negatifs",
			entree:  []int{3, -1, 0, -5},
			attendu: []int{-5, -1, 0, 3},
		},
		{
			nom:     "un seul element",
			entree:  []int{42},
			attendu: []int{42},
		},
		{
			nom:     "slice vide",
			entree:  []int{},
			attendu: []int{},
		},
	}

	for _, c := range cas {
		t.Run(c.nom, func(t *testing.T) {
			Tricroissant(c.entree) // trie en place, pas de valeur de retour
			if !reflect.DeepEqual(c.entree, c.attendu) {
				t.Errorf("attendu: %v\nobtenu:  %v", c.attendu, c.entree)
			}
		})
	}
}
