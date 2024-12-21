package commands

import (
	"fmt"
	"os"
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

	fmt.Fprintf(os.Stderr, "%s: not found\n", command)
}
