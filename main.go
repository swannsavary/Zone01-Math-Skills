package main

import (
	"fmt"
	"math"
	"math-skills/tools"
	"os"
)

func main() {
	//regade quel fichier regarder
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("need a file with numbers")
		return
	}
	//lit le fichier et le copie dans file
	file, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	//met tout les nombres dans allnb
	allnb, err := tools.Splitnb(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if len(allnb) == 0 {
		fmt.Println("need nomber in file")
		return
	}
	//appel de tout les fonctions et les arrondits a l'unite la plus proche
	tools.Tricroissant(allnb)
	fmt.Println("Average:", int(math.Round(tools.Average(allnb))))
	fmt.Println("Median:", int(math.Round(tools.Median(allnb))))
	fmt.Println("Variance:", int(math.Round(tools.Variance(allnb))))
	fmt.Println("Standard Deviation:", int(math.Round(tools.StandardDeviation(allnb))))
}
