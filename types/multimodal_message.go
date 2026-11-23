package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type MultimodalMessage struct {
	Type *string                `json:"type,omitempty" xml:"type,omitempty"`
	Text *MultimodalMessageText `json:"text,omitempty" xml:"text,omitempty"`
	// Deprecated: single attachment, kept for clients that predate `files`. Ignored when `files` is set.
	File *MultimodalMessageFile `json:"file,omitempty" xml:"file,omitempty"`
	// Files attached to this message, at most 5. Each file must already have been uploaded via the conversation file upload endpoint.
	Files           []*MultimodalMessageFilesItem `json:"files,omitempty" xml:"files,omitempty"`
	extraProperties map[string]interface{}
}

func (m *MultimodalMessage) GetType() (value *string) {
	if m == nil {
		return
	}
	return m.Type
}

func (m *MultimodalMessage) GetText() (value *MultimodalMessageText) {
	if m == nil {
		return
	}
	return m.Text
}

func (m *MultimodalMessage) GetFile() (value *MultimodalMessageFile) {
	if m == nil {
		return
	}
	return m.File
}

func (m *MultimodalMessage) GetFiles() (value []*MultimodalMessageFilesItem) {
	if m == nil {
		return
	}
	return m.Files
}

func (m *MultimodalMessage) SetType(type_ *string) {
	m.Type = type_
}

func (m *MultimodalMessage) SetText(text *MultimodalMessageText) {
	m.Text = text
}

func (m *MultimodalMessage) SetFile(file *MultimodalMessageFile) {
	m.File = file
}

func (m *MultimodalMessage) SetFiles(files []*MultimodalMessageFilesItem) {
	m.Files = files
}

func (m *MultimodalMessage) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MultimodalMessage) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MultimodalMessage to string"
	}
	return string(jsonData)
}

func (m *MultimodalMessage) UnmarshalJSON(data []byte) error {
	type unmarshaler MultimodalMessage
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MultimodalMessage(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MultimodalMessage) MarshalJSON() ([]byte, error) {
	type embed MultimodalMessage
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
