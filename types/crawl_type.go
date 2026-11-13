package types

import (
	"fmt"
)

type CrawlType string

const (
	CrawlTypeDiscovery CrawlType = "discovery"
	CrawlTypeSitemap   CrawlType = "sitemap"
)

func NewCrawlTypeFromString(s string) (CrawlType, error) {
	switch s {
	case "discovery":
		return CrawlTypeDiscovery, nil
	case "sitemap":
		return CrawlTypeSitemap, nil
	}
	var t CrawlType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CrawlType) Ptr() *CrawlType {
	return &e
}
