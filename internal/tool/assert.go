package tool

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func GetFolderFrom(url string) (string, error) {
	path := filepath.Join(os.Getenv(HomeName), SaveFolder)

	absolutePath, err := filepath.Abs(filepath.Join(os.Getenv(HomeName), SaveFolder, filepath.Base(url)))
	if err != nil {
		return "", errors.WithStack(err)
	}

	relPath, err := filepath.Rel(path, absolutePath)
	if err != nil {
		return "", errors.WithStack(err)
	}

	if strings.Contains(relPath, "..") {
		return "", errors.WithStack(errors.New("Your download file may have a path traversal attack"))
	}

	return absolutePath, nil
}

func IsFolderExisted(folder string) bool {
	_, err := os.Stat(folder)
	return os.IsExist(err)
}
