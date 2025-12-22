package v2

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func MapGoExtFiles(workingDirectory string) ([]string, error) {
	var goFiles []string

	err := filepath.WalkDir(workingDirectory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			matched, err := filepath.Match("*.go", d.Name())
			if err != nil {
				return err
			}
			if matched {
				goFiles = append(goFiles, path)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	fmt.Println(goFiles)

	return goFiles, nil
}
