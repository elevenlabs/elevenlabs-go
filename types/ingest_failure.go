package types

import (
	"fmt"
)

// Why processing a knowledge source failed.
type IngestFailure string

const (
	IngestFailureBlockedUrl      IngestFailure = "blocked_url"
	IngestFailureDnsFailure      IngestFailure = "dns_failure"
	IngestFailureHostUnreachable IngestFailure = "host_unreachable"
	IngestFailureAccessRefused   IngestFailure = "access_refused"
	IngestFailureCrawlTimeout    IngestFailure = "crawl_timeout"
	IngestFailureNoPages         IngestFailure = "no_pages"
	IngestFailureNoContent       IngestFailure = "no_content"
	IngestFailureUpstreamTimeout IngestFailure = "upstream_timeout"
	IngestFailureUpstreamError   IngestFailure = "upstream_error"
	IngestFailureUnknown         IngestFailure = "unknown"
)

func NewIngestFailureFromString(s string) (IngestFailure, error) {
	switch s {
	case "blocked_url":
		return IngestFailureBlockedUrl, nil
	case "dns_failure":
		return IngestFailureDnsFailure, nil
	case "host_unreachable":
		return IngestFailureHostUnreachable, nil
	case "access_refused":
		return IngestFailureAccessRefused, nil
	case "crawl_timeout":
		return IngestFailureCrawlTimeout, nil
	case "no_pages":
		return IngestFailureNoPages, nil
	case "no_content":
		return IngestFailureNoContent, nil
	case "upstream_timeout":
		return IngestFailureUpstreamTimeout, nil
	case "upstream_error":
		return IngestFailureUpstreamError, nil
	case "unknown":
		return IngestFailureUnknown, nil
	}
	var t IngestFailure
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e IngestFailure) Ptr() *IngestFailure {
	return &e
}
