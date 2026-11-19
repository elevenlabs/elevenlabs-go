package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AlertingMonitorConfig struct {
	// Whether this monitor is enabled and can notify
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// Failure rate threshold at which this monitor can notify.
	Threshold *float64 `json:"threshold,omitempty" xml:"threshold,omitempty" min:"0" max:"1"`
	// Relative increase over the trailing baseline at which this monitor can notify (0.2 = 20% above baseline, 0 = any failure). For conversation_volume_drop it is the share of the usual volume that must be missing, between 0.1 and 0.9 (0.5 = half).
	RelativeIncreaseThreshold *float64 `json:"relative_increase_threshold,omitempty" xml:"relative_increase_threshold,omitempty" min:"0" max:"10"`
	// Minimum failures in the window before this monitor can fire.
	MinFailureCount *int `json:"min_failure_count,omitempty" xml:"min_failure_count,omitempty" min:"1"`
	// Minimum trailing buckets with traffic before spike detection can fire.
	MinHistoryBucketCount *int `json:"min_history_bucket_count,omitempty" xml:"min_history_bucket_count,omitempty" min:"1"`
	// Minimum samples in the window before this monitor can fire.
	MinSampleCount *int `json:"min_sample_count,omitempty" xml:"min_sample_count,omitempty" min:"1"`
	// How many suspect buckets within the lookback window are required to promote a suspect to an alert.
	SuspectTriggerThreshold *int `json:"suspect_trigger_threshold,omitempty" xml:"suspect_trigger_threshold,omitempty" min:"1" max:"10"`
	// How many minutes an alert can stay inactive before it is auto-resolved.
	AutoResolveAfterInactiveMinutes *int `json:"auto_resolve_after_inactive_minutes,omitempty" xml:"auto_resolve_after_inactive_minutes,omitempty" min:"1"`
	extraProperties                 map[string]interface{}
}

func (a *AlertingMonitorConfig) GetEnabled() (value *bool) {
	if a == nil {
		return
	}
	return a.Enabled
}

func (a *AlertingMonitorConfig) GetThreshold() (value *float64) {
	if a == nil {
		return
	}
	return a.Threshold
}

func (a *AlertingMonitorConfig) GetRelativeIncreaseThreshold() (value *float64) {
	if a == nil {
		return
	}
	return a.RelativeIncreaseThreshold
}

func (a *AlertingMonitorConfig) GetMinFailureCount() (value *int) {
	if a == nil {
		return
	}
	return a.MinFailureCount
}

func (a *AlertingMonitorConfig) GetMinHistoryBucketCount() (value *int) {
	if a == nil {
		return
	}
	return a.MinHistoryBucketCount
}

func (a *AlertingMonitorConfig) GetMinSampleCount() (value *int) {
	if a == nil {
		return
	}
	return a.MinSampleCount
}

func (a *AlertingMonitorConfig) GetSuspectTriggerThreshold() (value *int) {
	if a == nil {
		return
	}
	return a.SuspectTriggerThreshold
}

func (a *AlertingMonitorConfig) GetAutoResolveAfterInactiveMinutes() (value *int) {
	if a == nil {
		return
	}
	return a.AutoResolveAfterInactiveMinutes
}

func (a *AlertingMonitorConfig) SetEnabled(enabled *bool) {
	a.Enabled = enabled
}

func (a *AlertingMonitorConfig) SetThreshold(threshold *float64) {
	a.Threshold = threshold
}

func (a *AlertingMonitorConfig) SetRelativeIncreaseThreshold(relativeIncreaseThreshold *float64) {
	a.RelativeIncreaseThreshold = relativeIncreaseThreshold
}

func (a *AlertingMonitorConfig) SetMinFailureCount(minFailureCount *int) {
	a.MinFailureCount = minFailureCount
}

func (a *AlertingMonitorConfig) SetMinHistoryBucketCount(minHistoryBucketCount *int) {
	a.MinHistoryBucketCount = minHistoryBucketCount
}

func (a *AlertingMonitorConfig) SetMinSampleCount(minSampleCount *int) {
	a.MinSampleCount = minSampleCount
}

func (a *AlertingMonitorConfig) SetSuspectTriggerThreshold(suspectTriggerThreshold *int) {
	a.SuspectTriggerThreshold = suspectTriggerThreshold
}

func (a *AlertingMonitorConfig) SetAutoResolveAfterInactiveMinutes(autoResolveAfterInactiveMinutes *int) {
	a.AutoResolveAfterInactiveMinutes = autoResolveAfterInactiveMinutes
}

func (a *AlertingMonitorConfig) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AlertingMonitorConfig) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AlertingMonitorConfig to string"
	}
	return string(jsonData)
}

func (a *AlertingMonitorConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AlertingMonitorConfig) MarshalJSON() ([]byte, error) {
	type embed AlertingMonitorConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
