package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Every location, inactive ones included, with every owner-editable field.
type OwnerListLocationsParams struct {
	extraProperties map[string]interface{}
}

func (o *OwnerListLocationsParams) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OwnerListLocationsParams) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OwnerListLocationsParams to string"
	}
	return string(jsonData)
}

func (o *OwnerListLocationsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler OwnerListLocationsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OwnerListLocationsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o, "smb_tool_type")
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OwnerListLocationsParams) MarshalJSON() ([]byte, error) {
	type embed OwnerListLocationsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
