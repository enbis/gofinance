package scraping

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/enbis/gofinance/internal/app/gofinance/types"
	"gopkg.in/go-playground/assert.v1"
)

func TestScrapeBody(t *testing.T) {
	response, err := http.Get("https://finance.yahoo.com/quote/AAPL")
	if err != nil {
		fmt.Errorf("%s", err.Error())
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		r := ScrapeBody(response.Body)
		var in types.Scraped
		if err := json.Unmarshal([]byte(r), &in); err != nil {
			t.Errorf(err.Error())
		}
		assert.Equal(t, "AAPL", in.Context.Dispatcher.Stores.QSS.Symbol)
	}
}
