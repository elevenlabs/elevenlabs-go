package types

import (
	"fmt"
)

// Single-word processing step shown in the UI status badge.
type ProcessingProgress string

const (
	ProcessingProgressConverting ProcessingProgress = "Converting"
	ProcessingProgressStoring    ProcessingProgress = "Storing"
	ProcessingProgressScraping   ProcessingProgress = "Scraping"
	ProcessingProgressCrawling   ProcessingProgress = "Crawling"
	ProcessingProgressExtracting ProcessingProgress = "Extracting"
)

func NewProcessingProgressFromString(s string) (ProcessingProgress, error) {
	switch s {
	case "Converting":
		return ProcessingProgressConverting, nil
	case "Storing":
		return ProcessingProgressStoring, nil
	case "Scraping":
		return ProcessingProgressScraping, nil
	case "Crawling":
		return ProcessingProgressCrawling, nil
	case "Extracting":
		return ProcessingProgressExtracting, nil
	}
	var t ProcessingProgress
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProcessingProgress) Ptr() *ProcessingProgress {
	return &e
}
