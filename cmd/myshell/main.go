package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/commands"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/interfaces"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/shared"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	repl()
}

func repl() {
	defer repl()
	defer shared.Reset()

	commandsList := []interfaces.Command{commands.Pwd, commands.Exit, commands.Echo, commands.Type, commands.Cd}
	fmt.Fprint(os.Stdout, "$ ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	cmd, args := ParseInput(input)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input")
		os.Exit(1)
	}

	customStdout, customSderr := handleRedirect(&args)
	defer customStdout.Close()
	defer customSderr.Close()

	for _, command := range commandsList {
		if cmd == command.Name {
			command.Handler(&args)
			return
		}
	}
	// fmt.Printf("cmd: %v\n", cmd)
	has, _ := commands.ExecCommandHandler(cmd, &args)
	// fmt.Printf("err: %v isNil: %v\n", err, err != nil)
	if !has {
		fmt.Fprintf(shared.Stderr, "%s: command not found\n", cmd)
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

func handleRedirect(args *[]string) (*os.File, *os.File) {
	var stdoutRedirect string = ""
	var stderrRedirect string = ""
	var customStdout *os.File
	var customStderr *os.File
	var stdoutSimble string
	hasStdoutRedirect := false
	hasStderrRedirect := false
	for i, arg := range *args {
		switch arg {
		case ">", "1>", ">>", "1>>":
			hasStdoutRedirect = true
			stdoutSimble = arg
			stdoutRedirect = (*args)[i+1]
			*args = append((*args)[:i], (*args)[i+2:]...)
		case "2>", "2>>":
			hasStderrRedirect = true
			stderrRedirect = (*args)[i+1]
			*args = append((*args)[:i], (*args)[i+2:]...)
		}
	}

	if hasStdoutRedirect {
		var f *os.File
		var err error

		if stdoutSimble == ">" || stdoutSimble == "1>" {
			f, err = os.Create(stdoutRedirect)
		} else {
			f, err = os.OpenFile(stdoutRedirect, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		}
		if err != nil {
			fmt.Println("Erro ao criar arquivo:", err)
		}
		shared.Stdout = f
		customStdout = f
	}

	if hasStderrRedirect {
		var f *os.File
		var err error

		if stderrRedirect == "2>" {
			f, err = os.Create(stderrRedirect)
		} else {
			f, err = os.OpenFile(stderrRedirect, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		}
		if err != nil {
			fmt.Println("Erro ao criar arquivo:", err)
		}
		shared.Stderr = f
		customStderr = f
	}

	return customStdout, customStderr
}
