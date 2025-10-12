package shared

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func CreateBackup() error {
	folderPath := "backup_folder"

	// Create backup folder if not exists
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create backup folder: %w", err)
	}

	workingDirectory, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	// Walk through all files/folders recursively
	err = filepath.WalkDir(workingDirectory, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip the backup folder itself (prevents infinite loop)
		if d.IsDir() && filepath.Base(path) == folderPath {
			return filepath.SkipDir
		}

		// Only copy files
		if !d.IsDir() {
			relPath, err := filepath.Rel(workingDirectory, path)
			if err != nil {
				return err
			}

			// Create a backup copy with .refx extension
			destPath := filepath.Join(folderPath, relPath+".refx")

			// Ensure destination directories exist
			if err := os.MkdirAll(filepath.Dir(destPath), os.ModePerm); err != nil {
				return err
			}

			src, err := os.Open(path)
			if err != nil {
				return err
			}
			defer src.Close()

			dst, err := os.Create(destPath)
			if err != nil {
				src.Close()
				return err
			}
			defer dst.Close()

			// Copy file content
			if _, err := io.Copy(dst, src); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("backup failed: %w", err)
	}

	return nil
}
