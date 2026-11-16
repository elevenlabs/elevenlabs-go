package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OpenAiToolChoiceFunction struct {
	Type            *string                        `json:"type,omitempty" xml:"type,omitempty"`
	Function        *OpenAiToolChoiceFunctionInner `json:"function" xml:"function" required:"true"`
	extraProperties map[string]interface{}
}

func (o *OpenAiToolChoiceFunction) GetType() (value *string) {
	if o == nil {
		return
	}
	return o.Type
}

func (o *OpenAiToolChoiceFunction) GetFunction() (value *OpenAiToolChoiceFunctionInner) {
	if o == nil {
		return
	}
	return o.Function
}

func (o *OpenAiToolChoiceFunction) SetType(type_ *string) {
	o.Type = type_
}

func (o *OpenAiToolChoiceFunction) SetFunction(function *OpenAiToolChoiceFunctionInner) {
	o.Function = function
}

func (o *OpenAiToolChoiceFunction) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiToolChoiceFunction) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiToolChoiceFunction to string"
	}
	return string(jsonData)
}

func (o *OpenAiToolChoiceFunction) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, o); err != nil {
		return err
	}
	type alias OpenAiToolChoiceFunction
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*o = OpenAiToolChoiceFunction(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenAiToolChoiceFunction) MarshalJSON() ([]byte, error) {
	type embed OpenAiToolChoiceFunction
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
