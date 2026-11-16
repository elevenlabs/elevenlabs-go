package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Chunk-level detail of the transcription with timing information.
type SpeechToTextChunkResponseModel struct {
	// The detected language code (e.g. 'eng' for English).
	LanguageCode string `json:"language_code" xml:"language_code" required:"true"`
	// The confidence score of the language detection (0 to 1).
	LanguageProbability float64 `json:"language_probability" xml:"language_probability" required:"true"`
	// The raw text of the transcription.
	Text string `json:"text" xml:"text" required:"true"`
	// List of words with their timing information.
	Words []*SpeechToTextWordResponseModel `json:"words" xml:"words" required:"true"`
	// The channel index this transcript belongs to (for multichannel audio).
	ChannelIndex *int `json:"channel_index,omitempty" xml:"channel_index,omitempty"`
	// Requested additional formats of the transcript.
	AdditionalFormats []*AdditionalFormatResponseModel `json:"additional_formats,omitempty" xml:"additional_formats,omitempty"`
	// The transcription ID of the response.
	TranscriptionId *string `json:"transcription_id,omitempty" xml:"transcription_id,omitempty"`
	// List of detected entities with their text, type, and character positions in the transcript.
	Entities []*DetectedEntity `json:"entities,omitempty" xml:"entities,omitempty"`
	// The duration of the audio that was transcribed in seconds.
	AudioDurationSecs *float64 `json:"audio_duration_secs,omitempty" xml:"audio_duration_secs,omitempty"`
	// Result of the optional transcript edit: the edited text, or an error if it could not be produced. Absent when no edit was requested. On success, both `text` and `edited_text` hold the edited transcript; compare them with the top-level `text` to see what changed.
	EditedTranscript *SpeechToTextChunkResponseModelEditedTranscript `json:"edited_transcript,omitempty" xml:"edited_transcript,omitempty"`
	extraProperties  map[string]interface{}
}

func (s *SpeechToTextChunkResponseModel) GetLanguageCode() (value string) {
	if s == nil {
		return
	}
	return s.LanguageCode
}

func (s *SpeechToTextChunkResponseModel) GetLanguageProbability() (value float64) {
	if s == nil {
		return
	}
	return s.LanguageProbability
}

func (s *SpeechToTextChunkResponseModel) GetText() (value string) {
	if s == nil {
		return
	}
	return s.Text
}

func (s *SpeechToTextChunkResponseModel) GetWords() (value []*SpeechToTextWordResponseModel) {
	if s == nil {
		return
	}
	return s.Words
}

func (s *SpeechToTextChunkResponseModel) GetChannelIndex() (value *int) {
	if s == nil {
		return
	}
	return s.ChannelIndex
}

func (s *SpeechToTextChunkResponseModel) GetAdditionalFormats() (value []*AdditionalFormatResponseModel) {
	if s == nil {
		return
	}
	return s.AdditionalFormats
}

func (s *SpeechToTextChunkResponseModel) GetTranscriptionId() (value *string) {
	if s == nil {
		return
	}
	return s.TranscriptionId
}

func (s *SpeechToTextChunkResponseModel) GetEntities() (value []*DetectedEntity) {
	if s == nil {
		return
	}
	return s.Entities
}

func (s *SpeechToTextChunkResponseModel) GetAudioDurationSecs() (value *float64) {
	if s == nil {
		return
	}
	return s.AudioDurationSecs
}

func (s *SpeechToTextChunkResponseModel) GetEditedTranscript() (value *SpeechToTextChunkResponseModelEditedTranscript) {
	if s == nil {
		return
	}
	return s.EditedTranscript
}

func (s *SpeechToTextChunkResponseModel) SetLanguageCode(languageCode string) {
	s.LanguageCode = languageCode
}

func (s *SpeechToTextChunkResponseModel) SetLanguageProbability(languageProbability float64) {
	s.LanguageProbability = languageProbability
}

func (s *SpeechToTextChunkResponseModel) SetText(text string) {
	s.Text = text
}

func (s *SpeechToTextChunkResponseModel) SetWords(words []*SpeechToTextWordResponseModel) {
	s.Words = words
}

func (s *SpeechToTextChunkResponseModel) SetChannelIndex(channelIndex *int) {
	s.ChannelIndex = channelIndex
}

func (s *SpeechToTextChunkResponseModel) SetAdditionalFormats(additionalFormats []*AdditionalFormatResponseModel) {
	s.AdditionalFormats = additionalFormats
}

func (s *SpeechToTextChunkResponseModel) SetTranscriptionId(transcriptionId *string) {
	s.TranscriptionId = transcriptionId
}

func (s *SpeechToTextChunkResponseModel) SetEntities(entities []*DetectedEntity) {
	s.Entities = entities
}

func (s *SpeechToTextChunkResponseModel) SetAudioDurationSecs(audioDurationSecs *float64) {
	s.AudioDurationSecs = audioDurationSecs
}

func (s *SpeechToTextChunkResponseModel) SetEditedTranscript(editedTranscript *SpeechToTextChunkResponseModelEditedTranscript) {
	s.EditedTranscript = editedTranscript
}

func (s *SpeechToTextChunkResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeechToTextChunkResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeechToTextChunkResponseModel to string"
	}
	return string(jsonData)
}

func (s *SpeechToTextChunkResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SpeechToTextChunkResponseModel) MarshalJSON() ([]byte, error) {
	type embed SpeechToTextChunkResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
