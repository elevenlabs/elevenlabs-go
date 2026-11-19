package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DependentAvailableToolIdentifier struct {
	Id                string                     `json:"id" xml:"id" required:"true"`
	Name              string                     `json:"name" xml:"name" required:"true"`
	CreatedAtUnixSecs int                        `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	AccessLevel       ProjectResponseAccessLevel `json:"access_level" xml:"access_level" required:"true"`
	extraProperties   map[string]interface{}
}

func (d *DependentAvailableToolIdentifier) GetId() (value string) {
	if d == nil {
		return
	}
	return d.Id
}

func (d *DependentAvailableToolIdentifier) GetName() (value string) {
	if d == nil {
		return
	}
	return d.Name
}

func (d *DependentAvailableToolIdentifier) GetCreatedAtUnixSecs() (value int) {
	if d == nil {
		return
	}
	return d.CreatedAtUnixSecs
}

func (d *DependentAvailableToolIdentifier) GetAccessLevel() (value ProjectResponseAccessLevel) {
	if d == nil {
		return
	}
	return d.AccessLevel
}

func (d *DependentAvailableToolIdentifier) SetId(id string) {
	d.Id = id
}

func (d *DependentAvailableToolIdentifier) SetName(name string) {
	d.Name = name
}

func (d *DependentAvailableToolIdentifier) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	d.CreatedAtUnixSecs = createdAtUnixSecs
}

func (d *DependentAvailableToolIdentifier) SetAccessLevel(accessLevel ProjectResponseAccessLevel) {
	d.AccessLevel = accessLevel
}

func (d *DependentAvailableToolIdentifier) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DependentAvailableToolIdentifier) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DependentAvailableToolIdentifier to string"
	}
	return string(jsonData)
}

func (d *DependentAvailableToolIdentifier) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DependentAvailableToolIdentifier
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DependentAvailableToolIdentifier(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DependentAvailableToolIdentifier) MarshalJSON() ([]byte, error) {
	type embed DependentAvailableToolIdentifier
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
