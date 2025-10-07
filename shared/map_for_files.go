package shared

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

func mapForFiles() ([]string, error) {
	var goFiles []string
	working_directory, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	err = filepath.WalkDir(working_directory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Check if it's a .go file
		if filepath.Ext(path) == ".go" {
			goFiles = append(goFiles, path)
		}

		return nil
	})

	return goFiles, err
}

func readFile(goFiles []string, old_arg, new_arg string) error {
	lenght := len(goFiles)
	deleted := color.RGB(232, 90, 102)
	added := color.RGB(121, 232, 90)

	var fileLines string
	for _, filePath := range goFiles {
		readFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't open file: %s", filePath)
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			readFile.Close()
			return err
		}

		// Make a map to make a atomic change
		//key will be the name of the temp file, and the value will be the original name
		fileKeyVal := make(map[string]string)
		temp_file_path := filePath + ".temp"
		fileKeyVal[temp_file_path] = filePath

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)

		tempFile, err := os.OpenFile(temp_file_path, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening the tempfile", err)
			tempFile.Close()
			return err
		}

		for fileScanner.Scan() {

			fileLines = fileScanner.Text()

			if strings.Contains(fileLines, old_arg) {
				fileLines = strings.ReplaceAll(fileLines, old_arg, new_arg)
			}

			fileLines = fileLines + "\n"
			tempFile.Write([]byte(fileLines))

		}
		tempFile.Close()
		readFile.Close()

		err = os.Rename(temp_file_path, filePath)
		if err != nil {
			return err
		}
	}

	fmt.Fprintf(os.Stdout, "+ ")
	added.Fprintf(os.Stdout, "%s\n", new_arg)
	fmt.Fprintf(os.Stdout, "- ")
	deleted.Fprintf(os.Stdout, "%s\n", old_arg)
	fmt.Fprintf(os.Stdout, "%d", lenght)
	fmt.Fprintf(os.Stdout, " files modified\n")

	return nil
}

func Fileio(oldArg, newArg string) {
	arr, err := mapForFiles()
	if err != nil {
		fmt.Print("error: ", err)
	}

	err = readFile(arr, oldArg, newArg)
	if err != nil {
		fmt.Print("Error \n", err)
	}
}
