package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	isRenew     bool
	isDecompile bool
)

var rootCmd = &cobra.Command{
	Use:   "extract-apks",
	Short: "Extract APK files from a directory",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	targetDirectory := "."
	scriptDir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	currentDir, _ := os.Getwd()

	filepath.Walk(targetDirectory, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && path != scriptDir && path != currentDir && !strings.HasPrefix(filepath.Base(path), ".") && filepath.Base(path) != "templates" && filepath.Dir(path) == targetDirectory {
			codeExtractFolder := filepath.Join(path, "CodeExtract")
			androidProjectFolder := filepath.Join(path, "AndroidProject")

			if isRenew {
				os.RemoveAll(androidProjectFolder)
			}
			os.MkdirAll(codeExtractFolder, os.ModePerm)
			os.MkdirAll(androidProjectFolder, os.ModePerm)
			fmt.Printf("%s :: Folders created inside %s\n", time.Now().Format("2006-01-02 15:04:05"), path)
		}
		return nil
	})

	if isDecompile {
		filepath.Walk(targetDirectory, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() && filepath.Ext(info.Name()) == ".apk" {
				cmd := exec.Command("jadx", "-d", filepath.Join(currentDir, filepath.Dir(path), "CodeExtract"), path)
				_, err := cmd.CombinedOutput()
				if err != nil {
					fmt.Printf("%s :: Error when extract :: %s", time.Now().Format("2006-01-02 15:04:05"), path)
				} else {
					fmt.Printf("%s :: %s extracted\n", time.Now().Format("2006-01-02 15:04:05"), path)
				}
			}
			return nil
		})
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05") + " :: Completed")
}

func main() {
	rootCmd.Flags().BoolVarP(&isRenew, "renewal", "r", false, "Remove old folders")
	rootCmd.Flags().BoolVarP(&isDecompile, "decompile", "d", false, "Decompile apks")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
