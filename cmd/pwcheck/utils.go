package main

import (
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

func readFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(f), "\n"), nil
}

func readInput(showPW bool) (string, error) {
	input, err := terminal.ReadPassword(0)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(input)), nil
}

func getCommand(args []string) (string, []string) {
	if len(args) == 0 {
		return "", nil
	}

	command := ""
	remain := make([]string, 0)
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") {
			remain = append(remain, arg)
		} else if command == "" {
			command = arg
		} else {
			remain = append(remain, arg)
		}
	}

	return command, remain
}

func getFlag(args []string, name string, defValue string) (string, []string) {
	if len(args) == 0 {
		return "", nil
	}

	remain := make([]string, 0)
	value := defValue
	skipNext := false
	for i, arg := range args {
		if skipNext {
			skipNext = false
			continue
		}
		if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
			if strings.TrimPrefix(arg, "-") == name || strings.TrimPrefix(arg, "--") == name {
				if i+1 < len(args) {
					value = args[i+1]
					skipNext = true
				}
			} else {
				remain = append(remain, arg)
			}
		} else {
			remain = append(remain, arg)
		}
	}

	return value, remain
}

func hasFlag(args []string, name string) (bool, []string) {
	remain := make([]string, 0)
	has := false
	for _, arg := range args {
		if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
			if strings.TrimPrefix(arg, "-") == name || strings.TrimPrefix(arg, "--") == name {
				has = true
			} else {
				remain = append(remain, arg)
			}
		} else {
			remain = append(remain, arg)
		}
	}

	return has, remain
}
