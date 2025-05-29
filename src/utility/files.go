package utility

import (
	"os"
	"errors"
	"log"
	"path/filepath"
	"strings"
)

func ReadFile(filePath string) (string, error) {
	if filePath == "" {
		return "", errors.New("file path cannot be empty")
	}

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func WriteFile(filePath, content string) error {
	if filePath == "" {
		return errors.New("file path cannot be empty")
	}

	if content == "" {
		return errors.New("content cannot be empty")
	}

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		return err
	}

	log.Printf("File written successfully: %s", filePath)
	return nil
}

func ClearCache(cacheDir string) error {
	if cacheDir == "" {
		return errors.New("cache directory path cannot be empty")
	}

	err := os.RemoveAll(cacheDir)
	if err != nil {
		return err
	}

	err = os.MkdirAll(cacheDir, 0755)
	if err != nil {
		return err
	}

	log.Printf("Cache cleared at %s", cacheDir)
	return nil
}


func GetFiles(dir string, ignoredFiles []string) ([]string, error) {
	if dir == "" {
		return nil, errors.New("directory path cannot be empty")
	}

	files, err := filepath.Glob(filepath.Join(dir, "*.md"))
	if err != nil {
		return nil, err
	}

	for _, ignored := range ignoredFiles {
		ignored = strings.TrimSpace(ignored)
		for i := len(files) - 1; i >= 0; i-- {
			if strings.Contains(files[i], ignored) {
				files = append(files[:i], files[i+1:]...)
			}
		}
	}

	for i, file := range files {
		files[i] = strings.TrimSpace(file)
		if strings.HasPrefix(file, dir) {
			files[i] = strings.TrimPrefix(file, dir)
		} else {
			files[i] = strings.TrimPrefix(file, "/") // Handle absolute paths

		}
		files[i] = strings.TrimPrefix(files[i], "\\") // Handle Windows paths
		files[i] = strings.TrimPrefix(files[i], "./") // Handle relative paths
		log.Print("File found: ", files[i])
	}


	if len(files) == 0 {
		return nil, errors.New("no files found in the directory")
	}

	return files, nil
}