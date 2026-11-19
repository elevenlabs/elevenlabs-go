package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceAccentResponseModel struct {
	// The accent value used for filtering shared voices via the `accent` query parameter on `GET /v1/shared-voices`.
	Accent string `json:"accent" xml:"accent" required:"true"`
	// The language code this accent belongs to, e.g. `en`.
	Language string `json:"language" xml:"language" required:"true"`
	// The full accent code, e.g. `en-american`.
	Code string `json:"code" xml:"code" required:"true"`
	// The human-readable accent name, e.g. `American`.
	Name            string `json:"name" xml:"name" required:"true"`
	extraProperties map[string]interface{}
}

func (v *VoiceAccentResponseModel) GetAccent() (value string) {
	if v == nil {
		return
	}
	return v.Accent
}

func (v *VoiceAccentResponseModel) GetLanguage() (value string) {
	if v == nil {
		return
	}
	return v.Language
}

func (v *VoiceAccentResponseModel) GetCode() (value string) {
	if v == nil {
		return
	}
	return v.Code
}

func (v *VoiceAccentResponseModel) GetName() (value string) {
	if v == nil {
		return
	}
	return v.Name
}

func (v *VoiceAccentResponseModel) SetAccent(accent string) {
	v.Accent = accent
}

func (v *VoiceAccentResponseModel) SetLanguage(language string) {
	v.Language = language
}

func (v *VoiceAccentResponseModel) SetCode(code string) {
	v.Code = code
}

func (v *VoiceAccentResponseModel) SetName(name string) {
	v.Name = name
}

func (v *VoiceAccentResponseModel) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoiceAccentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceAccentResponseModel to string"
	}
	return string(jsonData)
}

func (v *VoiceAccentResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, v); err != nil {
		return err
	}
	type alias VoiceAccentResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*v = VoiceAccentResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VoiceAccentResponseModel) MarshalJSON() ([]byte, error) {
	type embed VoiceAccentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
