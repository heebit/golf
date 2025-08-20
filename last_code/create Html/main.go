package main

import (
	"golf/utils"
	"strings"
)

func main() {
	rows, err := utils.ReadCSV("links_Japan_Hotels.csv")
	if err != nil {
		panic(err)
	}

	outputDir := "japan_hotels"

	for _, row := range rows {

		url := strings.ReplaceAll(row[0], "dk/", "")
		filename := extractFileNameFromURL(url)
		utils.DownloadHTMLAsFile(url, outputDir+"/"+row[1], filename)
	}
}

func extractFileNameFromURL(url string) string {
	parts := strings.Split(url, "/")

	if len(parts) == 0 {
		return "index"
	}
	fileName := parts[len(parts)-1]
	if fileName == "" {
		return "index"
	}
	return fileName
}
