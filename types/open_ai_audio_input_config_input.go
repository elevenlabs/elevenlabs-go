package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OpenAiAudioInputConfigInput struct {
	Format        *OpenAiAudioInputFormat        `json:"format,omitempty" xml:"format,omitempty"`
	Transcription *OpenAiAudioInputTranscription `json:"transcription,omitempty" xml:"transcription,omitempty"`
	TurnDetection *OpenAiTurnDetection           `json:"turn_detection,omitempty" xml:"turn_detection,omitempty"`
	// ``BURSTY`` clients stream ahead of wall time, so decisions use the audio clock.
	Pacing          *AudioInputPacing `json:"pacing,omitempty" xml:"pacing,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OpenAiAudioInputConfigInput) GetFormat() (value *OpenAiAudioInputFormat) {
	if o == nil {
		return
	}
	return o.Format
}

func (o *OpenAiAudioInputConfigInput) GetTranscription() (value *OpenAiAudioInputTranscription) {
	if o == nil {
		return
	}
	return o.Transcription
}

func (o *OpenAiAudioInputConfigInput) GetTurnDetection() (value *OpenAiTurnDetection) {
	if o == nil {
		return
	}
	return o.TurnDetection
}

func (o *OpenAiAudioInputConfigInput) GetPacing() (value *AudioInputPacing) {
	if o == nil {
		return
	}
	return o.Pacing
}

func (o *OpenAiAudioInputConfigInput) SetFormat(format *OpenAiAudioInputFormat) {
	o.Format = format
}

func (o *OpenAiAudioInputConfigInput) SetTranscription(transcription *OpenAiAudioInputTranscription) {
	o.Transcription = transcription
}

func (o *OpenAiAudioInputConfigInput) SetTurnDetection(turnDetection *OpenAiTurnDetection) {
	o.TurnDetection = turnDetection
}

func (o *OpenAiAudioInputConfigInput) SetPacing(pacing *AudioInputPacing) {
	o.Pacing = pacing
}

func (o *OpenAiAudioInputConfigInput) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiAudioInputConfigInput) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiAudioInputConfigInput to string"
	}
	return string(jsonData)
}

func (o *OpenAiAudioInputConfigInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, o); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OpenAiAudioInputConfigInput) MarshalJSON() ([]byte, error) {
	type embed OpenAiAudioInputConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
