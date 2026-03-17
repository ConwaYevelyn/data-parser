package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetAbsolutePath returns the absolute path of a file or directory.
func GetAbsolutePath(path string) string {
	absPath, err := filepath.Abs(path)
	if err!= nil {
		log.Fatal(err)
	}
	return absPath
}

// ReadJSONFile reads the contents of a JSON file and unmarshals it into a struct.
func ReadJSONFile(filePath string, result interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err!= nil {
		return err
	}
	return json.Unmarshal(data, result)
}

// WriteJSONFile writes a struct to a JSON file.
func WriteJSONFile(filePath string, result interface{}) error {
	data, err := json.MarshalIndent(result, "", "  ")
	if err!= nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, os.ModePerm)
}

// GetEnvironmentVariable returns the value of an environment variable.
func GetEnvironmentVariable(name string) string {
	return os.Getenv(name)
}

// IsEmptyString checks if a string is empty.
func IsEmptyString(s string) bool {
	return s == ""
}

// IsDirectory checks if a path is a directory.
func IsDirectory(path string) bool {
	fi, err := os.Stat(path)
	return err == nil && fi.IsDir()
}

// RemoveDirectory removes a directory and all its contents.
func RemoveDirectory(path string) error {
	return os.RemoveAll(path)
}

// GetDirectoryPath returns the directory path of a file.
func GetDirectoryPath(path string) string {
	return filepath.Dir(path)
}

// GetFileName returns the file name of a file path.
func GetFileName(path string) string {
	return filepath.Base(path)
}

// GetFileExtension returns the file extension of a file path.
func GetFileExtension(path string) string {
	return filepath.Ext(path)
}

// IsFile checks if a path is a file.
func IsFile(path string) bool {
	fi, err := os.Stat(path)
	return err == nil &&!fi.IsDir()
}

// GetFileSize returns the size of a file in bytes.
func GetFileSize(path string) (int64, error) {
	fi, err := os.Stat(path)
	if err!= nil {
		return 0, err
	}
	return fi.Size(), nil
}

// PrintError prints an error message to the standard error stream.
func PrintError(err error) {
	fmt.Fprintln(os.Stderr, err)
}