package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A model that represents an tool dependent on a knowledge base/tools
// to which the user has no direct access.
type DependentUnknownToolIdentifier struct {
	Id              string `json:"id" xml:"id" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DependentUnknownToolIdentifier) GetId() (value string) {
	if d == nil {
		return
	}
	return d.Id
}

func (d *DependentUnknownToolIdentifier) SetId(id string) {
	d.Id = id
}

func (d *DependentUnknownToolIdentifier) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DependentUnknownToolIdentifier) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DependentUnknownToolIdentifier to string"
	}
	return string(jsonData)
}

func (d *DependentUnknownToolIdentifier) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DependentUnknownToolIdentifier
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DependentUnknownToolIdentifier(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DependentUnknownToolIdentifier) MarshalJSON() ([]byte, error) {
	type embed DependentUnknownToolIdentifier
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
