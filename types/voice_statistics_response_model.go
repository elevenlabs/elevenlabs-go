package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceStatisticsResponseModel struct {
	// The project voice reference ID.
	ProjectVoiceRefId string `json:"project_voice_ref_id" xml:"project_voice_ref_id" required:"true"`
	// The number of unconverted characters for this voice.
	CharactersUnconverted int `json:"characters_unconverted" xml:"characters_unconverted" required:"true"`
	// The number of converted characters for this voice.
	CharactersConverted int `json:"characters_converted" xml:"characters_converted" required:"true"`
	// The number of credits needed to convert the remaining audio for this voice.
	CreditsNeededToConvert *int `json:"credits_needed_to_convert,omitempty" xml:"credits_needed_to_convert,omitempty"`
	// The voice ID.
	VoiceId         *string `json:"voice_id,omitempty" xml:"voice_id,omitempty"`
	extraProperties map[string]interface{}
}

func (v *VoiceStatisticsResponseModel) GetProjectVoiceRefId() (value string) {
	if v == nil {
		return
	}
	return v.ProjectVoiceRefId
}

func (v *VoiceStatisticsResponseModel) GetCharactersUnconverted() (value int) {
	if v == nil {
		return
	}
	return v.CharactersUnconverted
}

func (v *VoiceStatisticsResponseModel) GetCharactersConverted() (value int) {
	if v == nil {
		return
	}
	return v.CharactersConverted
}

func (v *VoiceStatisticsResponseModel) GetCreditsNeededToConvert() (value *int) {
	if v == nil {
		return
	}
	return v.CreditsNeededToConvert
}

func (v *VoiceStatisticsResponseModel) GetVoiceId() (value *string) {
	if v == nil {
		return
	}
	return v.VoiceId
}

func (v *VoiceStatisticsResponseModel) SetProjectVoiceRefId(projectVoiceRefId string) {
	v.ProjectVoiceRefId = projectVoiceRefId
}

func (v *VoiceStatisticsResponseModel) SetCharactersUnconverted(charactersUnconverted int) {
	v.CharactersUnconverted = charactersUnconverted
}

func (v *VoiceStatisticsResponseModel) SetCharactersConverted(charactersConverted int) {
	v.CharactersConverted = charactersConverted
}

func (v *VoiceStatisticsResponseModel) SetCreditsNeededToConvert(creditsNeededToConvert *int) {
	v.CreditsNeededToConvert = creditsNeededToConvert
}

func (v *VoiceStatisticsResponseModel) SetVoiceId(voiceId *string) {
	v.VoiceId = voiceId
}

func (v *VoiceStatisticsResponseModel) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoiceStatisticsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceStatisticsResponseModel to string"
	}
	return string(jsonData)
}

func (v *VoiceStatisticsResponseModel) UnmarshalJSON(data []byte) error {
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

func (v *VoiceStatisticsResponseModel) MarshalJSON() ([]byte, error) {
	type embed VoiceStatisticsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
