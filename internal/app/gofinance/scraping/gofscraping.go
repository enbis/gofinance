package scraping

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	goftype "github.com/enbis/gofinance/internal/app/gofinance/types"
	colly "github.com/gocolly/colly/v2"
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

func ScrapeTable(url string) goftype.Holders {
	holders := goftype.Holders{}
	majors := []goftype.Major{}
	tops := []goftype.Top{}

	c := colly.NewCollector()

	c.OnHTML("table tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			buff := []string{}
			index := 0
			row.ForEach("td", func(_ int, el *colly.HTMLElement) {
				index = el.Index
				buff = append(buff, el.Text)
			})
			fmt.Println("index ", index)
			fmt.Println("buff ", buff)
			if index > 1 && len(buff) == 5 {
				tops = append(tops, goftype.Top{Holder: buff[0], Shares: buff[1], Data: buff[2], Out: buff[3], Value: buff[4]})
			} else {
				majors = append(majors, goftype.Major{Percentage: buff[0], Detail: buff[1]})
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.Visit(url)

	holders.Majors = majors
	holders.Tops = tops

	return holders
}

/*


// Before making a request print "Visiting ..."

err := c.Visit("https://gianarb.it/conferences.html")
if err != nil {
	println(err)
}
*/
