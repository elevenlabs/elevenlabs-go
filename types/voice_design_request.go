package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceDesignRequest struct {
	// Description to use for the created voice.
	VoiceDescription string `json:"voice_description" xml:"voice_description" required:"true" maxLength:"1000" minLength:"20"`
	// Text to generate, text length has to be between 100 and 1000.
	Text *string `json:"text,omitempty" xml:"text,omitempty" maxLength:"1000" minLength:"100"`
	// Whether to automatically generate a text suitable for the voice description.
	AutoGenerateText *bool `json:"auto_generate_text,omitempty" xml:"auto_generate_text,omitempty"`
	// Controls the volume level of the generated voice. -1 is quietest, 1 is loudest, 0 corresponds to roughly -24 LUFS.
	Loudness *float64 `json:"loudness,omitempty" xml:"loudness,omitempty" min:"-1" max:"1"`
	// Higher quality results in better voice output but less variety.
	Quality *float64 `json:"quality,omitempty" xml:"quality,omitempty" min:"-1" max:"1"`
	// Random number that controls the voice generation. Same seed with same inputs produces same voice.
	Seed *int `json:"seed,omitempty" xml:"seed,omitempty" min:"0" max:"2147483647"`
	// Controls how closely the AI follows the prompt. Lower numbers give the AI more freedom to be creative, while higher numbers force it to stick more to the prompt. High numbers can cause voice to sound artificial or robotic. We recommend to use longer, more detailed prompts at lower Guidance Scale.
	GuidanceScale *float64 `json:"guidance_scale,omitempty" xml:"guidance_scale,omitempty" min:"0" max:"100"`
	// Whether to enhance the voice description using AI to add more detail and improve voice generation quality. When enabled, the system will automatically expand simple prompts into more detailed voice descriptions. Defaults to False
	ShouldEnhance   *bool `json:"should_enhance,omitempty" xml:"should_enhance,omitempty"`
	extraProperties map[string]interface{}
}

func (v *VoiceDesignRequest) GetVoiceDescription() (value string) {
	if v == nil {
		return
	}
	return v.VoiceDescription
}

func (v *VoiceDesignRequest) GetText() (value *string) {
	if v == nil {
		return
	}
	return v.Text
}

func (v *VoiceDesignRequest) GetAutoGenerateText() (value *bool) {
	if v == nil {
		return
	}
	return v.AutoGenerateText
}

func (v *VoiceDesignRequest) GetLoudness() (value *float64) {
	if v == nil {
		return
	}
	return v.Loudness
}

func (v *VoiceDesignRequest) GetQuality() (value *float64) {
	if v == nil {
		return
	}
	return v.Quality
}

func (v *VoiceDesignRequest) GetSeed() (value *int) {
	if v == nil {
		return
	}
	return v.Seed
}

func (v *VoiceDesignRequest) GetGuidanceScale() (value *float64) {
	if v == nil {
		return
	}
	return v.GuidanceScale
}

func (v *VoiceDesignRequest) GetShouldEnhance() (value *bool) {
	if v == nil {
		return
	}
	return v.ShouldEnhance
}

func (v *VoiceDesignRequest) SetVoiceDescription(voiceDescription string) {
	v.VoiceDescription = voiceDescription
}

func (v *VoiceDesignRequest) SetText(text *string) {
	v.Text = text
}

func (v *VoiceDesignRequest) SetAutoGenerateText(autoGenerateText *bool) {
	v.AutoGenerateText = autoGenerateText
}

func (v *VoiceDesignRequest) SetLoudness(loudness *float64) {
	v.Loudness = loudness
}

func (v *VoiceDesignRequest) SetQuality(quality *float64) {
	v.Quality = quality
}

func (v *VoiceDesignRequest) SetSeed(seed *int) {
	v.Seed = seed
}

func (v *VoiceDesignRequest) SetGuidanceScale(guidanceScale *float64) {
	v.GuidanceScale = guidanceScale
}

func (v *VoiceDesignRequest) SetShouldEnhance(shouldEnhance *bool) {
	v.ShouldEnhance = shouldEnhance
}

func (v *VoiceDesignRequest) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoiceDesignRequest) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceDesignRequest to string"
	}
	return string(jsonData)
}

func (v *VoiceDesignRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, v); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VoiceDesignRequest) MarshalJSON() ([]byte, error) {
	type embed VoiceDesignRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
