package paperless

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

// PaperlessClient is the main struct for the paperless client
type PaperlessClient struct {
	url    string
	token  string
	client *resty.Client
}

// New creates a new PaperlessClient
func New(url string, token string) *PaperlessClient {
	var client = resty.New().
		//SetDebug(true).
		SetHostURL(url).
		SetHeader("Accept", "application/json; version=3").
		SetAuthScheme("Token").
		SetAuthToken(token)

	return &PaperlessClient{
		url:    url,
		token:  token,
		client: client,
	}
}

func (p *PaperlessClient) GetDocument(id int) (Document, error) {
	var document Document
	_, err := p.client.R().
		SetResult(&document).
		Get(fmt.Sprintf("/api/documents/%d/", id))

	// Preserve a reference to the client
	document.client = p
	return document, err
}

func (p *PaperlessClient) SearchDocuments(query SearchQuery) ([]Document, error) {
	var documentResult DocumentResult
	_, err := p.client.R().
		SetResult(&documentResult).
		Get(fmt.Sprintf("/api/documents/?%s", query.Encode()))

	// Preserve a reference to the client
	for i := range documentResult.Results {
		documentResult.Results[i].client = p
	}
	return documentResult.Results, err
}

func (p *PaperlessClient) GetDocumentsByTagId(tag int) ([]Document, error) {
	return p.SearchDocuments(Query("tags__id", fmt.Sprintf("%d", tag)))
}

func (p *PaperlessClient) GetDocumentsByTag(tag string) ([]Document, error) {
	return p.SearchDocuments(Query("tags__name__iexact", tag))
}
