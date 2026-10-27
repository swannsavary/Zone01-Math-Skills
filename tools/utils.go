package tools

import (
	"errors"
	"math"
	"os"
	"strconv"
	"strings"
)

// separe un slice de byte a chaque \n afin de recuperer les nombres du fichier
func Splitnb(file []byte) ([]int, error) {
	nbs := strings.Split(string(file), "\n")
	var result []int
	for _, nb := range nbs {
		nb = strings.TrimSpace(nb)
		if nb != "" {
			nombre, err := strconv.Atoi(nb)
			if err != nil {
				return nil, errors.New("not a nomber")

			}
			result = append(result, nombre)
		}
	}
	return result, nil
}

// tri a bulle d'un slice de int
func Tricroissant(nbs []int) {
	for i := 0; i < len(nbs)-1; i++ {
		for e := 0; e < len(nbs)-i-1; e++ {
			if nbs[e] > nbs[e+1] {
				nbs[e], nbs[e+1] = nbs[e+1], nbs[e]
			}
		}
	}
}

// affiche les resultats sous la forme "label" : valeur
func PrintStat(label string, value float64) {
	os.Stdout.WriteString(label + ": " + strconv.Itoa(int(math.Round(value))) + "\n")
}
