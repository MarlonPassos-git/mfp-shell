package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/interfaces"
)

var Echo interfaces.Command = interfaces.Command{
	Name: "echo",
	Handler: func(args *[]string) {
		fmt.Fprintf(os.Stdout, "%s\n", strings.Join(*args, " "))
	},
}
