package types

import (
	"fmt"
)

// Canonical representation of a knowledge base document's stored content.
//
// HTML is the legacy default; documents created before this field existed are
// interpreted as HTML.
type ContentFormat string

const (
	ContentFormatHtml     ContentFormat = "html"
	ContentFormatMarkdown ContentFormat = "markdown"
)

func NewContentFormatFromString(s string) (ContentFormat, error) {
	switch s {
	case "html":
		return ContentFormatHtml, nil
	case "markdown":
		return ContentFormatMarkdown, nil
	}
	var t ContentFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ContentFormat) Ptr() *ContentFormat {
	return &e
}
