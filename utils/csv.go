package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func WriteCSV(fileName string, rows [][]string) {

	outputFile, err := os.Create(fileName + ".csv")
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer outputFile.Close()

	csvWriter := csv.NewWriter(outputFile)
	defer csvWriter.Flush()

	for _, row := range rows {
		if err := csvWriter.Write(row); err != nil {
			log.Fatalf("Error writing row to CSV: %v", err)
		}
	}
	log.Printf("Data written to %s.csv successfully", fileName)
}

func ReadCSV(filename string) ([][]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	return csvReader.ReadAll()
}
