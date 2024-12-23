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
	var quote string
	for {
		if strings.HasPrefix(s, "'") {
			quote = "'"
		} else if strings.HasPrefix(s, "\"") {
			quote = "\""
		} else {
			quote = ""
		}

		if quote != "" {
			s = s[1:]
			end := strings.Index(s, quote)
			if end == -1 {
				tokens = append(tokens, s)
				break
			}
			token := s[:end]
			tokens = append(tokens, token)
			s = s[end+1:]
		} else {
			space := strings.Index(s, " ")
			if space == -1 {
				tokens = append(tokens, s)
				break
			}
			tokens = append(tokens, s[:space])
			s = s[space+1:]
		}

		s = strings.TrimSpace(s)
		if s == "" {
			break
		}
	}

	if len(tokens) > 0 {
		cmd = strings.ToLower(tokens[0])
		if len(tokens) > 1 {
			args = tokens[1:]
		}
	}

	return cmd, args
}
