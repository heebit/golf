package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func DownloadHTMLAsFile(url, dirName, fileName string) error {
	MustCreateDirectory(dirName)

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download %s: %s", url, resp.Status)
	}

	outputFile, err := CreateFile(dirName, fileName, ".html")
	if err != nil {
		return err
	}

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
	