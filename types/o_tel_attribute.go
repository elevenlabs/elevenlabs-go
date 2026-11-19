package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OTelAttribute struct {
	Key             string                              `json:"key" xml:"key" required:"true"`
	Value           map[string]*OTelAttributeValueValue `json:"value" required:"true"`
	extraProperties map[string]interface{}
}

func (o *OTelAttribute) GetKey() (value string) {
	if o == nil {
		return
	}
	return o.Key
}

func (o *OTelAttribute) GetValue() (value map[string]*OTelAttributeValueValue) {
	if o == nil {
		return
	}
	return o.Value
}

func (o *OTelAttribute) SetKey(key string) {
	o.Key = key
}

func (o *OTelAttribute) SetValue(value map[string]*OTelAttributeValueValue) {
	o.Value = value
}

func (o *OTelAttribute) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OTelAttribute) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OTelAttribute to string"
	}
	return string(jsonData)
}

func (o *OTelAttribute) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, o); err != nil {
		return err
	}
	type alias OTelAttribute
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*o = OTelAttribute(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OTelAttribute) MarshalJSON() ([]byte, error) {
	type embed OTelAttribute
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
