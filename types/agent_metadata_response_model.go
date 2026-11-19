package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentMetadataResponseModel struct {
	// The creation time of the agent in unix seconds
	CreatedAtUnixSecs int `json:"created_at_unix_secs" xml:"created_at_unix_secs" required:"true"`
	// The last update time of the agent in unix seconds
	UpdatedAtUnixSecs int `json:"updated_at_unix_secs" xml:"updated_at_unix_secs" required:"true"`
	extraProperties   map[string]interface{}
}

func (a *AgentMetadataResponseModel) GetCreatedAtUnixSecs() (value int) {
	if a == nil {
		return
	}
	return a.CreatedAtUnixSecs
}

func (a *AgentMetadataResponseModel) GetUpdatedAtUnixSecs() (value int) {
	if a == nil {
		return
	}
	return a.UpdatedAtUnixSecs
}

func (a *AgentMetadataResponseModel) SetCreatedAtUnixSecs(createdAtUnixSecs int) {
	a.CreatedAtUnixSecs = createdAtUnixSecs
}

func (a *AgentMetadataResponseModel) SetUpdatedAtUnixSecs(updatedAtUnixSecs int) {
	a.UpdatedAtUnixSecs = updatedAtUnixSecs
}

func (a *AgentMetadataResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentMetadataResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentMetadataResponseModel to string"
	}
	return string(jsonData)
}

func (a *AgentMetadataResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentMetadataResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentMetadataResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentMetadataResponseModel) MarshalJSON() ([]byte, error) {
	type embed AgentMetadataResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
