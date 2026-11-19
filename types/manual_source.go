package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ManualSource struct {
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	CreatedByUserId string  `json:"created_by_user_id" xml:"created_by_user_id" required:"true"`
	Notes           *string `json:"notes,omitempty" xml:"notes,omitempty"`
	extraProperties map[string]interface{}
}

func (m *ManualSource) GetType() (value *string) {
	if m == nil {
		return
	}
	return m.Type
}

func (m *ManualSource) GetCreatedByUserId() (value string) {
	if m == nil {
		return
	}
	return m.CreatedByUserId
}

func (m *ManualSource) GetNotes() (value *string) {
	if m == nil {
		return
	}
	return m.Notes
}

func (m *ManualSource) SetType(type_ *string) {
	m.Type = type_
}

func (m *ManualSource) SetCreatedByUserId(createdByUserId string) {
	m.CreatedByUserId = createdByUserId
}

func (m *ManualSource) SetNotes(notes *string) {
	m.Notes = notes
}

func (m *ManualSource) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m ManualSource) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: ManualSource to string"
	}
	return string(jsonData)
}

func (m *ManualSource) UnmarshalJSON(data []byte) error {
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

func (m *ManualSource) MarshalJSON() ([]byte, error) {
	type embed ManualSource
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
