package commands

import (
	"fmt"
	"os"
	"strings"
)

const TypeCommand = "type"

var allBuiltinComands = []string{TypeCommand, EchoCommand, ExitCommand}

func TypeCommandHandler(args *[]string) {
	if len(*args) == 0 {
		return
	}

	command := (*args)[0]

	for _, builtinCommand := range allBuiltinComands {
		if builtinCommand != command {
			continue
		}

		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", command)

		return
	}

	sucess := handlePath(command)

	if sucess {
		return
	}

	fmt.Fprintf(os.Stderr, "%s: not found\n", command)
}

func handlePath(command string) bool {
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

			fmt.Fprintf(os.Stdout, "%s is %s/%s\n", command, path, file.Name())
			return true
		}

	}

	return false
}
