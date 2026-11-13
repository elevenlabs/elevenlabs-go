package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OpenOutcome struct {
	extraProperties map[string]interface{}
}

func (o *OpenOutcome) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenOutcome) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenOutcome to string"
	}
	return string(jsonData)
}

func (o *OpenOutcome) UnmarshalJSON(data []byte) error {
	type unmarshaler OpenOutcome
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OpenOutcome(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o, "status")
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenOutcome) MarshalJSON() ([]byte, error) {
	type embed OpenOutcome
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
