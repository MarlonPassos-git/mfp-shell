package commands

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/interfaces"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/shared"
)

var Pwd interfaces.Command = interfaces.Command{
	Name: "pwd",
	Handler: func(_ *[]string) {
		dir, _ := os.Getwd()

		fmt.Fprintln(shared.Stdout, dir)
	},
}
