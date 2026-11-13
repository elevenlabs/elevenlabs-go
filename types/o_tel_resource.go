package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OTelResource struct {
	Attributes      []*OTelAttribute `json:"attributes" xml:"attributes" required:"true"`
	extraProperties map[string]interface{}
}

func (o *OTelResource) GetAttributes() (value []*OTelAttribute) {
	if o == nil {
		return
	}
	return o.Attributes
}

func (o *OTelResource) SetAttributes(attributes []*OTelAttribute) {
	o.Attributes = attributes
}

func (o *OTelResource) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OTelResource) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OTelResource to string"
	}
	return string(jsonData)
}

func (o *OTelResource) UnmarshalJSON(data []byte) error {
	type unmarshaler OTelResource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OTelResource(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OTelResource) MarshalJSON() ([]byte, error) {
	type embed OTelResource
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
