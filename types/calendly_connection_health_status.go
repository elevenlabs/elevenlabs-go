package types

import (
	"fmt"
)

type CalendlyConnectionHealthStatus string

const (
	CalendlyConnectionHealthStatusHealthy     CalendlyConnectionHealthStatus = "healthy"
	CalendlyConnectionHealthStatusNeedsReauth CalendlyConnectionHealthStatus = "needs_reauth"
)

func NewCalendlyConnectionHealthStatusFromString(s string) (CalendlyConnectionHealthStatus, error) {
	switch s {
	case "healthy":
		return CalendlyConnectionHealthStatusHealthy, nil
	case "needs_reauth":
		return CalendlyConnectionHealthStatusNeedsReauth, nil
	}
	var t CalendlyConnectionHealthStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CalendlyConnectionHealthStatus) Ptr() *CalendlyConnectionHealthStatus {
	return &e
}
