package main

import (
	"fmt"
	"golf/utils"
	"io"
	"path/filepath"
	"strings"
)

type Hotel struct {
	Title   string   `json:"title"`
	Address string   `json:"address"`
	Rooms   string   `json:"rooms"`
	Style   string   `json:"style"`
	Images  []string `json:"images"`
	Price   string   `json:"price"`
}

func main() {

	imgDir := "/Users/ivanivanov/Desktop/go/golf/imagesOfHotels"
	paths, err := utils.GetDirectoryFilesWithExtension("/Users/ivanivanov/Desktop/go/golf/jsonJpanHotels", ".json")
	if err != nil {
		panic(err)
	}
	utils.MustCreateDirectory(imgDir)

	for _, path := range paths {

		city := getNameFromPath(path)

		var hotel Hotel
		err := utils.ReadJSONFromFile(path, &hotel)
		if err != nil {
			fmt.Printf("Ошибка чтения %s: %v\n", path, err)
			continue
		}

		title := utils.SanitizeFileName(hotel.Title)
		hotelDir := filepath.Join(imgDir, city, title)
		utils.MustCreateDirectory(hotelDir)

		for idx, image := range hotel.Images {
			resp, err := utils.DawnloadFile(image)
			if err != nil {
				fmt.Printf("Ошибка скачивания %s: %v\n", image, err)
				continue
			}

			fileName := fmt.Sprintf("%s_%d", title, idx+1)

			file, err := utils.CreateFile(hotelDir, fileName, ".jpg")
			if err != nil {
				resp.Body.Close()
				fmt.Printf("Ошибка создания файла %s: %v\n", fileName, err)
				continue
			}

			_, err = io.Copy(file, resp.Body)
			resp.Body.Close()
			file.Close()
			if err != nil {
				fmt.Printf("Ошибка сохранения %s: %v\n", fileName, err)
			}

		}
	}
}

func getNameFromPath(path string) string {
	cityName := strings.Split(path, "/")
	return cityName[len(cityName)-2]
}
