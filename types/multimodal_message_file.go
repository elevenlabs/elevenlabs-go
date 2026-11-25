package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Deprecated: single attachment, kept for clients that predate `files`. Ignored when `files` is set.
type MultimodalMessageFile struct {
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	FileId          string  `json:"file_id" xml:"file_id" required:"true"`
	extraProperties map[string]interface{}
}

func (m *MultimodalMessageFile) GetType() (value *string) {
	if m == nil {
		return
	}
	return m.Type
}

func (m *MultimodalMessageFile) GetFileId() (value string) {
	if m == nil {
		return
	}
	return m.FileId
}

func (m *MultimodalMessageFile) SetType(type_ *string) {
	m.Type = type_
}

func (m *MultimodalMessageFile) SetFileId(fileId string) {
	m.FileId = fileId
}

func (m *MultimodalMessageFile) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MultimodalMessageFile) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MultimodalMessageFile to string"
	}
	return string(jsonData)
}

func (m *MultimodalMessageFile) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias MultimodalMessageFile
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = MultimodalMessageFile(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MultimodalMessageFile) MarshalJSON() ([]byte, error) {
	type embed MultimodalMessageFile
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
