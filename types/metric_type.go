package types

import (
	"fmt"
)

type MetricType string

const (
	MetricTypeCredits            MetricType = "credits"
	MetricTypeTtsCharacters      MetricType = "tts_characters"
	MetricTypeMinutesUsed        MetricType = "minutes_used"
	MetricTypeRequestCount       MetricType = "request_count"
	MetricTypeTtfbAvg            MetricType = "ttfb_avg"
	MetricTypeTtfbP95            MetricType = "ttfb_p95"
	MetricTypeFiatUnitsSpent     MetricType = "fiat_units_spent"
	MetricTypeConcurrency        MetricType = "concurrency"
	MetricTypeConcurrencyAverage MetricType = "concurrency_average"
)

func NewMetricTypeFromString(s string) (MetricType, error) {
	switch s {
	case "credits":
		return MetricTypeCredits, nil
	case "tts_characters":
		return MetricTypeTtsCharacters, nil
	case "minutes_used":
		return MetricTypeMinutesUsed, nil
	case "request_count":
		return MetricTypeRequestCount, nil
	case "ttfb_avg":
		return MetricTypeTtfbAvg, nil
	case "ttfb_p95":
		return MetricTypeTtfbP95, nil
	case "fiat_units_spent":
		return MetricTypeFiatUnitsSpent, nil
	case "concurrency":
		return MetricTypeConcurrency, nil
	case "concurrency_average":
		return MetricTypeConcurrencyAverage, nil
	}
	var t MetricType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MetricType) Ptr() *MetricType {
	return &e
}
