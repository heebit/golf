package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJSONFromFile(filename string, dest any) error {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", filename, err)
		return  err
	}

	if err := json.Unmarshal(fileContent, dest); err != nil {
		fmt.Printf("Error parsing %s: %v\n", filename, err)
		return err
	}
	return nil
}
