package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TimeRange struct {
	StartMs         int `json:"start_ms" xml:"start_ms" required:"true"`
	EndMs           int `json:"end_ms" xml:"end_ms" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TimeRange) GetStartMs() (value int) {
	if t == nil {
		return
	}
	return t.StartMs
}

func (t *TimeRange) GetEndMs() (value int) {
	if t == nil {
		return
	}
	return t.EndMs
}

func (t *TimeRange) SetStartMs(startMs int) {
	t.StartMs = startMs
}

func (t *TimeRange) SetEndMs(endMs int) {
	t.EndMs = endMs
}

func (t *TimeRange) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TimeRange) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TimeRange to string"
	}
	return string(jsonData)
}

func (t *TimeRange) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TimeRange
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TimeRange(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TimeRange) MarshalJSON() ([]byte, error) {
	type embed TimeRange
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
