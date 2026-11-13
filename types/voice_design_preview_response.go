package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceDesignPreviewResponse struct {
	// The previews of the generated voices.
	Previews []*VoicePreviewResponseModel `json:"previews" xml:"previews" required:"true"`
	// The text used to preview the voices.
	Text            string `json:"text" xml:"text" required:"true"`
	extraProperties map[string]interface{}
}

func (v *VoiceDesignPreviewResponse) GetPreviews() (value []*VoicePreviewResponseModel) {
	if v == nil {
		return
	}
	return v.Previews
}

func (v *VoiceDesignPreviewResponse) GetText() (value string) {
	if v == nil {
		return
	}
	return v.Text
}

func (v *VoiceDesignPreviewResponse) SetPreviews(previews []*VoicePreviewResponseModel) {
	v.Previews = previews
}

func (v *VoiceDesignPreviewResponse) SetText(text string) {
	v.Text = text
}

func (v *VoiceDesignPreviewResponse) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoiceDesignPreviewResponse) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceDesignPreviewResponse to string"
	}
	return string(jsonData)
}

func (v *VoiceDesignPreviewResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, v); err != nil {
		return err
	}
	type alias VoiceDesignPreviewResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*v = VoiceDesignPreviewResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VoiceDesignPreviewResponse) MarshalJSON() ([]byte, error) {
	type embed VoiceDesignPreviewResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
