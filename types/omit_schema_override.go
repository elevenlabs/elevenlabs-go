package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OmitSchemaOverride struct {
	extraProperties map[string]interface{}
}

func (o *OmitSchemaOverride) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OmitSchemaOverride) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OmitSchemaOverride to string"
	}
	return string(jsonData)
}

func (o *OmitSchemaOverride) UnmarshalJSON(data []byte) error {
	type unmarshaler OmitSchemaOverride
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OmitSchemaOverride(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o, "source")
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OmitSchemaOverride) MarshalJSON() ([]byte, error) {
	type embed OmitSchemaOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
