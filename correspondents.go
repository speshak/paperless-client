package paperless

type Correspondent struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	Slug          string `json:"slug"`
	DocumentCount int    `json:"document_count"`
}

type CorrespondentResult struct {
	SearchResult
	Results []Correspondent `json:"results"`
}
