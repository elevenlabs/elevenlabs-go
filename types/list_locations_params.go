package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Params + Reception-proxied “run()“ for list_locations.
type ListLocationsParams struct {
	extraProperties map[string]interface{}
}

func (l *ListLocationsParams) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListLocationsParams) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListLocationsParams to string"
	}
	return string(jsonData)
}

func (l *ListLocationsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler ListLocationsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListLocationsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l, "smb_tool_type")
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListLocationsParams) MarshalJSON() ([]byte, error) {
	type embed ListLocationsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
