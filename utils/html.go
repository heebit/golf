package utils

import (
	"fmt"

	"io"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func DownloadHTMLAsFile(url, dirName, fileName string) error {
	MustCreateDirectory(dirName)

	resp, err := DawnloadFile(url)

	if err != nil {
		return fmt.Errorf("error downloading file from %s: %w", url, err)
	}
	outputFile, err := CreateFile(dirName, fileName, ".html")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	defer outputFile.Close()
	_, err = io.Copy(outputFile, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func GetHtmlContent(file *os.File) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file %s: %w", file.Name(), err)
	}
	return doc, nil
}
