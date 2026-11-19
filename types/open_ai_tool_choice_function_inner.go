package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OpenAiToolChoiceFunctionInner struct {
	Name            string `json:"name" xml:"name" required:"true"`
	extraProperties map[string]interface{}
}

func (o *OpenAiToolChoiceFunctionInner) GetName() (value string) {
	if o == nil {
		return
	}
	return o.Name
}

func (o *OpenAiToolChoiceFunctionInner) SetName(name string) {
	o.Name = name
}

func (o *OpenAiToolChoiceFunctionInner) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiToolChoiceFunctionInner) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiToolChoiceFunctionInner to string"
	}
	return string(jsonData)
}

func (o *OpenAiToolChoiceFunctionInner) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, o); err != nil {
		return err
	}
	type alias OpenAiToolChoiceFunctionInner
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*o = OpenAiToolChoiceFunctionInner(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenAiToolChoiceFunctionInner) MarshalJSON() ([]byte, error) {
	type embed OpenAiToolChoiceFunctionInner
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
