package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	dir := "/usr/local/share/nvm/current/bin" // Especifique o diretório que você deseja listar

	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Printf("isRegular: %v, name: %s, type: %s\n", file.Type().IsRegular(), file.Name(), file.Type())
	}
}
