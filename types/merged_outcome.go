package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type MergedOutcome struct {
	VersionId       string  `json:"version_id" xml:"version_id" required:"true"`
	At              int     `json:"at" xml:"at" required:"true"`
	ByUserId        *string `json:"by_user_id" xml:"by_user_id" required:"true"`
	extraProperties map[string]interface{}
}

func (m *MergedOutcome) GetVersionId() (value string) {
	if m == nil {
		return
	}
	return m.VersionId
}

func (m *MergedOutcome) GetAt() (value int) {
	if m == nil {
		return
	}
	return m.At
}

func (m *MergedOutcome) GetByUserId() (value *string) {
	if m == nil {
		return
	}
	return m.ByUserId
}

func (m *MergedOutcome) SetVersionId(versionId string) {
	m.VersionId = versionId
}

func (m *MergedOutcome) SetAt(at int) {
	m.At = at
}

func (m *MergedOutcome) SetByUserId(byUserId *string) {
	m.ByUserId = byUserId
}

func (m *MergedOutcome) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MergedOutcome) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MergedOutcome to string"
	}
	return string(jsonData)
}

func (m *MergedOutcome) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m, "status")
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MergedOutcome) MarshalJSON() ([]byte, error) {
	type embed MergedOutcome
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
