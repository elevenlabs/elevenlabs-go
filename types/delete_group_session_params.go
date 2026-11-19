package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Permanently remove a previously-cancelled group session.
//
// Group analogue of “delete_calendar_event“: cancel
// (“cancel_group_session_for_all“) is the soft, history-preserving step;
// this tool is the irreversible follow-up that drops the row from Mongo
// and the staff Google Calendar entirely. The backend rejects the call
// (422) if the session hasn't been cancelled yet, so the only safe path
// is cancel-then-delete.
type DeleteGroupSessionParams struct {
	extraProperties map[string]interface{}
}

func (d *DeleteGroupSessionParams) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeleteGroupSessionParams) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeleteGroupSessionParams to string"
	}
	return string(jsonData)
}

func (d *DeleteGroupSessionParams) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteGroupSessionParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteGroupSessionParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "smb_tool_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeleteGroupSessionParams) MarshalJSON() ([]byte, error) {
	type embed DeleteGroupSessionParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
