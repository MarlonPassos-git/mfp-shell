package commands

import (
	"errors"
	"os/exec"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/shared"
)

func ExecCommandHandler(comand string, args *[]string) error {
	has, fullPath := findCommandPath(comand)

	if !has {
		return errors.New("not find")
	}
	cmd := exec.Command(fullPath, (*args)...)
	cmd.Stderr = shared.Stderr
	cmd.Stdout = shared.Stdout

	err := cmd.Run()

	return err
}
