package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

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
	cmd, args := ParseInput(input)

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

func ParseInput(s string) (string, []string) {
	s = strings.TrimSuffix(s, "\n")
	var result []string
	var current []rune
	var quote rune
	var nestedQuote rune
	escaped := false

	for i, r := range s {
		switch {
		case escaped:
			current = append(current, r)
			escaped = false
		case r == '\\':
			if nestedQuote != '\'' && quote != '\'' {
				if quote == 0 || (quote != 0 && (s[i+1] == '"' || s[i+1] == '\\' || s[i+1] == '$')) {
					escaped = true
				} else {
					current = append(current, r)
				}
			} else {
				current = append(current, r)
			}
		case quote != 0:
			if r == quote {
				quote = 0
			} else {
				if r == '"' || r == '\'' {
					if nestedQuote == r {
						nestedQuote = 0
					} else {
						nestedQuote = r
					}
				}
				current = append(current, r)
			}
		case r == '"' || r == '\'':
			quote = r
		case unicode.IsSpace(r):
			if len(current) > 0 {
				result = append(result, string(current))
				current = nil
			}
		default:
			current = append(current, r)
		}
	}

	if len(current) > 0 {
		result = append(result, string(current))
	}

	cmd := result[0]
	args := result[1:]

	return cmd, args
}
