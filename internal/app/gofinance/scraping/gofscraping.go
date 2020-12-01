package scraping

import (
	"io"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeBody(input io.ReadCloser) string {
	var response string
	doc, err := goquery.NewDocumentFromReader(input)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		content := s.Contents().Text()
		if strings.Contains(content, "root.App.main") {
			splitted := strings.Split((strings.Split(content, "(this)")[0]), "root.App.main =")[1]
			response = strings.TrimSpace(strings.Replace(splitted, ";\n}", "", 1))
		}
	})
	return response
}
