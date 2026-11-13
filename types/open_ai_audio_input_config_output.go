package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type OpenAiAudioInputConfigOutput struct {
	Format        *OpenAiAudioInputFormat        `json:"format,omitempty" xml:"format,omitempty"`
	Transcription *OpenAiAudioInputTranscription `json:"transcription,omitempty" xml:"transcription,omitempty"`
	TurnDetection *OpenAiTurnDetection           `json:"turn_detection,omitempty" xml:"turn_detection,omitempty"`
	// ``BURSTY`` clients stream ahead of wall time, so decisions use the audio clock.
	Pacing          *AudioInputPacing `json:"pacing,omitempty" xml:"pacing,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OpenAiAudioInputConfigOutput) GetFormat() (value *OpenAiAudioInputFormat) {
	if o == nil {
		return
	}
	return o.Format
}

func (o *OpenAiAudioInputConfigOutput) GetTranscription() (value *OpenAiAudioInputTranscription) {
	if o == nil {
		return
	}
	return o.Transcription
}

func (o *OpenAiAudioInputConfigOutput) GetTurnDetection() (value *OpenAiTurnDetection) {
	if o == nil {
		return
	}
	return o.TurnDetection
}

func (o *OpenAiAudioInputConfigOutput) GetPacing() (value *AudioInputPacing) {
	if o == nil {
		return
	}
	return o.Pacing
}

func (o *OpenAiAudioInputConfigOutput) SetFormat(format *OpenAiAudioInputFormat) {
	o.Format = format
}

func (o *OpenAiAudioInputConfigOutput) SetTranscription(transcription *OpenAiAudioInputTranscription) {
	o.Transcription = transcription
}

func (o *OpenAiAudioInputConfigOutput) SetTurnDetection(turnDetection *OpenAiTurnDetection) {
	o.TurnDetection = turnDetection
}

func (o *OpenAiAudioInputConfigOutput) SetPacing(pacing *AudioInputPacing) {
	o.Pacing = pacing
}

func (o *OpenAiAudioInputConfigOutput) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OpenAiAudioInputConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OpenAiAudioInputConfigOutput to string"
	}
	return string(jsonData)
}

func (o *OpenAiAudioInputConfigOutput) UnmarshalJSON(data []byte) error {
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

func (o *OpenAiAudioInputConfigOutput) MarshalJSON() ([]byte, error) {
	type embed OpenAiAudioInputConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
