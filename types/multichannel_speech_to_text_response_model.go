package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Response model for multichannel speech-to-text transcription.
type MultichannelSpeechToTextResponseModel struct {
	// List of transcripts, one for each audio channel. Each transcript contains the text and word-level details for its respective channel.
	Transcripts []*SpeechToTextChunkResponseModel `json:"transcripts" xml:"transcripts" required:"true"`
	// The transcription ID of the response.
	TranscriptionId *string `json:"transcription_id,omitempty" xml:"transcription_id,omitempty"`
	// The duration of the audio that was transcribed across all channels in seconds.
	AudioDurationSecs *float64 `json:"audio_duration_secs,omitempty" xml:"audio_duration_secs,omitempty"`
	extraProperties   map[string]interface{}
}

func (m *MultichannelSpeechToTextResponseModel) GetTranscripts() (value []*SpeechToTextChunkResponseModel) {
	if m == nil {
		return
	}
	return m.Transcripts
}

func (m *MultichannelSpeechToTextResponseModel) GetTranscriptionId() (value *string) {
	if m == nil {
		return
	}
	return m.TranscriptionId
}

func (m *MultichannelSpeechToTextResponseModel) GetAudioDurationSecs() (value *float64) {
	if m == nil {
		return
	}
	return m.AudioDurationSecs
}

func (m *MultichannelSpeechToTextResponseModel) SetTranscripts(transcripts []*SpeechToTextChunkResponseModel) {
	m.Transcripts = transcripts
}

func (m *MultichannelSpeechToTextResponseModel) SetTranscriptionId(transcriptionId *string) {
	m.TranscriptionId = transcriptionId
}

func (m *MultichannelSpeechToTextResponseModel) SetAudioDurationSecs(audioDurationSecs *float64) {
	m.AudioDurationSecs = audioDurationSecs
}

func (m *MultichannelSpeechToTextResponseModel) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MultichannelSpeechToTextResponseModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MultichannelSpeechToTextResponseModel to string"
	}
	return string(jsonData)
}

func (m *MultichannelSpeechToTextResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MultichannelSpeechToTextResponseModel) MarshalJSON() ([]byte, error) {
	type embed MultichannelSpeechToTextResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
