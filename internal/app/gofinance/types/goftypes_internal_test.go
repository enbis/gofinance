package types

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var in Scraped
	input := `{
		"context":{
			"dispatcher":{
				"stores":{
					"QuoteSummaryStore":{
						"symbol":"MSFT"
					}
				}
			}
		}
	}`

	if err := json.Unmarshal([]byte(input), &in); err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println("Symbol ", in.Context.Dispatcher.Stores.QSS.Symbol)
}
