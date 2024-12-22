package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/commands"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/interfaces"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	repl()
}

func repl() {
	defer repl()

	commandsList := []interfaces.Command{commands.Pwd, commands.Exit, commands.Echo, commands.Type, commands.Cd}
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

	for _, cmd := range commandsList {
		if comand == cmd.Name {
			cmd.Handler(&args)
			return
		}
	}

	err = commands.ExecCommandHandler(comand, &args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: command not found\n", comand)
	}
}
