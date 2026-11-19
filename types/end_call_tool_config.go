package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type EndCallToolConfig struct {
	extraProperties map[string]interface{}
}

func (e *EndCallToolConfig) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EndCallToolConfig) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EndCallToolConfig to string"
	}
	return string(jsonData)
}

func (e *EndCallToolConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler EndCallToolConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EndCallToolConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "system_tool_type")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EndCallToolConfig) MarshalJSON() ([]byte, error) {
	type embed EndCallToolConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
