package commands

import (
	"fmt"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/interfaces"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/shared"
)

var Echo interfaces.Command = interfaces.Command{
	Name: "echo",
	Handler: func(args *[]string) {
		fmt.Fprintf(shared.Stdout, "%s\n", strings.Join(*args, " "))
	},
}
