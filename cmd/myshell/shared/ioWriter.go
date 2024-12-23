package shared

import (
	"io"
	"os"
)

var Stdout io.Writer = os.Stdout
var Stderr io.Writer = os.Stderr

func Reset() {
	Stdout = os.Stdout
	Stderr = os.Stderr
}
