package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type QueryParamsJsonSchemaOutput struct {
	Properties      map[string]*QueryParamsJsonSchemaOutputPropertiesValue `json:"properties" required:"true"`
	Required        []string                                               `json:"required,omitempty" xml:"required,omitempty"`
	extraProperties map[string]interface{}
}

func (q *QueryParamsJsonSchemaOutput) GetProperties() (value map[string]*QueryParamsJsonSchemaOutputPropertiesValue) {
	if q == nil {
		return
	}
	return q.Properties
}

func (q *QueryParamsJsonSchemaOutput) GetRequired() (value []string) {
	if q == nil {
		return
	}
	return q.Required
}

func (q *QueryParamsJsonSchemaOutput) SetProperties(properties map[string]*QueryParamsJsonSchemaOutputPropertiesValue) {
	q.Properties = properties
}

func (q *QueryParamsJsonSchemaOutput) SetRequired(required []string) {
	q.Required = required
}

func (q *QueryParamsJsonSchemaOutput) GetExtraProperties() map[string]interface{} {
	if q == nil {
		return nil
	}
	return q.extraProperties
}

func (q QueryParamsJsonSchemaOutput) String() string {
	jsonData, err := json.MarshalIndent(q, "", "  ")
	if err != nil {
		return "error converting struct: QueryParamsJsonSchemaOutput to string"
	}
	return string(jsonData)
}

func (q *QueryParamsJsonSchemaOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler QueryParamsJsonSchemaOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*q = QueryParamsJsonSchemaOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *q)
	if err != nil {
		return err
	}
	q.extraProperties = extraProperties
	return nil
}

func (q *QueryParamsJsonSchemaOutput) MarshalJSON() ([]byte, error) {
	type embed QueryParamsJsonSchemaOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*q),
	}
	return json.Marshal(marshaler)
}
