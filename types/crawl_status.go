package types

import (
	"fmt"
)

type CrawlStatus string

const (
	CrawlStatusQueued     CrawlStatus = "queued"
	CrawlStatusProcessing CrawlStatus = "processing"
	CrawlStatusSucceeded  CrawlStatus = "succeeded"
	CrawlStatusFailed     CrawlStatus = "failed"
	CrawlStatusSkipped    CrawlStatus = "skipped"
	CrawlStatusCancelled  CrawlStatus = "cancelled"
)

func NewCrawlStatusFromString(s string) (CrawlStatus, error) {
	switch s {
	case "queued":
		return CrawlStatusQueued, nil
	case "processing":
		return CrawlStatusProcessing, nil
	case "succeeded":
		return CrawlStatusSucceeded, nil
	case "failed":
		return CrawlStatusFailed, nil
	case "skipped":
		return CrawlStatusSkipped, nil
	case "cancelled":
		return CrawlStatusCancelled, nil
	}
	var t CrawlStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CrawlStatus) Ptr() *CrawlStatus {
	return &e
}
