package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UserDataCollectionAnalysisItemInput struct {
	Name   string  `json:"name" xml:"name" required:"true"`
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// LLM to use for this item. If unset, the attaching agent's analysis_llm default is used.
	Llm  *Llm    `json:"llm,omitempty" xml:"llm,omitempty"`
	Kind *string `json:"kind,omitempty" xml:"kind,omitempty"`
	// Schema property for literal JSON types. IMPORTANT: Only ONE of the following fields can be set: description (LLM provides value), dynamic_variable (value from variable), is_system_provided (system provides value), constant_value (fixed value), or is_omitted (parameter is omitted). These are mutually exclusive.
	ValueSchema     *DataCollectionValueSchema `json:"value_schema" xml:"value_schema" required:"true"`
	extraProperties map[string]interface{}
}

func (u *UserDataCollectionAnalysisItemInput) GetName() (value string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UserDataCollectionAnalysisItemInput) GetSource() (value *string) {
	if u == nil {
		return
	}
	return u.Source
}

func (u *UserDataCollectionAnalysisItemInput) GetLlm() (value *Llm) {
	if u == nil {
		return
	}
	return u.Llm
}

func (u *UserDataCollectionAnalysisItemInput) GetKind() (value *string) {
	if u == nil {
		return
	}
	return u.Kind
}

func (u *UserDataCollectionAnalysisItemInput) GetValueSchema() (value *DataCollectionValueSchema) {
	if u == nil {
		return
	}
	return u.ValueSchema
}

func (u *UserDataCollectionAnalysisItemInput) SetName(name string) {
	u.Name = name
}

func (u *UserDataCollectionAnalysisItemInput) SetSource(source *string) {
	u.Source = source
}

func (u *UserDataCollectionAnalysisItemInput) SetLlm(llm *Llm) {
	u.Llm = llm
}

func (u *UserDataCollectionAnalysisItemInput) SetKind(kind *string) {
	u.Kind = kind
}

func (u *UserDataCollectionAnalysisItemInput) SetValueSchema(valueSchema *DataCollectionValueSchema) {
	u.ValueSchema = valueSchema
}

func (u *UserDataCollectionAnalysisItemInput) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UserDataCollectionAnalysisItemInput) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UserDataCollectionAnalysisItemInput to string"
	}
	return string(jsonData)
}

func (u *UserDataCollectionAnalysisItemInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, u); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UserDataCollectionAnalysisItemInput) MarshalJSON() ([]byte, error) {
	type embed UserDataCollectionAnalysisItemInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
