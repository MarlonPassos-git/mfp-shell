package commands

import (
	"fmt"
	"os"
	"strconv"
)

const ExitCommand = "exit"

func ExitCommandHandler(args *[]string) {
	var exitCode int

	if len(*args) < 1 {
		exitCode = 0
	} else {
		first := (*args)[0]
		num, err := strconv.Atoi(first)

		if err != nil {
			fmt.Fprintf(os.Stderr, "exit: %s: numeric argument required\n", first)
			return
		}

		exitCode = num
	}

	os.Exit(exitCode)
}
