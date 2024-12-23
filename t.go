package main

import (
	"fmt"
	"os"
)

func main() {
	// f, err := os.Create("a.txt")
	f, err := os.OpenFile("aa.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Erro ao criar arquivo:", err)
	}

	fmt.Fprintln(f, "hi2")
}
