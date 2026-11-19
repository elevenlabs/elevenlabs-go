package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Permanently remove a previously-cancelled calendar event.
//
// This delete tool is the irreversible follow-up to cancel_calendar_event.
// The backend rejects the call (422) if the event hasn't been
// cancelled yet, so the only safe path is cancel-then-delete.
type DeleteCalendarEventParams struct {
	extraProperties map[string]interface{}
}

func (d *DeleteCalendarEventParams) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeleteCalendarEventParams) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeleteCalendarEventParams to string"
	}
	return string(jsonData)
}

func (d *DeleteCalendarEventParams) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteCalendarEventParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteCalendarEventParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "smb_tool_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeleteCalendarEventParams) MarshalJSON() ([]byte, error) {
	type embed DeleteCalendarEventParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
