package types

import (
	"fmt"
)

// The time interval over which to aggregate the usage data.
type UsageAggregationInterval string

const (
	UsageAggregationIntervalHour       UsageAggregationInterval = "hour"
	UsageAggregationIntervalDay        UsageAggregationInterval = "day"
	UsageAggregationIntervalWeek       UsageAggregationInterval = "week"
	UsageAggregationIntervalMonth      UsageAggregationInterval = "month"
	UsageAggregationIntervalCumulative UsageAggregationInterval = "cumulative"
)

func NewUsageAggregationIntervalFromString(s string) (UsageAggregationInterval, error) {
	switch s {
	case "hour":
		return UsageAggregationIntervalHour, nil
	case "day":
		return UsageAggregationIntervalDay, nil
	case "week":
		return UsageAggregationIntervalWeek, nil
	case "month":
		return UsageAggregationIntervalMonth, nil
	case "cumulative":
		return UsageAggregationIntervalCumulative, nil
	}
	var t UsageAggregationInterval
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e UsageAggregationInterval) Ptr() *UsageAggregationInterval {
	return &e
}
