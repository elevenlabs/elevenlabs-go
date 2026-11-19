package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SendCustomEmailParams struct {
	extraProperties map[string]interface{}
}

func (s *SendCustomEmailParams) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SendCustomEmailParams) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SendCustomEmailParams to string"
	}
	return string(jsonData)
}

func (s *SendCustomEmailParams) UnmarshalJSON(data []byte) error {
	type unmarshaler SendCustomEmailParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SendCustomEmailParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "smb_tool_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SendCustomEmailParams) MarshalJSON() ([]byte, error) {
	type embed SendCustomEmailParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
