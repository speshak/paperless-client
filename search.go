/**
 * Types & Functions for interacting with the paperless search interface
 */
package paperless

// SearchResult is the result of a search query
// This will be overridden by the specific search result
type SearchResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	All      []int  `json:"all"`
}

type SearchTerm struct {
	Key   string
	Value string
}

type SearchQuery []SearchTerm

// Create URL encoded string from search query
func (q SearchQuery) Encode() string {
	var result string
	for _, term := range q {
		result += term.Key + "=" + term.Value + "&"
	}
	return result
}

func (q SearchQuery) Add(term SearchTerm) SearchQuery {
	return append(q, term)
}

func Query(key string, value string) SearchQuery {
	return SearchQuery{SearchTerm{key, value}}
}
