package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/christopher-kleine/pwcheck"
)

func calc(args []string) {
	showPW, _ := hasFlag(args, "show")

	for {
		fmt.Print("Password: ")
		password, err := readInput(showPW)
		if err != nil {
			panic(err)
		}

		entropy := pwcheck.GetEntropy(password)
		fmt.Printf("%f\n", entropy)
	}
}

func calcFile(args []string) {
	showPW, args := hasFlag(args, "show")

	if len(args) == 0 {
		help()
		return
	}

	lines, err := readFile(args[0])
	if err != nil {
		panic(err)
	}

	entropies := pwcheck.GetEntropySlice(lines)
	if showPW {
		l := 0
		for _, line := range lines {
			if len(line) > l {
				l = len(line)
			}
		}
		header := fmt.Sprintf("%%-%ds | %%12s\n", l)
		rowFormat := fmt.Sprintf("%%-%ds | %%12.6f\n", l)

		fmt.Printf(header, "Password", "Entropy")
		fmt.Printf(header, strings.Repeat("-", l), strings.Repeat("-", 12))
		for i, line := range lines {
			fmt.Printf(rowFormat, line, entropies[i])
		}
	} else {
		l := int(math.Ceil(math.Log10(float64(len(lines)) + 1)))
		if len("Row") > l {
			l = len("Row")
		}
		header := fmt.Sprintf("%%-%ds | %%12s\n", l)
		rowFormat := fmt.Sprintf("%%%dd | %%12.6f\n", l)

		fmt.Printf(header, "Row", "Entropy")
		fmt.Printf(header, strings.Repeat("-", l), strings.Repeat("-", 12))
		for i, entropy := range entropies {
			fmt.Printf(rowFormat, i+1, entropy)
		}
	}
}
