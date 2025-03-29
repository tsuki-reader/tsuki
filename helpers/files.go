package helpers

import (
	"os"
	"tsuki/core"
)

func FileExists(file string) bool {
	_, statErr := os.Stat(file)
	return !(statErr != nil && os.IsNotExist(statErr))
}

func CreateAndWriteToFile(file string, contents string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(contents)
	if err != nil {
		return err
	}

	return nil
}

func CreateDirectory(location string) error {
	err := os.MkdirAll(location, 0700)
	return err
}

func ReadFileContents(file string) (string, error) {
	contents, err := os.ReadFile(file)
	if err != nil {
		core.CONFIG.Logger.Println(file)
		return "", err
	}

	return string(contents), err
}
