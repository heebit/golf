package utils

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func CreateFile(dir, fileName, extension string) (*os.File, error) {
	fullPath := filepath.Join(dir, fileName+extension)
	outputFile, err := os.Create(fullPath)

	if err != nil {
		return nil, err
	}

	log.Printf("File %s created successfully", fileName)
	return outputFile, nil
}

func MustCreateDirectory(dirName string) {
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating directory %s: %v", dirName, err)
	}
	log.Printf("Directory %s created successfully", dirName)
}

func GetDirectoryFilesWithExtension(pathDir, extension string) ([]string, error) {
	var filepaths []string

	err := filepath.WalkDir(pathDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || filepath.Ext(path) != extension {
			return nil
		}

		filepaths = append(filepaths, path)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error walking through directory %s: %w", pathDir, err)
	}
	if len(filepaths) == 0 {
		return nil, fmt.Errorf("no files with extension .%s found in directory %s", extension, pathDir)
	}
	return filepaths, nil
}

func GetOpenFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s: %w", filePath, err)
	}
	return file, nil
}

func DawnloadFile(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download %s: %s", url, resp.Status)
	}
	return resp, nil
}
