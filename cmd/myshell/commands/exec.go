package commands

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func ExecCommandHandler(comand string, args *[]string) error {
	has, fullPath := findCommandPath(comand)

	if !has {
		return errors.New("not find")
	}
	var cmd *exec.Cmd
	if len(*args) > 0 {
		cmd = exec.Command(fullPath, (*args)...)
	} else {
		cmd = exec.Command(comand)
	}

	output, _ := cmd.CombinedOutput()

	fmt.Println(strings.TrimSpace(string(output)))
	return nil
}
