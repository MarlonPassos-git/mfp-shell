package commands

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/interfaces"
)

var Cd interfaces.Command = interfaces.Command{
	Name: "cd",
	Handler: func(args *[]string) {
		var path string

		if len(*args) < 1 {
			path = "~"
		} else {
			path = (*args)[0]
		}

		err := os.Chdir(path)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", path)
		}
	},
}
