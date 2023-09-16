package paperless

import (
	"testing"
	//"fmt"
	"github.com/stretchr/testify/assert"
	"os"
)

func TestGetDocument(t *testing.T) {
	client := New(os.Getenv("PAPERLESS_URL"), os.Getenv("PAPERLESS_TOKEN"))
	doc, err := client.GetDocument(471)

	if err != nil {
		t.Error(err)
	}

	//fmt.Printf("%#v\n", doc)
	assert.Equal(t, 471, doc.Id)
}

func TestGetDocumentsByTagId(t *testing.T) {
	client := New(os.Getenv("PAPERLESS_URL"), os.Getenv("PAPERLESS_TOKEN"))
	docs, err := client.GetDocumentsByTagId(12)

	if err != nil {
		t.Error(err)
	}

	//fmt.Printf("%#v\n", docs)
	assert.Equal(t, 2, len(docs))
}

func TestGetDocumentsByTag(t *testing.T) {
	client := New(os.Getenv("PAPERLESS_URL"), os.Getenv("PAPERLESS_TOKEN"))
	docs, err := client.GetDocumentsByTag("Appliance")

	if err != nil {
		t.Error(err)
	}

	//fmt.Printf("%#v\n", docs)
	assert.Equal(t, 2, len(docs))
}
