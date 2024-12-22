package commands

import (
	"errors"
	"os"
	"os/exec"
)

func ExecCommandHandler(comand string, args *[]string) error {
	has, fullPath := findCommandPath(comand)

	if !has {
		return errors.New("not find")
	}
	cmd := exec.Command(fullPath, (*args)...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()

	return err
}
func handle() {

}
