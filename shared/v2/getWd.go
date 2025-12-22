package v2

import "os"

func GetWorkDirectory() (string, error) {
	workingDirectory, err := os.Getwd()
	if err != nil {
		return "", err
	}
	//
	return workingDirectory, nil
}
