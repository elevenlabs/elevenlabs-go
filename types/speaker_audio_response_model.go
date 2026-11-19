package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SpeakerAudioResponseModel struct {
	// The base64 encoded audio.
	AudioBase64 string `json:"audio_base_64" xml:"audio_base_64" required:"true"`
	// The media type of the audio.
	MediaType string `json:"media_type" xml:"media_type" required:"true"`
	// The duration of the audio in seconds.
	DurationSecs    float64 `json:"duration_secs" xml:"duration_secs" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SpeakerAudioResponseModel) GetAudioBase64() (value string) {
	if s == nil {
		return
	}
	return s.AudioBase64
}

func (s *SpeakerAudioResponseModel) GetMediaType() (value string) {
	if s == nil {
		return
	}
	return s.MediaType
}

func (s *SpeakerAudioResponseModel) GetDurationSecs() (value float64) {
	if s == nil {
		return
	}
	return s.DurationSecs
}

func (s *SpeakerAudioResponseModel) SetAudioBase64(audioBase64 string) {
	s.AudioBase64 = audioBase64
}

func (s *SpeakerAudioResponseModel) SetMediaType(mediaType string) {
	s.MediaType = mediaType
}

func (s *SpeakerAudioResponseModel) SetDurationSecs(durationSecs float64) {
	s.DurationSecs = durationSecs
}

func (s *SpeakerAudioResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeakerAudioResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeakerAudioResponseModel to string"
	}
	return string(jsonData)
}

func (s *SpeakerAudioResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SpeakerAudioResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SpeakerAudioResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SpeakerAudioResponseModel) MarshalJSON() ([]byte, error) {
	type embed SpeakerAudioResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
