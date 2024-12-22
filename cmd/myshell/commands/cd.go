package commands

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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

		if strings.HasPrefix(path, "~") {
			home := os.Getenv("HOME")
			path = strings.Replace(path, "~", home, -1)
		} else if strings.HasPrefix(path, ".") {
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			path = filepath.Join(pwd, path)
		}

		err := os.Chdir(path)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", path)
		}
	},
}
