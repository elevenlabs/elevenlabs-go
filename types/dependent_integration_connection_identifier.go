package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Identifier for an integration connection that depends on an auth connection
type DependentIntegrationConnectionIdentifier struct {
	Id              string `json:"id" xml:"id" required:"true"`
	Name            string `json:"name" xml:"name" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DependentIntegrationConnectionIdentifier) GetId() (value string) {
	if d == nil {
		return
	}
	return d.Id
}

func (d *DependentIntegrationConnectionIdentifier) GetName() (value string) {
	if d == nil {
		return
	}
	return d.Name
}

func (d *DependentIntegrationConnectionIdentifier) SetId(id string) {
	d.Id = id
}

func (d *DependentIntegrationConnectionIdentifier) SetName(name string) {
	d.Name = name
}

func (d *DependentIntegrationConnectionIdentifier) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DependentIntegrationConnectionIdentifier) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DependentIntegrationConnectionIdentifier to string"
	}
	return string(jsonData)
}

func (d *DependentIntegrationConnectionIdentifier) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DependentIntegrationConnectionIdentifier
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DependentIntegrationConnectionIdentifier(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DependentIntegrationConnectionIdentifier) MarshalJSON() ([]byte, error) {
	type embed DependentIntegrationConnectionIdentifier
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
