package types

import (
	"fmt"
)

type ReceptionScrapingMode string

const (
	ReceptionScrapingModeSinglePage ReceptionScrapingMode = "single_page"
	ReceptionScrapingModeDeepScan   ReceptionScrapingMode = "deep_scan"
	ReceptionScrapingModeQuickScan  ReceptionScrapingMode = "quick_scan"
)

func NewReceptionScrapingModeFromString(s string) (ReceptionScrapingMode, error) {
	switch s {
	case "single_page":
		return ReceptionScrapingModeSinglePage, nil
	case "deep_scan":
		return ReceptionScrapingModeDeepScan, nil
	case "quick_scan":
		return ReceptionScrapingModeQuickScan, nil
	}
	var t ReceptionScrapingMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReceptionScrapingMode) Ptr() *ReceptionScrapingMode {
	return &e
}
