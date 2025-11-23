package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// panic("a problem")

	path := filepath.Join(os.TempDir(), "file")
	_, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("File created at", path)
}
