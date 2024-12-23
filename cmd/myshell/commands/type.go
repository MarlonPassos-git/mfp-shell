package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/interfaces"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/shared"
)

var allBuiltinComands = []string{Exit.Name, Pwd.Name, Cd.Name, Exit.Name, Pwd.Name, Echo.Name, "type"}

var Type interfaces.Command = interfaces.Command{
	Name:    "type",
	Handler: handler,
}

func handler(args *[]string) {
	if len(*args) == 0 {
		return
	}

	command := (*args)[0]

	if handleBuiltin(command) {
		return
	}

	hasInPath, fullPath := findCommandPath(command)

	if hasInPath {
		fmt.Fprintf(shared.Stdout, "%s is %s\n", command, fullPath)

		return
	}

	fmt.Fprintf(shared.Stderr, "%s: not found\n", command)
}

func handleBuiltin(command string) bool {
	for _, builtinCommand := range allBuiltinComands {
		if builtinCommand != command {
			continue
		}

		fmt.Fprintf(shared.Stdout, "%s is a shell builtin\n", command)
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
