package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	targetDirectory := "."
	scriptDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	currentDir, _ := os.Getwd()

	filepath.Walk(targetDirectory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && path != scriptDir && path != currentDir && !strings.HasPrefix(filepath.Base(path), ".") && filepath.Dir(path) == targetDirectory {
			codeExtractFolder := filepath.Join(path, "CodeExtract")
			androidProjectFolder := filepath.Join(path, "AndroidProject")
			os.MkdirAll(codeExtractFolder, os.ModePerm)
			os.MkdirAll(androidProjectFolder, os.ModePerm)
			fmt.Printf("%s :: Folders created inside %s\n", time.Now().Format("2006-01-02 15:04:05"), path)
		}
		return nil
	})

	filepath.Walk(targetDirectory, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Ext(info.Name()) == ".apk" {
			cmd := exec.Command("jadx", "-d", filepath.Join(currentDir, filepath.Dir(path), "CodeExtract"), path)
			cmd.CombinedOutput()
		}
		return nil
	})

	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " :: Completed")
}
