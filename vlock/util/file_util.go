package util

import (
	"os"
	"strings"
	"errors"
	"os/exec"
	"path/filepath"
)

func GetCurrentPath() (string, error) {
    file, err := exec.LookPath(os.Args[0])
    if err != nil {
        return "", err
    }
    path, err := filepath.Abs(file)
    if err != nil {
        return "", err
    }
    i := strings.LastIndex(path, "/")
    if i < 0 {
        i = strings.LastIndex(path, "\\")
    }
    if i < 0 {
		err := errors.New(`error: Can't find "/" or "\".`)
        return "", err
    }
    return string(path[0 : i+1]), nil
}