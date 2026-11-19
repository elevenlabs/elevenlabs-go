package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ListCalendarEventsParams struct {
	IncludeCancelled *bool `json:"include_cancelled,omitempty" xml:"include_cancelled,omitempty"`
	extraProperties  map[string]interface{}
}

func (l *ListCalendarEventsParams) GetIncludeCancelled() (value *bool) {
	if l == nil {
		return
	}
	return l.IncludeCancelled
}

func (l *ListCalendarEventsParams) SetIncludeCancelled(includeCancelled *bool) {
	l.IncludeCancelled = includeCancelled
}

func (l *ListCalendarEventsParams) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListCalendarEventsParams) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListCalendarEventsParams to string"
	}
	return string(jsonData)
}

func (l *ListCalendarEventsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListCalendarEventsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListCalendarEventsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "smb_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListCalendarEventsParams) MarshalJSON() ([]byte, error) {
	type embed ListCalendarEventsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
