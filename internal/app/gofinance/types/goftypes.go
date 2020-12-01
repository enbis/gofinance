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
	Symbol string `json:"symbol"`
}
