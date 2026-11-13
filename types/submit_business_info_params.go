package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Finalize the onboarding interview: create a text knowledge source and mark done.
type SubmitBusinessInfoParams struct {
	extraProperties map[string]interface{}
}

func (s *SubmitBusinessInfoParams) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubmitBusinessInfoParams) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubmitBusinessInfoParams to string"
	}
	return string(jsonData)
}

func (s *SubmitBusinessInfoParams) UnmarshalJSON(data []byte) error {
	type unmarshaler SubmitBusinessInfoParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SubmitBusinessInfoParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "smb_tool_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubmitBusinessInfoParams) MarshalJSON() ([]byte, error) {
	type embed SubmitBusinessInfoParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
