package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/christopher-kleine/pwcheck"
)

func validate(args []string) {
	showPW, args := hasFlag(args, "show")
	entropyStr, args := getFlag(args, "entropy", "60")
	showReason, _ := hasFlag(args, "reason")

	entropy, err := strconv.ParseFloat(entropyStr, 64)
	if err != nil {
		help()
		return
	}

	for {
		fmt.Print("Password: ")
		password, err := readInput(showPW)
		if err != nil {
			panic(err)
		}

		if err := pwcheck.Validate(password, entropy); err != nil {
			if showReason {
				fmt.Printf("FAILED: %s\n", err)
			} else {
				fmt.Println("FAILED")
			}
		} else {
			fmt.Println("OK")
		}
	}
}

func validateFile(args []string) {
	showPW, args := hasFlag(args, "show")
	entropyStr, args := getFlag(args, "entropy", "60")
	showReason, args := hasFlag(args, "reason")
	noOk, args := hasFlag(args, "no-ok")
	noFailed, args := hasFlag(args, "no-failed")
	noSummary, args := hasFlag(args, "no-summary")
	summaryOnly, args := hasFlag(args, "summary-only")

	entropy, err := strconv.ParseFloat(entropyStr, 64)
	if err != nil {
		help()
		return
	}

	if len(args) == 0 {
		help()
		return
	}

	summaryOnly = !summaryOnly
	if noOk && noFailed {
		summaryOnly = false
	}

	lines, err := readFile(args[0])
	if err != nil {
		panic(err)
	}

	errors := pwcheck.ValidateSlice(lines, entropy)
	ok := 0
	failed := 0
	if summaryOnly {
		if noOk {
			for i, err := range errors {
				if err == nil {
					lines[i] = ""
				}
			}
		}
		if noFailed {
			for i, err := range errors {
				if err != nil {
					lines[i] = ""
				}
			}
		}

		lPass := 0
		if showPW {
			for _, line := range lines {
				if len(line) > lPass {
					lPass = len(line)
				}
			}
		} else {
			lPass = int(math.Ceil(math.Log10(float64(len(lines)) + 1)))
			if len("Row") > lPass {
				lPass = len("Row")
			}
		}

		lErr := 0
		if showReason {
			for _, err := range errors {
				if err != nil {
					if len(err.Error()) > lErr {
						lErr = len(err.Error())
					}
				}
			}
		}

		header := fmt.Sprintf("%%-%ds | %%6s\n", lPass)
		rowFormat := fmt.Sprintf("%%%dd | %%-6s\n", lPass)
		if showPW && showReason {
			header = fmt.Sprintf("%%-%ds | %%6s | %%-%ds\n", lPass, lErr)
			rowFormat = fmt.Sprintf("%%-%ds | %%-6s | %%-%ds\n", lPass, lErr)
		} else if showPW {
			header = fmt.Sprintf("%%-%ds | %%6s\n", lPass)
			rowFormat = fmt.Sprintf("%%-%ds | %%-6s\n", lPass)
		} else if showReason {
			header = fmt.Sprintf("%%-%ds | %%6s | %%-%ds\n", lPass, lErr)
			rowFormat = fmt.Sprintf("%%%dd | %%-6s | %%-%ds\n", lPass, lErr)
		}

		if showPW && showReason {
			fmt.Printf(header, "Password", "Result", "Reason")
			fmt.Printf(header, strings.Repeat("-", lPass), "------", strings.Repeat("-", lErr))
			for i, line := range lines {
				if errors[i] != nil {
					failed++
				} else {
					ok++
				}

				if line == "" {
					continue
				}
				if errors[i] != nil {
					fmt.Printf(rowFormat, line, "FAILED", errors[i])
				} else {
					fmt.Printf(rowFormat, line, "OK", "")
				}
			}
		} else if showPW {
			fmt.Printf(header, "Password", "Result")
			fmt.Printf(header, strings.Repeat("-", lPass), "------")
			for i, line := range lines {
				if errors[i] != nil {
					failed++
				} else {
					ok++
				}

				if line == "" {
					continue
				}
				if errors[i] != nil {
					fmt.Printf(rowFormat, line, "FAILED")
				} else {
					fmt.Printf(rowFormat, line, "OK")
				}
			}
		} else if showReason {
			fmt.Printf(header, "Row", "Result", "Reason")
			fmt.Printf(header, strings.Repeat("-", lPass), "------", strings.Repeat("-", lErr))
			for i, line := range lines {
				if errors[i] != nil {
					failed++
				} else {
					ok++
				}

				if line == "" {
					continue
				}
				if errors[i] != nil {
					fmt.Printf(rowFormat, i+1, "FAILED", errors[i])
				} else {
					fmt.Printf(rowFormat, i+1, "OK", "")
				}
			}
		} else {
			fmt.Printf(header, "Row", "Result")
			fmt.Printf(header, strings.Repeat("-", lPass), "------")
			for i, line := range lines {
				if errors[i] != nil {
					failed++
				} else {
					ok++
				}

				if line == "" {
					continue
				}
				if errors[i] != nil {
					fmt.Printf(rowFormat, i+1, "FAILED")
				} else {
					fmt.Printf(rowFormat, i+1, "OK")
				}
			}
		}

		fmt.Println()
	} else {
		for _, err := range errors {
			if err != nil {
				failed++
			} else {
				ok++
			}
		}
	}

	if !noSummary {
		fmt.Printf("OK    : %d\n", ok)
		fmt.Printf("FAILED: %d\n", failed)
	}
}
