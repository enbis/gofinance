package foundamentals

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/enbis/gofinance/internal/app/gofinance/scraping"
	"github.com/enbis/gofinance/internal/app/gofinance/types"
)

const (
	prefixUrl string = "https://finance.yahoo.com/quote"
)

func Get(t string) string {
	var response string
	var jsonResponse types.Scraped
	url := fmt.Sprintf("%s/%s", prefixUrl, strings.ToUpper(t))
	fmt.Printf("%v\n", url)
	res, err := http.Get(url)
	if err != nil {
		//TODO error handling
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		//TODO error handling
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	response = scraping.ScrapeBody(res.Body)
	if err = json.Unmarshal([]byte(response), &jsonResponse); err != nil {
		log.Fatalf("unable to unmarshal %v ", err)
	}

	fmt.Printf("Symbol %s", jsonResponse.Context.Dispatcher.Stores.QSS.Symbol)

	return response
}
