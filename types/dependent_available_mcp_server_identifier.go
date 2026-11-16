package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DependentAvailableMcpServerIdentifier struct {
	Id                string                     `json:"id" xml:"id" required:"true"`
	Name              string                     `json:"name" xml:"name" required:"true"`
	CreatedAtUnixSecs int                        `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	AccessLevel       ProjectResponseAccessLevel `json:"access_level" xml:"access_level" required:"true"`
	extraProperties   map[string]interface{}
}

func (d *DependentAvailableMcpServerIdentifier) GetId() (value string) {
	if d == nil {
		return
	}
	return d.Id
}

func (d *DependentAvailableMcpServerIdentifier) GetName() (value string) {
	if d == nil {
		return
	}
	return d.Name
}

func (d *DependentAvailableMcpServerIdentifier) GetCreatedAtUnixSecs() (value int) {
	if d == nil {
		return
	}
	return d.CreatedAtUnixSecs
}

func (d *DependentAvailableMcpServerIdentifier) GetAccessLevel() (value ProjectResponseAccessLevel) {
	if d == nil {
		return
	}
	return d.AccessLevel
}

func (d *DependentAvailableMcpServerIdentifier) SetId(id string) {
	d.Id = id
}

func (d *DependentAvailableMcpServerIdentifier) SetName(name string) {
	d.Name = name
}

func (d *DependentAvailableMcpServerIdentifier) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	d.CreatedAtUnixSecs = createdAtUnixSecs
}

func (d *DependentAvailableMcpServerIdentifier) SetAccessLevel(accessLevel ProjectResponseAccessLevel) {
	d.AccessLevel = accessLevel
}

func (d *DependentAvailableMcpServerIdentifier) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DependentAvailableMcpServerIdentifier) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DependentAvailableMcpServerIdentifier to string"
	}
	return string(jsonData)
}

func (d *DependentAvailableMcpServerIdentifier) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DependentAvailableMcpServerIdentifier
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DependentAvailableMcpServerIdentifier(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DependentAvailableMcpServerIdentifier) MarshalJSON() ([]byte, error) {
	type embed DependentAvailableMcpServerIdentifier
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
