package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DateOverridePolicy struct {
	// First local date the override applies on.
	StartDate string `json:"start_date" xml:"start_date" required:"true"`
	// Last local date the override applies on, inclusive.
	EndDate string `json:"end_date" xml:"end_date" required:"true"`
	// Calling windows on every day from `start_date` to `end_date`, replacing the weekly ones. Empty means no calls on those days.
	Intervals       []*CallingIntervalPolicy `json:"intervals,omitempty" xml:"intervals,omitempty"`
	Name            *string                  `json:"name,omitempty" xml:"name,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DateOverridePolicy) GetStartDate() (value string) {
	if d == nil {
		return
	}
	return d.StartDate
}

func (d *DateOverridePolicy) GetEndDate() (value string) {
	if d == nil {
		return
	}
	return d.EndDate
}

func (d *DateOverridePolicy) GetIntervals() (value []*CallingIntervalPolicy) {
	if d == nil {
		return
	}
	return d.Intervals
}

func (d *DateOverridePolicy) GetName() (value *string) {
	if d == nil {
		return
	}
	return d.Name
}

func (d *DateOverridePolicy) SetStartDate(startDate string) {
	d.StartDate = startDate
}

func (d *DateOverridePolicy) SetEndDate(endDate string) {
	d.EndDate = endDate
}

func (d *DateOverridePolicy) SetIntervals(intervals []*CallingIntervalPolicy) {
	d.Intervals = intervals
}

func (d *DateOverridePolicy) SetName(name *string) {
	d.Name = name
}

func (d *DateOverridePolicy) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DateOverridePolicy) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DateOverridePolicy to string"
	}
	return string(jsonData)
}

func (d *DateOverridePolicy) UnmarshalJSON(data []byte) error {
	data = internal.ParseDateFields(data, "start_date", "end_date")
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DateOverridePolicy) MarshalJSON() ([]byte, error) {
	type embed DateOverridePolicy
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return internal.FormatDateFields("start_date", "end_date")(json.Marshal(marshaler))
}
