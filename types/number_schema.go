package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type NumberSchema struct {
	Title           *string   `json:"title,omitempty" xml:"title,omitempty"`
	Description     *string   `json:"description,omitempty" xml:"description,omitempty" maxLength:"500"`
	Type            *string   `json:"type,omitempty" xml:"type,omitempty"`
	Enum            []float64 `json:"enum,omitempty" xml:"enum,omitempty"`
	extraProperties map[string]interface{}
}

func (n *NumberSchema) GetTitle() (value *string) {
	if n == nil {
		return
	}
	return n.Title
}

func (n *NumberSchema) GetDescription() (value *string) {
	if n == nil {
		return
	}
	return n.Description
}

func (n *NumberSchema) GetType() (value *string) {
	if n == nil {
		return
	}
	return n.Type
}

func (n *NumberSchema) GetEnum() (value []float64) {
	if n == nil {
		return
	}
	return n.Enum
}

func (n *NumberSchema) SetTitle(title *string) {
	n.Title = title
}

func (n *NumberSchema) SetDescription(description *string) {
	n.Description = description
}

func (n *NumberSchema) SetType(type_ *string) {
	n.Type = type_
}

func (n *NumberSchema) SetEnum(enum []float64) {
	n.Enum = enum
}

func (n *NumberSchema) GetExtraProperties() map[string]interface{} {
	if n == nil {
		return nil
	}
	return n.extraProperties
}

func (n NumberSchema) String() string {
	jsonData, err := json.MarshalIndent(n, "", "  ")
	if err != nil {
		return "error converting struct: NumberSchema to string"
	}
	return string(jsonData)
}

func (n *NumberSchema) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, n); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *n)
	if err != nil {
		return err
	}
	n.extraProperties = extraProperties
	return nil
}

func (n *NumberSchema) MarshalJSON() ([]byte, error) {
	type embed NumberSchema
	var marshaler = struct {
		embed
	}{
		embed: embed(*n),
	}
	return json.Marshal(marshaler)
}
