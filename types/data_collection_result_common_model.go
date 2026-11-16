package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DataCollectionResultCommonModel struct {
	DataCollectionId string                     `json:"data_collection_id" xml:"data_collection_id" required:"true"`
	Name             *string                    `json:"name,omitempty" xml:"name,omitempty"`
	Value            any                        `json:"value,omitempty" xml:"value,omitempty"`
	JsonSchema       *LiteralJsonSchemaProperty `json:"json_schema,omitempty" xml:"json_schema,omitempty"`
	Rationale        string                     `json:"rationale" xml:"rationale" required:"true"`
	extraProperties  map[string]interface{}
}

func (d *DataCollectionResultCommonModel) GetDataCollectionId() (value string) {
	if d == nil {
		return
	}
	return d.DataCollectionId
}

func (d *DataCollectionResultCommonModel) GetName() (value *string) {
	if d == nil {
		return
	}
	return d.Name
}

func (d *DataCollectionResultCommonModel) GetValue() (value any) {
	if d == nil {
		return
	}
	return d.Value
}

func (d *DataCollectionResultCommonModel) GetJsonSchema() (value *LiteralJsonSchemaProperty) {
	if d == nil {
		return
	}
	return d.JsonSchema
}

func (d *DataCollectionResultCommonModel) GetRationale() (value string) {
	if d == nil {
		return
	}
	return d.Rationale
}

func (d *DataCollectionResultCommonModel) SetDataCollectionId(dataCollectionId string) {
	d.DataCollectionId = dataCollectionId
}

func (d *DataCollectionResultCommonModel) SetName(name *string) {
	d.Name = name
}

func (d *DataCollectionResultCommonModel) SetValue(value any) {
	d.Value = value
}

func (d *DataCollectionResultCommonModel) SetJsonSchema(jsonSchema *LiteralJsonSchemaProperty) {
	d.JsonSchema = jsonSchema
}

func (d *DataCollectionResultCommonModel) SetRationale(rationale string) {
	d.Rationale = rationale
}

func (d *DataCollectionResultCommonModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DataCollectionResultCommonModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DataCollectionResultCommonModel to string"
	}
	return string(jsonData)
}

func (d *DataCollectionResultCommonModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DataCollectionResultCommonModel) MarshalJSON() ([]byte, error) {
	type embed DataCollectionResultCommonModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
