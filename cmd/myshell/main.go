package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/commands"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	repl()
}

func repl() {
	fmt.Fprint(os.Stdout, "$ ")
	fullCommand, err := bufio.NewReader(os.Stdin).ReadString('\n')
	fullCommand = strings.TrimSpace(fullCommand)
	part := strings.Fields(fullCommand)
	comand := part[0]
	args := part[1:]
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input")
		os.Exit(1)
	}

	switch comand {
	case commands.ExitCommand:
		commands.ExitCommandHandler(&args)
	case commands.EchoCommand:
		commands.EchoCommandHandler(&args)
	case commands.TypeCommand:
		commands.TypeCommandHandler(&args)
	default:
		err := commands.ExecCommandHandler(comand, &args)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: command not found\n", comand)
		}
	}

	repl()
}
