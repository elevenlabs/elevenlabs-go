package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DependentAvailableAgentIdentifier struct {
	// If the agent is a transitive dependent, contains IDs of the resources that the agent depends on directly.
	ReferencedResourceIds []string                   `json:"referenced_resource_ids,omitempty" xml:"referenced_resource_ids,omitempty"`
	Id                    string                     `json:"id" xml:"id" required:"true"`
	Name                  string                     `json:"name" xml:"name" required:"true"`
	CreatedAtUnixSecs     int                        `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	AccessLevel           ProjectResponseAccessLevel `json:"access_level" xml:"access_level" required:"true"`
	extraProperties       map[string]interface{}
}

func (d *DependentAvailableAgentIdentifier) GetReferencedResourceIds() (value []string) {
	if d == nil {
		return
	}
	return d.ReferencedResourceIds
}

func (d *DependentAvailableAgentIdentifier) GetId() (value string) {
	if d == nil {
		return
	}
	return d.Id
}

func (d *DependentAvailableAgentIdentifier) GetName() (value string) {
	if d == nil {
		return
	}
	return d.Name
}

func (d *DependentAvailableAgentIdentifier) GetCreatedAtUnixSecs() (value int) {
	if d == nil {
		return
	}
	return d.CreatedAtUnixSecs
}

func (d *DependentAvailableAgentIdentifier) GetAccessLevel() (value ProjectResponseAccessLevel) {
	if d == nil {
		return
	}
	return d.AccessLevel
}

func (d *DependentAvailableAgentIdentifier) SetReferencedResourceIds(referencedResourceIds []string) {
	d.ReferencedResourceIds = referencedResourceIds
}

func (d *DependentAvailableAgentIdentifier) SetId(id string) {
	d.Id = id
}

func (d *DependentAvailableAgentIdentifier) SetName(name string) {
	d.Name = name
}

func (d *DependentAvailableAgentIdentifier) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	d.CreatedAtUnixSecs = createdAtUnixSecs
}

func (d *DependentAvailableAgentIdentifier) SetAccessLevel(accessLevel ProjectResponseAccessLevel) {
	d.AccessLevel = accessLevel
}

func (d *DependentAvailableAgentIdentifier) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DependentAvailableAgentIdentifier) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DependentAvailableAgentIdentifier to string"
	}
	return string(jsonData)
}

func (d *DependentAvailableAgentIdentifier) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DependentAvailableAgentIdentifier
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DependentAvailableAgentIdentifier(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DependentAvailableAgentIdentifier) MarshalJSON() ([]byte, error) {
	type embed DependentAvailableAgentIdentifier
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
