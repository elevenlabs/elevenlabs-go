package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ObjectSchema struct {
	Title           *string                   `json:"title,omitempty" xml:"title,omitempty"`
	Description     *string                   `json:"description,omitempty" xml:"description,omitempty" maxLength:"500"`
	Type            *string                   `json:"type,omitempty" xml:"type,omitempty"`
	Properties      map[string]*ContentSchema `json:"properties,omitempty"`
	Required        []string                  `json:"required,omitempty" xml:"required,omitempty"`
	extraProperties map[string]interface{}
}

func (o *ObjectSchema) GetTitle() (value *string) {
	if o == nil {
		return
	}
	return o.Title
}

func (o *ObjectSchema) GetDescription() (value *string) {
	if o == nil {
		return
	}
	return o.Description
}

func (o *ObjectSchema) GetType() (value *string) {
	if o == nil {
		return
	}
	return o.Type
}

func (o *ObjectSchema) GetProperties() (value map[string]*ContentSchema) {
	if o == nil {
		return
	}
	return o.Properties
}

func (o *ObjectSchema) GetRequired() (value []string) {
	if o == nil {
		return
	}
	return o.Required
}

func (o *ObjectSchema) SetTitle(title *string) {
	o.Title = title
}

func (o *ObjectSchema) SetDescription(description *string) {
	o.Description = description
}

func (o *ObjectSchema) SetType(type_ *string) {
	o.Type = type_
}

func (o *ObjectSchema) SetProperties(properties map[string]*ContentSchema) {
	o.Properties = properties
}

func (o *ObjectSchema) SetRequired(required []string) {
	o.Required = required
}

func (o *ObjectSchema) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o ObjectSchema) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: ObjectSchema to string"
	}
	return string(jsonData)
}

func (o *ObjectSchema) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, o); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *ObjectSchema) MarshalJSON() ([]byte, error) {
	type embed ObjectSchema
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
