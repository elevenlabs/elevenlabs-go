package types

import (
	"fmt"
)

type CalComConnectionHealthStatus string

const (
	CalComConnectionHealthStatusHealthy     CalComConnectionHealthStatus = "healthy"
	CalComConnectionHealthStatusNeedsReauth CalComConnectionHealthStatus = "needs_reauth"
)

func NewCalComConnectionHealthStatusFromString(s string) (CalComConnectionHealthStatus, error) {
	switch s {
	case "healthy":
		return CalComConnectionHealthStatusHealthy, nil
	case "needs_reauth":
		return CalComConnectionHealthStatusNeedsReauth, nil
	}
	var t CalComConnectionHealthStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CalComConnectionHealthStatus) Ptr() *CalComConnectionHealthStatus {
	return &e
}
