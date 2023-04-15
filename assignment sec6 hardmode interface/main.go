package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := os.Args[1]
	openFileOrCreate(fileName)
}

func openFileOrCreate(fileName string) {
	f, err := os.Open(fileName) // For read access.
	if err != nil {
		_, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
	}
	io.Copy(os.Stdout, f)
}
