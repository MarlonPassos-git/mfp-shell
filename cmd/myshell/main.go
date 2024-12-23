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
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	cmd, args := parseInput(input)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input")
		os.Exit(1)
	}

	for _, command := range commandsList {
		if cmd == command.Name {
			command.Handler(&args)
			return
		}
	}

	err = commands.ExecCommandHandler(cmd, &args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: command not found\n", cmd)
	}
}

func parseInput(input string) (cmd string, args []string) {
	s := strings.Trim(input, "\r\n")
	var tokens []string
	for {
		start := strings.Index(s, "'")
		if start == -1 {
			tokens = append(tokens, strings.Fields(s)...)
			break
		}
		tokens = append(tokens, strings.Fields(s[:start])...)
		s = s[start+1:]
		end := strings.Index(s, "'")
		token := s[:end]
		tokens = append(tokens, token)
		s = s[end+1:]
	}
	cmd = strings.ToLower(tokens[0])
	if len(tokens) > 1 {
		args = tokens[1:]
	}

	return cmd, args
}
