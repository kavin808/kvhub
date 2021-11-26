package tool

import (
	"log"
	"os"
	"path/filepath"
)

// InitLog var log
func InitLog(dirPath string, filename string) (string, error) {
	if err := createLogDir(dirPath); err != nil {
		return "", err
	}
	path := filepath.Join(dirPath, filename)
	location, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
	if err == nil {
		log.SetOutput(location)
	}
	return path, err
}

func createLogDir(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.Mkdir(dirPath, 0755)
	}
	return nil
}
