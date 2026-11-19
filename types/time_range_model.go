package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Represents a time-of-day range (e.g., 09:00-17:00) commonly used for
// working hours, available hours, or recurring schedules.
type TimeRangeModel struct {
	// Start hour (0-23)
	StartHour int `json:"start_hour" xml:"start_hour" required:"true" min:"0"`
	// Start minute (0-59)
	StartMinute int `json:"start_minute" xml:"start_minute" required:"true" min:"0"`
	// End hour (0-24, where 24:00 means midnight end-of-day)
	EndHour int `json:"end_hour" xml:"end_hour" required:"true" min:"0" max:"24"`
	// End minute (0-59)
	EndMinute       int `json:"end_minute" xml:"end_minute" required:"true" min:"0"`
	extraProperties map[string]interface{}
}

func (t *TimeRangeModel) GetStartHour() (value int) {
	if t == nil {
		return
	}
	return t.StartHour
}

func (t *TimeRangeModel) GetStartMinute() (value int) {
	if t == nil {
		return
	}
	return t.StartMinute
}

func (t *TimeRangeModel) GetEndHour() (value int) {
	if t == nil {
		return
	}
	return t.EndHour
}

func (t *TimeRangeModel) GetEndMinute() (value int) {
	if t == nil {
		return
	}
	return t.EndMinute
}

func (t *TimeRangeModel) SetStartHour(startHour int) {
	t.StartHour = startHour
}

func (t *TimeRangeModel) SetStartMinute(startMinute int) {
	t.StartMinute = startMinute
}

func (t *TimeRangeModel) SetEndHour(endHour int) {
	t.EndHour = endHour
}

func (t *TimeRangeModel) SetEndMinute(endMinute int) {
	t.EndMinute = endMinute
}

func (t *TimeRangeModel) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TimeRangeModel) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TimeRangeModel to string"
	}
	return string(jsonData)
}

func (t *TimeRangeModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TimeRangeModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TimeRangeModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TimeRangeModel) MarshalJSON() ([]byte, error) {
	type embed TimeRangeModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
