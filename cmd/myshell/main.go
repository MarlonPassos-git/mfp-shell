package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

var validCommands = []string{}

func main() {
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	trimmedCommand := strings.TrimSpace(command)
	if err != nil {
		fmt.Println("Error reading input")
		os.Exit(1)
	}

	if !contains(validCommands, trimmedCommand) {
		fmt.Fprintf(os.Stderr, "%s: command not found\n", trimmedCommand)
	}
}

func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
