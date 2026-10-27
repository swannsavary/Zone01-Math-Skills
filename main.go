package main

import (
	"math-skills/tools"
	"os"
)

func main() {
	//regade quel fichier regarder
	args := os.Args[1:]
	if len(args) != 1 {
		os.Stderr.WriteString("need a file with numbers\n")
		return
	}
	//lit le fichier et le copie dans file
	file, err := os.ReadFile(args[0])
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}
	//met tout les nombres dans allnb
	allnb, err := tools.Splitnb(file)
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		return
	}

	if len(allnb) == 0 {
		os.Stderr.WriteString("need numbers in file\n")
		return
	}
	//appel de toute les fonctions et les arrondits a l'unite la plus proche
	tools.Tricroissant(allnb)
	tools.PrintStat("Average", tools.Average(allnb))
	tools.PrintStat("Median", tools.Median(allnb))
	tools.PrintStat("Variance", tools.Variance(allnb))
	tools.PrintStat("Standard Deviation", tools.StandardDeviation(allnb))
}
