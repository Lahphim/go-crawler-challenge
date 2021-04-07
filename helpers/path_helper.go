package helpers

import (
	"path"
	"path/filepath"
	"runtime"
)

// RootDir returns the current path of the project
func RootDir() string {
	_, currentFile, _, _ := runtime.Caller(0)
	currentFilePath := path.Join(path.Dir(currentFile))
	return filepath.Dir(currentFilePath)
}
