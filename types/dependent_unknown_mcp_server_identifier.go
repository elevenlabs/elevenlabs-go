package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DependentUnknownMcpServerIdentifier struct {
	Id              string `json:"id" xml:"id" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DependentUnknownMcpServerIdentifier) GetId() (value string) {
	if d == nil {
		return
	}
	return d.Id
}

func (d *DependentUnknownMcpServerIdentifier) SetId(id string) {
	d.Id = id
}

func (d *DependentUnknownMcpServerIdentifier) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DependentUnknownMcpServerIdentifier) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DependentUnknownMcpServerIdentifier to string"
	}
	return string(jsonData)
}

func (d *DependentUnknownMcpServerIdentifier) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DependentUnknownMcpServerIdentifier
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DependentUnknownMcpServerIdentifier(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DependentUnknownMcpServerIdentifier) MarshalJSON() ([]byte, error) {
	type embed DependentUnknownMcpServerIdentifier
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
