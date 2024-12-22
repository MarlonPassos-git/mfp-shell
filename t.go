package main

import (
	"fmt"
)

func main() {
	str := "exemplo de string"
	var result string

	for _, char := range str {
		if char != ' ' {
			result += string(char)
		}
	}

	fmt.Println("Resultado:", result)
}
