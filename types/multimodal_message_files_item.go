package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type MultimodalMessageFilesItem struct {
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	FileId          string  `json:"file_id" xml:"file_id" required:"true"`
	extraProperties map[string]interface{}
}

func (m *MultimodalMessageFilesItem) GetType() (value *string) {
	if m == nil {
		return
	}
	return m.Type
}

func (m *MultimodalMessageFilesItem) GetFileId() (value string) {
	if m == nil {
		return
	}
	return m.FileId
}

func (m *MultimodalMessageFilesItem) SetType(type_ *string) {
	m.Type = type_
}

func (m *MultimodalMessageFilesItem) SetFileId(fileId string) {
	m.FileId = fileId
}

func (m *MultimodalMessageFilesItem) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MultimodalMessageFilesItem) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MultimodalMessageFilesItem to string"
	}
	return string(jsonData)
}

func (m *MultimodalMessageFilesItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias MultimodalMessageFilesItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = MultimodalMessageFilesItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MultimodalMessageFilesItem) MarshalJSON() ([]byte, error) {
	type embed MultimodalMessageFilesItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
