package shared

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func mapForFiles() ([]string, error) {
	var go_files []string
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
			go_files = append(go_files, path)
		}

		return nil
	})

	return go_files, err
}

func readFile(go_files []string, old_arg, new_arg string) error {
	var fileLines string
	for _, file_path := range go_files {
		readFile, err := os.OpenFile(file_path, os.O_RDONLY, 0644)
		if err != nil {
			fmt.Println("Couldn't open file: ", readFile, "Error: ", err)
			return err
		}

		// Make a map to make a atomic change
		//key will be the name of the temp file, and the value will be the original name
		fileKeyVal := make(map[string]string)
		temp_file_path := file_path + ".temp"
		fileKeyVal[temp_file_path] = file_path

		fileScanner := bufio.NewScanner(readFile)
		fileScanner.Split(bufio.ScanLines)

		tempFile, err := os.OpenFile(temp_file_path, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening the tempfile", err)
			return err
		}

		for fileScanner.Scan() {

			fileLines = fileScanner.Text()

			if strings.Contains(fileLines, old_arg) {
				fileLines = strings.ReplaceAll(fileLines, old_arg, new_arg)
				fmt.Println("✓ Replaced:", fileLines) // Debug
			}

			// ALWAYS write the line (whether replaced or not)
			fileLines = fileLines + "\n"
			tempFile.Write([]byte(fileLines))

		}
		tempFile.Close()
		readFile.Close()

		err = os.Rename(temp_file_path, file_path)
		if err != nil {
			return err
		}
	}

	fmt.Println(fileLines)

	return nil
}

func Fileio(old_arg, new_arg string) {
	arr, err := mapForFiles()
	if err != nil {
		fmt.Print("error: ", err)
	}

	err = readFile(arr, old_arg, new_arg)
	if err != nil {
		fmt.Print("Error \n", err)
	}
}
