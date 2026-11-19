package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OTelScope struct {
	Name            string `json:"name" xml:"name" required:"true"`
	Version         string `json:"version" xml:"version" required:"true"`
	extraProperties map[string]interface{}
}

func (o *OTelScope) GetName() (value string) {
	if o == nil {
		return
	}
	return o.Name
}

func (o *OTelScope) GetVersion() (value string) {
	if o == nil {
		return
	}
	return o.Version
}

func (o *OTelScope) SetName(name string) {
	o.Name = name
}

func (o *OTelScope) SetVersion(version string) {
	o.Version = version
}

func (o *OTelScope) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OTelScope) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OTelScope to string"
	}
	return string(jsonData)
}

func (o *OTelScope) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, o); err != nil {
		return err
	}
	type alias OTelScope
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*o = OTelScope(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OTelScope) MarshalJSON() ([]byte, error) {
	type embed OTelScope
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
