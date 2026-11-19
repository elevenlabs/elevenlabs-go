package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Schedule a single instance of a group service.
//
// The session's duration is derived from the parent service so the assistant
// only has to pin start time, the (optional) instructor / room, and the
// location. Participants register separately via
// “register_for_group_session“.
type ScheduleGroupSessionParams struct {
	extraProperties map[string]interface{}
}

func (s *ScheduleGroupSessionParams) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScheduleGroupSessionParams) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScheduleGroupSessionParams to string"
	}
	return string(jsonData)
}

func (s *ScheduleGroupSessionParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ScheduleGroupSessionParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = ScheduleGroupSessionParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "smb_tool_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ScheduleGroupSessionParams) MarshalJSON() ([]byte, error) {
	type embed ScheduleGroupSessionParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
