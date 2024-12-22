package commands

import (
	"fmt"
	"os"
)

const PwdCommand = "pwd"

func PwdCommandHandler() {
	dir, _ := os.Getwd()

	fmt.Fprintln(os.Stdout, dir)
}
