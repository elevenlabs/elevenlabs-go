package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LlmListResponseModel struct {
	// List of all available LLM models that can be used with agents.
	Llms                     []*LlmInfoModel            `json:"llms" xml:"llms" required:"true"`
	DefaultDeprecationConfig *LlmDeprecationConfigModel `json:"default_deprecation_config" xml:"default_deprecation_config" required:"true"`
	extraProperties          map[string]interface{}
}

func (l *LlmListResponseModel) GetLlms() (value []*LlmInfoModel) {
	if l == nil {
		return
	}
	return l.Llms
}

func (l *LlmListResponseModel) GetDefaultDeprecationConfig() (value *LlmDeprecationConfigModel) {
	if l == nil {
		return
	}
	return l.DefaultDeprecationConfig
}

func (l *LlmListResponseModel) SetLlms(llms []*LlmInfoModel) {
	l.Llms = llms
}

func (l *LlmListResponseModel) SetDefaultDeprecationConfig(defaultDeprecationConfig *LlmDeprecationConfigModel) {
	l.DefaultDeprecationConfig = defaultDeprecationConfig
}

func (l *LlmListResponseModel) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmListResponseModel) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmListResponseModel to string"
	}
	return string(jsonData)
}

func (l *LlmListResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias LlmListResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = LlmListResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmListResponseModel) MarshalJSON() ([]byte, error) {
	type embed LlmListResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
