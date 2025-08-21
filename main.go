package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golf/utils"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
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
	rootDir := "/Users/ivanivanov/Desktop/go/golf/japan_hotels"
	outDir := "/Users/ivanivanov/Desktop/go/golf/jsonJpanHotels"
	data, err := utils.GetDirectoryFilesWithExtension(rootDir, ".html")
	if err != nil {
		panic(err)
	}
	utils.MustCreateDirectory(outDir)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	for _, file := range data {
		content, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
		if err != nil {
			panic(err)
		}

		relPath, err := filepath.Rel(rootDir, file)
		if err != nil {
			panic(err)
		}

		targetDir := filepath.Join(outDir, filepath.Dir(relPath))
		utils.MustCreateDirectory(targetDir)

		doc.Find(".data-sheet").Each(func(_ int, s *goquery.Selection) {
			title := strings.TrimSpace(s.Find(".data-sheet__title").Text())
			address := strings.TrimSpace(
				s.Find(".data-sheet__detail-info .data-sheet__block").Eq(0).
					Find(".data-sheet__block--text").First().Text(),
			)
			rooms := strings.TrimSpace(
				s.Find(".data-sheet__detail-info .data-sheet__block").Eq(1).
					Find(".data-sheet__block--text").First().Text(),
			)
			style := strings.TrimSpace(
				s.Find(".data-sheet__detail-info .data-sheet__block").Eq(1).
					Find(".data-sheet__block--text").Last().Text(),
			)

			var images []string
			doc.Find("#js-owl-carousel-gallery img").Each(func(i int, s *goquery.Selection) {
				if src, exists := s.Attr("src"); exists {
					images = append(images, src)
				}
			})

			price := ""
			err := chromedp.Run(ctx,
				chromedp.Navigate("file://"+file),
				chromedp.Text("h6.hotelpage__booking--title.text__width--bold", &price, chromedp.NodeVisible, chromedp.ByQuery),
			)
			if err != nil {
				log.Println("Не удалось получить цену:", err)
				price = ""
			}

			title = utils.SanitizeFileName(title)

			hotel := Hotel{
				Title:   title,
				Address: address,
				Rooms:   rooms,
				Style:   style,
				Price:   price,
				Images:  images,
			}

			jsonFile, err := utils.CreateFile(targetDir, title, ".json")
			if err != nil {
				panic(err)
			}
			defer jsonFile.Close()

			encoder := json.NewEncoder(jsonFile)
			encoder.SetIndent("", "  ")
			encoder.SetEscapeHTML(false) 
			if err := encoder.Encode(hotel); err != nil {
				panic(err)
			}

			
		})
	}

}
