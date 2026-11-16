package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OpenAiFunctionTool struct {
	Name            string         `json:"name" xml:"name" required:"true"`
	Description     *string        `json:"description,omitempty" xml:"description,omitempty"`
	Parameters      map[string]any `json:"parameters,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OpenAiFunctionTool) GetName() (value string) {
	if o == nil {
		return
	}
	return o.Name
}

func (o *OpenAiFunctionTool) GetDescription() (value *string) {
	if o == nil {
		return
	}
	return o.Description
}

func (o *OpenAiFunctionTool) GetParameters() (value map[string]any) {
	if o == nil {
		return
	}
	return o.Parameters
}

func (o *OpenAiFunctionTool) SetName(name string) {
	o.Name = name
}

func (o *OpenAiFunctionTool) SetDescription(description *string) {
	o.Description = description
}

func (o *OpenAiFunctionTool) SetParameters(parameters map[string]any) {
	o.Parameters = parameters
}

func (o *OpenAiFunctionTool) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiFunctionTool) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiFunctionTool to string"
	}
	return string(jsonData)
}

func (o *OpenAiFunctionTool) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, o); err != nil {
		return err
	}
	type alias OpenAiFunctionTool
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*o = OpenAiFunctionTool(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *o, "type")
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenAiFunctionTool) MarshalJSON() ([]byte, error) {
	type embed OpenAiFunctionTool
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
