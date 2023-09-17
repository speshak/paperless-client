package paperless

import (
	"fmt"
	"time"
)

// Paperless Document
type Document struct {
	Id                    int       `json:"id"`
	Correspondent         int       `json:"correspondent"`
	Document_type         int       `json:"document_type"`
	Storage_path          string    `json:"storage_path"`
	Title                 string    `json:"title"`
	Content               string    `json:"content"`
	Tags                  []int     `json:"tags"`
	Created               time.Time `json:"created"`
	Created_date          string    `json:"created_date"`
	Modified              time.Time `json:"modified"`
	Added                 time.Time `json:"added"`
	Archive_serial_number string    `json:"archive_serial_number"`
	Original_file_name    string    `json:"original_file_name"`
	Archived_file_name    string    `json:"archived_file_name"`
	Owner                 int       `json:"owner"`
	User_can_change       bool      `json:"user_can_change"`
	Notes                 []string  `json:"notes"`

	client *PaperlessClient
}

// Result of a document search query
type DocumentResult struct {
	SearchResult
	Results []Document `json:"results"`
}

// Update the document
func (d *Document) Update() error {
	_, err := d.client.client.R().
		SetBody(d).
		Put(fmt.Sprintf("/api/documents/%d/", d.Id))
	return err
}

// Get the download URL for the document
func (d *Document) DownloadUrl() string {
	return fmt.Sprintf("%s/api/documents/%d/download/", d.client.url, d.Id)
}

// Get the preview URL for the document
func (d *Document) PreviewUrl() string {
	return fmt.Sprintf("%s/api/documents/%d/preview/", d.client.url, d.Id)
}

// Get the thumbnail URL for the document
func (d *Document) ThumbnailUrl() string {
	return fmt.Sprintf("%s/api/documents/%d/thumbnail/", d.client.url, d.Id)
}

// Remove a tag from the document
func (d *Document) RemoveTag(tag int) error {
	// Paperless API doesn't allow directly removing a tag, instead we have to
	// update the entire document. This fakes it by removing the tag from the
	// existing list of tags and then updating the document.
	var newTags []int
	for _, t := range d.Tags {
		if t != tag {
			newTags = append(newTags, t)
		}
	}
	d.Tags = newTags
	return d.Update()
}

// Add a tag to the document
func (d *Document) AddTag(tag int) error {
	// Paperless API doesn't allow directly adding a tag, instead we have to
	// update the entire document. This fakes it by adding the tag to the
	// existing list of tags and then updating the document.
	d.Tags = append(d.Tags, tag)
	return d.Update()
}
