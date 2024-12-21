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

var validCommands = []string{"exit"}

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

	if !contains(validCommands, comand) {
		fmt.Fprintf(os.Stderr, "%s: command not found\n", comand)
	}

	switch comand {
	case commands.ExitCommand:
		commands.ExitCommandHandler(&args)
	}

	repl()
}

func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
