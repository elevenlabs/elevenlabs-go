package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type MultimodalMessageText struct {
	Id              *string                            `json:"id,omitempty" xml:"id,omitempty"`
	Type            *string                            `json:"type,omitempty" xml:"type,omitempty"`
	Text            *string                            `json:"text,omitempty" xml:"text,omitempty"`
	UserIdentifier  *string                            `json:"user_identifier,omitempty" xml:"user_identifier,omitempty"`
	SourceMedium    *MultimodalMessageTextSourceMedium `json:"source_medium,omitempty" xml:"source_medium,omitempty"`
	extraProperties map[string]interface{}
}

func (m *MultimodalMessageText) GetId() (value *string) {
	if m == nil {
		return
	}
	return m.Id
}

func (m *MultimodalMessageText) GetType() (value *string) {
	if m == nil {
		return
	}
	return m.Type
}

func (m *MultimodalMessageText) GetText() (value *string) {
	if m == nil {
		return
	}
	return m.Text
}

func (m *MultimodalMessageText) GetUserIdentifier() (value *string) {
	if m == nil {
		return
	}
	return m.UserIdentifier
}

func (m *MultimodalMessageText) GetSourceMedium() (value *MultimodalMessageTextSourceMedium) {
	if m == nil {
		return
	}
	return m.SourceMedium
}

func (m *MultimodalMessageText) SetId(id *string) {
	m.Id = id
}

func (m *MultimodalMessageText) SetType(type_ *string) {
	m.Type = type_
}

func (m *MultimodalMessageText) SetText(text *string) {
	m.Text = text
}

func (m *MultimodalMessageText) SetUserIdentifier(userIdentifier *string) {
	m.UserIdentifier = userIdentifier
}

func (m *MultimodalMessageText) SetSourceMedium(sourceMedium *MultimodalMessageTextSourceMedium) {
	m.SourceMedium = sourceMedium
}

func (m *MultimodalMessageText) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MultimodalMessageText) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MultimodalMessageText to string"
	}
	return string(jsonData)
}

func (m *MultimodalMessageText) UnmarshalJSON(data []byte) error {
	type unmarshaler MultimodalMessageText
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = MultimodalMessageText(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MultimodalMessageText) MarshalJSON() ([]byte, error) {
	type embed MultimodalMessageText
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
