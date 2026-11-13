package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SetBookingSlugParams struct {
	extraProperties map[string]interface{}
}

func (s *SetBookingSlugParams) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SetBookingSlugParams) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SetBookingSlugParams to string"
	}
	return string(jsonData)
}

func (s *SetBookingSlugParams) UnmarshalJSON(data []byte) error {
	type unmarshaler SetBookingSlugParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SetBookingSlugParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "smb_tool_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SetBookingSlugParams) MarshalJSON() ([]byte, error) {
	type embed SetBookingSlugParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
