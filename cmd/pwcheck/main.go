package main

import (
	"os"
)

func main() {
	subcmd, remain := getCommand(os.Args[1:])
	switch subcmd {
	case "validate":
		validate(remain)
	case "validate-file":
		validateFile(remain)
	case "calc":
		calc(remain)
	case "calc-file":
		calcFile(remain)
	case "help":
		help()
	case "version":
		version()
	default:
		help()
	}
}
