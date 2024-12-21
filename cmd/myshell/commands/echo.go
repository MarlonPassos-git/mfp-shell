package commands

import (
	"fmt"
	"os"
	"strings"
)

const EchoCommand = "echo"

func EchoCommandHandler(args *[]string) {
	fmt.Fprintf(os.Stdout, "%s\n", strings.Join(*args, " "))
}
