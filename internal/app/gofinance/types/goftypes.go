package types

type Scraped struct {
	Context Context `json:"context"`
}

type Context struct {
	Dispatcher Dispatcher `json:"dispatcher"`
}

type Dispatcher struct {
	Stores Stores `json:"stores"`
}

type Stores struct {
	QSS QuoteSummaryStore `json:"quotesummarystore"`
}

type QuoteSummaryStore struct {
	//DefaultKeyStatistics
	// Details
	SummaryProfile      SummaryProfile `json:"summaryprofile`
	Symbol              string         `json:"symbol"`
	RecommendationTrend RecoTrend      `json:"recommendationTrend"`
}

type SummaryProfile struct {
	Zip                 string `json:"zip"`
	Sector              string `json:"sector"`
	FullTimeEmployees   int    `json:"fulltimeemployees`
	LongBusinessSummary string `json:"longbusinesssummary"`
	City                string `json:"city"`
	Phone               string `json:"phone"`
	State               string `json:"state"`
	Country             string `json:"country"`
	Website             string `json:"website"`
	MaxAge              int    `json:"maxage"`
	Address1            string `json:"address1"`
	Industry            string `json:"industry"`
}

type RecoTrend struct {
	Trend  []Trend `json:"trend"`
	MaxAge int     `json:"maxage"`
}

type Trend struct {
	Period     string `json:"period"`
	StrongBuy  int    `json:"strongbuy"`
	Buy        int    `json:"buy"`
	Hold       int    `json:"hold"`
	Sell       int    `json:"sell"`
	StrongSell int    `json:"strongsell"`
}
