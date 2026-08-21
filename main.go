package main

import (
	"fmt"
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
	//met tout les nombre dans allnb
	allnb, err := tools.Splitnb(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if len(allnb) == 0 {
		fmt.Println("need nomber")
		return
	}
	//appel de tout les fonctions
	tools.Tricroissant(allnb)
	fmt.Println("Average:", int(tools.Average(allnb)))
	fmt.Println("Median:", int(tools.Median(allnb)))
	fmt.Println("Variance:", int(tools.Variance(allnb)))
	fmt.Println("Standard Deviation:", int(tools.StandardDeviation(allnb)))
}
