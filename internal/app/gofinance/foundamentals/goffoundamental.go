package foundamentals

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const (
	scrapeUrl string = "https://finance.yahoo.com/quote"
)

func Get(t string) string {
	var response string
	url := fmt.Sprintf("%s/%s", scrapeUrl, strings.ToUpper(t))
	fmt.Printf("%v\n", url)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// html scrapping
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ready!!!")
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//json_str = html.split('root.App.main =')[1].split(
		//'(this)')[0].split(';\n}')[0].strip()
		content := s.Contents().Text()
		if strings.Contains(content, "root.App.main") {
			xresponse := strings.Split((strings.Split(content, "(this)")[0]), "root.App.main =")[1]
			response = strings.Replace(xresponse, ";\n}", "", 1)
		}
		//p := s.Has("root.App.main").Text()
	})
	return response
	// .Each(func(index int, item *goquery.Selection) {
	// 	title := item.Text()
	// 	fmt.Println("Title ", title)
	// })
}
