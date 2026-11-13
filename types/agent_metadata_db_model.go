package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentMetadataDbModel struct {
	CreatedAtUnixSecs int                    `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	UpdatedAtUnixSecs int                    `json:"updated_at_unix_secs" xml:"updated_at_unix_secs" required:"true"`
	CreatedFrom       *AgentDefinitionSource `json:"created_from,omitempty" xml:"created_from,omitempty"`
	LastUpdatedFrom   *AgentDefinitionSource `json:"last_updated_from,omitempty" xml:"last_updated_from,omitempty"`
	extraProperties   map[string]interface{}
}

func (a *AgentMetadataDbModel) GetCreatedAtUnixSecs() (value int) {
	if a == nil {
		return
	}
	return a.CreatedAtUnixSecs
}

func (a *AgentMetadataDbModel) GetUpdatedAtUnixSecs() (value int) {
	if a == nil {
		return
	}
	return a.UpdatedAtUnixSecs
}

func (a *AgentMetadataDbModel) GetCreatedFrom() (value *AgentDefinitionSource) {
	if a == nil {
		return
	}
	return a.CreatedFrom
}

func (a *AgentMetadataDbModel) GetLastUpdatedFrom() (value *AgentDefinitionSource) {
	if a == nil {
		return
	}
	return a.LastUpdatedFrom
}

func (a *AgentMetadataDbModel) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	a.CreatedAtUnixSecs = createdAtUnixSecs
}

func (a *AgentMetadataDbModel) SetUpdatedAtUnixSecs(updatedAtUnixSecs int) {
	a.UpdatedAtUnixSecs = updatedAtUnixSecs
}

func (a *AgentMetadataDbModel) SetCreatedFrom(createdFrom *AgentDefinitionSource) {
	a.CreatedFrom = createdFrom
}

func (a *AgentMetadataDbModel) SetLastUpdatedFrom(lastUpdatedFrom *AgentDefinitionSource) {
	a.LastUpdatedFrom = lastUpdatedFrom
}

func (a *AgentMetadataDbModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentMetadataDbModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentMetadataDbModel to string"
	}
	return string(jsonData)
}

func (a *AgentMetadataDbModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentMetadataDbModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentMetadataDbModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentMetadataDbModel) MarshalJSON() ([]byte, error) {
	type embed AgentMetadataDbModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
