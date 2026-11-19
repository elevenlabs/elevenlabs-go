package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type QueryParamsJsonSchemaInput struct {
	Properties      map[string]*QueryParamsJsonSchemaInputPropertiesValue `json:"properties" required:"true"`
	Required        []string                                              `json:"required,omitempty" xml:"required,omitempty"`
	extraProperties map[string]interface{}
}

func (q *QueryParamsJsonSchemaInput) GetProperties() (value map[string]*QueryParamsJsonSchemaInputPropertiesValue) {
	if q == nil {
		return
	}
	return q.Properties
}

func (q *QueryParamsJsonSchemaInput) GetRequired() (value []string) {
	if q == nil {
		return
	}
	return q.Required
}

func (q *QueryParamsJsonSchemaInput) SetProperties(properties map[string]*QueryParamsJsonSchemaInputPropertiesValue) {
	q.Properties = properties
}

func (q *QueryParamsJsonSchemaInput) SetRequired(required []string) {
	q.Required = required
}

func (q *QueryParamsJsonSchemaInput) GetExtraProperties() map[string]interface{} {
	if q == nil {
		return nil
	}
	return q.extraProperties
}

func (q QueryParamsJsonSchemaInput) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueryParamsJsonSchemaInput to string"
	}
	return string(jsonData)
}

func (q *QueryParamsJsonSchemaInput) UnmarshalJSON(data []byte) error {
	type unmarshaler QueryParamsJsonSchemaInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*q = QueryParamsJsonSchemaInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *q)
	if err != nil {
		return err
	}
	q.extraProperties = extraProperties
	return nil
}

func (q *QueryParamsJsonSchemaInput) MarshalJSON() ([]byte, error) {
	type embed QueryParamsJsonSchemaInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*q),
	}
	return json.Marshal(marshaler)
}
