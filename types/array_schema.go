package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ArraySchema struct {
	Title           *string        `json:"title,omitempty" xml:"title,omitempty"`
	Description     *string        `json:"description,omitempty" xml:"description,omitempty" maxLength:"500"`
	Type            *string        `json:"type,omitempty" xml:"type,omitempty"`
	Items           *ContentSchema `json:"items" xml:"items" required:"true"`
	extraProperties map[string]interface{}
}

func (a *ArraySchema) GetTitle() (value *string) {
	if a == nil {
		return
	}
	return a.Title
}

func (a *ArraySchema) GetDescription() (value *string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *ArraySchema) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *ArraySchema) GetItems() (value *ContentSchema) {
	if a == nil {
		return
	}
	return a.Items
}

func (a *ArraySchema) SetTitle(title *string) {
	a.Title = title
}

func (a *ArraySchema) SetDescription(description *string) {
	a.Description = description
}

func (a *ArraySchema) SetType(type_ *string) {
	a.Type = type_
}

func (a *ArraySchema) SetItems(items *ContentSchema) {
	a.Items = items
}

func (a *ArraySchema) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a ArraySchema) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: ArraySchema to string"
	}
	return string(jsonData)
}

func (a *ArraySchema) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *ArraySchema) MarshalJSON() ([]byte, error) {
	type embed ArraySchema
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
