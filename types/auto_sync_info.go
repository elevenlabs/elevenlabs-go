package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AutoSyncInfo struct {
	// Minimum frequency (in days) at which the document is refreshed. The actual interval may be shorter, never longer.
	MinimumFrequencyDays *int `json:"minimum_frequency_days,omitempty" xml:"minimum_frequency_days,omitempty" min:"1" max:"180"`
	// Whether to remove the document if the URL becomes unavailable
	AutoRemove *bool `json:"auto_remove,omitempty" xml:"auto_remove,omitempty"`
	// Number of consecutive sync failures
	ConsecFailures *int `json:"consec_failures,omitempty" xml:"consec_failures,omitempty" min:"0"`
	// Unix timestamp for the next scheduled sync or None (in case of folders)
	NextRefreshBy   *int `json:"next_refresh_by,omitempty" xml:"next_refresh_by,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AutoSyncInfo) GetMinimumFrequencyDays() (value *int) {
	if a == nil {
		return
	}
	return a.MinimumFrequencyDays
}

func (a *AutoSyncInfo) GetAutoRemove() (value *bool) {
	if a == nil {
		return
	}
	return a.AutoRemove
}

func (a *AutoSyncInfo) GetConsecFailures() (value *int) {
	if a == nil {
		return
	}
	return a.ConsecFailures
}

func (a *AutoSyncInfo) GetNextRefreshBy() (value *int) {
	if a == nil {
		return
	}
	return a.NextRefreshBy
}

func (a *AutoSyncInfo) SetMinimumFrequencyDays(minimumFrequencyDays *int) {
	a.MinimumFrequencyDays = minimumFrequencyDays
}

func (a *AutoSyncInfo) SetAutoRemove(autoRemove *bool) {
	a.AutoRemove = autoRemove
}

func (a *AutoSyncInfo) SetConsecFailures(consecFailures *int) {
	a.ConsecFailures = consecFailures
}

func (a *AutoSyncInfo) SetNextRefreshBy(nextRefreshBy *int) {
	a.NextRefreshBy = nextRefreshBy
}

func (a *AutoSyncInfo) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AutoSyncInfo) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AutoSyncInfo to string"
	}
	return string(jsonData)
}

func (a *AutoSyncInfo) UnmarshalJSON(data []byte) error {
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

func (a *AutoSyncInfo) MarshalJSON() ([]byte, error) {
	type embed AutoSyncInfo
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
