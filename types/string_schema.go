package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type StringSchema struct {
	Title           *string  `json:"title,omitempty" xml:"title,omitempty"`
	Description     *string  `json:"description,omitempty" xml:"description,omitempty" maxLength:"500"`
	Type            *string  `json:"type,omitempty" xml:"type,omitempty"`
	Enum            []string `json:"enum,omitempty" xml:"enum,omitempty"`
	extraProperties map[string]interface{}
}

func (s *StringSchema) GetTitle() (value *string) {
	if s == nil {
		return
	}
	return s.Title
}

func (s *StringSchema) GetDescription() (value *string) {
	if s == nil {
		return
	}
	return s.Description
}

func (s *StringSchema) GetType() (value *string) {
	if s == nil {
		return
	}
	return s.Type
}

func (s *StringSchema) GetEnum() (value []string) {
	if s == nil {
		return
	}
	return s.Enum
}

func (s *StringSchema) SetTitle(title *string) {
	s.Title = title
}

func (s *StringSchema) SetDescription(description *string) {
	s.Description = description
}

func (s *StringSchema) SetType(type_ *string) {
	s.Type = type_
}

func (s *StringSchema) SetEnum(enum []string) {
	s.Enum = enum
}

func (s *StringSchema) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StringSchema) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StringSchema to string"
	}
	return string(jsonData)
}

func (s *StringSchema) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *StringSchema) MarshalJSON() ([]byte, error) {
	type embed StringSchema
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
