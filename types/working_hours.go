package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type WorkingHours struct {
	// Open time ranges per day of the week, keyed 0 (Monday) to 6 (Sunday). A day with no ranges is closed.
	Hours map[string][]*TimeRangeModel `json:"hours,omitempty"`
	// Timezone of these hours. Null uses the business's timezone.
	Timezone        *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	extraProperties map[string]interface{}
}

func (w *WorkingHours) GetHours() (value map[string][]*TimeRangeModel) {
	if w == nil {
		return
	}
	return w.Hours
}

func (w *WorkingHours) GetTimezone() (value *string) {
	if w == nil {
		return
	}
	return w.Timezone
}

func (w *WorkingHours) SetHours(hours map[string][]*TimeRangeModel) {
	w.Hours = hours
}

func (w *WorkingHours) SetTimezone(timezone *string) {
	w.Timezone = timezone
}

func (w *WorkingHours) GetExtraProperties() map[string]interface{} {
	if w == nil {
		return nil
	}
	return w.extraProperties
}

func (w WorkingHours) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WorkingHours to string"
	}
	return string(jsonData)
}

func (w *WorkingHours) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, w); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *w)
	if err != nil {
		return err
	}
	w.extraProperties = extraProperties
	return nil
}

func (w *WorkingHours) MarshalJSON() ([]byte, error) {
	type embed WorkingHours
	var marshaler = struct {
		embed
	}{
		embed: embed(*w),
	}
	return json.Marshal(marshaler)
}
