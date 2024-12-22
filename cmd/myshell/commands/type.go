package commands

import (
	"fmt"
	"os"
	"strings"
)

const TypeCommand = "type"

var allBuiltinComands = []string{TypeCommand, EchoCommand, ExitCommand, PwdCommand}

func TypeCommandHandler(args *[]string) {
	if len(*args) == 0 {
		return
	}

	command := (*args)[0]

	if handleBuiltin(command) {
		return
	}

	hasInPath, fullPath := findCommandPath(command)

	if hasInPath {
		fmt.Fprintf(os.Stdout, "%s is %s\n", command, fullPath)

		return
	}

	fmt.Fprintf(os.Stderr, "%s: not found\n", command)
}

func handleBuiltin(command string) bool {
	for _, builtinCommand := range allBuiltinComands {
		if builtinCommand != command {
			continue
		}

		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", command)
		return true
	}
	return false
}

func findCommandPath(command string) (bool, string) {
	envPath := os.Getenv("PATH")
	paths := strings.Split(envPath, ":")

	for _, path := range paths {
		files, err := os.ReadDir(path)
		if err != nil {
			continue
		}

		for _, file := range files {
			if file.IsDir() || file.Name() != command {
				continue
			}

			fullPath := fmt.Sprintf("%s/%s", path, file.Name())
			return true, fullPath
		}

	}

	return false, ""
}
