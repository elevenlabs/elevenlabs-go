package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ToolMockConfig struct {
	DefaultReturnValue *string `json:"default_return_value,omitempty" xml:"default_return_value,omitempty"`
	DefaultIsError     *bool   `json:"default_is_error,omitempty" xml:"default_is_error,omitempty"`
	extraProperties    map[string]interface{}
}

func (t *ToolMockConfig) GetDefaultReturnValue() (value *string) {
	if t == nil {
		return
	}
	return t.DefaultReturnValue
}

func (t *ToolMockConfig) GetDefaultIsError() (value *bool) {
	if t == nil {
		return
	}
	return t.DefaultIsError
}

func (t *ToolMockConfig) SetDefaultReturnValue(defaultReturnValue *string) {
	t.DefaultReturnValue = defaultReturnValue
}

func (t *ToolMockConfig) SetDefaultIsError(defaultIsError *bool) {
	t.DefaultIsError = defaultIsError
}

func (t *ToolMockConfig) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t ToolMockConfig) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: ToolMockConfig to string"
	}
	return string(jsonData)
}

func (t *ToolMockConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler ToolMockConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = ToolMockConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *ToolMockConfig) MarshalJSON() ([]byte, error) {
	type embed ToolMockConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
