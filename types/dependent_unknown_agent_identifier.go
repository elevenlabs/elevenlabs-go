package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A model that represents an agent dependent on a knowledge base/tools
// to which the user has no direct access.
type DependentUnknownAgentIdentifier struct {
	// If the agent is a transitive dependent, contains IDs of the resources that the agent depends on directly.
	ReferencedResourceIds []string `json:"referenced_resource_ids,omitempty" xml:"referenced_resource_ids,omitempty"`
	Id                    string   `json:"id" xml:"id" required:"true"`
	extraProperties       map[string]interface{}
}

func (d *DependentUnknownAgentIdentifier) GetReferencedResourceIds() (value []string) {
	if d == nil {
		return
	}
	return d.ReferencedResourceIds
}

func (d *DependentUnknownAgentIdentifier) GetId() (value string) {
	if d == nil {
		return
	}
	return d.Id
}

func (d *DependentUnknownAgentIdentifier) SetReferencedResourceIds(referencedResourceIds []string) {
	d.ReferencedResourceIds = referencedResourceIds
}

func (d *DependentUnknownAgentIdentifier) SetId(id string) {
	d.Id = id
}

func (d *DependentUnknownAgentIdentifier) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DependentUnknownAgentIdentifier) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DependentUnknownAgentIdentifier to string"
	}
	return string(jsonData)
}

func (d *DependentUnknownAgentIdentifier) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DependentUnknownAgentIdentifier
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DependentUnknownAgentIdentifier(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DependentUnknownAgentIdentifier) MarshalJSON() ([]byte, error) {
	type embed DependentUnknownAgentIdentifier
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
