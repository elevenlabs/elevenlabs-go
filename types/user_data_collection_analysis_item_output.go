package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UserDataCollectionAnalysisItemOutput struct {
	Name   string  `json:"name" xml:"name" required:"true"`
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// LLM to use for this item. If unset, the attaching agent's analysis_llm default is used.
	Llm *Llm `json:"llm,omitempty" xml:"llm,omitempty"`
	// Schema property for literal JSON types. IMPORTANT: Only ONE of the following fields can be set: description (LLM provides value), dynamic_variable (value from variable), is_system_provided (system provides value), constant_value (fixed value), or is_omitted (parameter is omitted). These are mutually exclusive.
	ValueSchema     *DataCollectionValueSchema `json:"value_schema" xml:"value_schema" required:"true"`
	extraProperties map[string]interface{}
}

func (u *UserDataCollectionAnalysisItemOutput) GetName() (value string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UserDataCollectionAnalysisItemOutput) GetSource() (value *string) {
	if u == nil {
		return
	}
	return u.Source
}

func (u *UserDataCollectionAnalysisItemOutput) GetLlm() (value *Llm) {
	if u == nil {
		return
	}
	return u.Llm
}

func (u *UserDataCollectionAnalysisItemOutput) GetValueSchema() (value *DataCollectionValueSchema) {
	if u == nil {
		return
	}
	return u.ValueSchema
}

func (u *UserDataCollectionAnalysisItemOutput) SetName(name string) {
	u.Name = name
}

func (u *UserDataCollectionAnalysisItemOutput) SetSource(source *string) {
	u.Source = source
}

func (u *UserDataCollectionAnalysisItemOutput) SetLlm(llm *Llm) {
	u.Llm = llm
}

func (u *UserDataCollectionAnalysisItemOutput) SetValueSchema(valueSchema *DataCollectionValueSchema) {
	u.ValueSchema = valueSchema
}

func (u *UserDataCollectionAnalysisItemOutput) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UserDataCollectionAnalysisItemOutput) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UserDataCollectionAnalysisItemOutput to string"
	}
	return string(jsonData)
}

func (u *UserDataCollectionAnalysisItemOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, u); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "kind")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UserDataCollectionAnalysisItemOutput) MarshalJSON() ([]byte, error) {
	type embed UserDataCollectionAnalysisItemOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
