package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	targetDirectory := "."

	scriptDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = filepath.Walk(targetDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != scriptDir && path != currentDir && !strings.HasPrefix(filepath.Base(path), ".") {
			codeExtractFolder := filepath.Join(path, "CodeExtract")
			androidProjectFolder := filepath.Join(path, "AndroidProject")

			if err := os.MkdirAll(codeExtractFolder, os.ModePerm); err != nil {
				return err
			}
			if err := os.MkdirAll(androidProjectFolder, os.ModePerm); err != nil {
				return err
			}
			fmt.Printf("Folders created inside %s\n", path)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Folders created successfully.")
}
