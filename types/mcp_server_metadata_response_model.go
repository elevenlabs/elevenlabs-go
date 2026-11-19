package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type McpServerMetadataResponseModel struct {
	CreatedAt       int     `json:"created_at" xml:"created_at" required:"true"`
	OwnerUserId     *string `json:"owner_user_id,omitempty" xml:"owner_user_id,omitempty"`
	extraProperties map[string]interface{}
}

func (m *McpServerMetadataResponseModel) GetCreatedAt() (value int) {
	if m == nil {
		return
	}
	return m.CreatedAt
}

func (m *McpServerMetadataResponseModel) GetOwnerUserId() (value *string) {
	if m == nil {
		return
	}
	return m.OwnerUserId
}

func (m *McpServerMetadataResponseModel) SetCreatedAt(createdAt int) {
	m.CreatedAt = createdAt
}

func (m *McpServerMetadataResponseModel) SetOwnerUserId(ownerUserId *string) {
	m.OwnerUserId = ownerUserId
}

func (m *McpServerMetadataResponseModel) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpServerMetadataResponseModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpServerMetadataResponseModel to string"
	}
	return string(jsonData)
}

func (m *McpServerMetadataResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpServerMetadataResponseModel) MarshalJSON() ([]byte, error) {
	type embed McpServerMetadataResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
