package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AssetTranscriptionData struct {
	LanguageCode     string    `json:"language_code" xml:"language_code" required:"true"`
	Text             string    `json:"text" xml:"text" required:"true"`
	Words            []string  `json:"words" xml:"words" required:"true"`
	WordStartTimesMs []int     `json:"word_start_times_ms" xml:"word_start_times_ms" required:"true"`
	WordEndTimesMs   []int     `json:"word_end_times_ms" xml:"word_end_times_ms" required:"true"`
	WordSpeakerIds   []*string `json:"word_speaker_ids" xml:"word_speaker_ids" required:"true"`
	extraProperties  map[string]interface{}
}

func (a *AssetTranscriptionData) GetLanguageCode() (value string) {
	if a == nil {
		return
	}
	return a.LanguageCode
}

func (a *AssetTranscriptionData) GetText() (value string) {
	if a == nil {
		return
	}
	return a.Text
}

func (a *AssetTranscriptionData) GetWords() (value []string) {
	if a == nil {
		return
	}
	return a.Words
}

func (a *AssetTranscriptionData) GetWordStartTimesMs() (value []int) {
	if a == nil {
		return
	}
	return a.WordStartTimesMs
}

func (a *AssetTranscriptionData) GetWordEndTimesMs() (value []int) {
	if a == nil {
		return
	}
	return a.WordEndTimesMs
}

func (a *AssetTranscriptionData) GetWordSpeakerIds() (value []*string) {
	if a == nil {
		return
	}
	return a.WordSpeakerIds
}

func (a *AssetTranscriptionData) SetLanguageCode(languageCode string) {
	a.LanguageCode = languageCode
}

func (a *AssetTranscriptionData) SetText(text string) {
	a.Text = text
}

func (a *AssetTranscriptionData) SetWords(words []string) {
	a.Words = words
}

func (a *AssetTranscriptionData) SetWordStartTimesMs(wordStartTimesMs []int) {
	a.WordStartTimesMs = wordStartTimesMs
}

func (a *AssetTranscriptionData) SetWordEndTimesMs(wordEndTimesMs []int) {
	a.WordEndTimesMs = wordEndTimesMs
}

func (a *AssetTranscriptionData) SetWordSpeakerIds(wordSpeakerIds []*string) {
	a.WordSpeakerIds = wordSpeakerIds
}

func (a *AssetTranscriptionData) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AssetTranscriptionData) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AssetTranscriptionData to string"
	}
	return string(jsonData)
}

func (a *AssetTranscriptionData) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AssetTranscriptionData
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AssetTranscriptionData(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AssetTranscriptionData) MarshalJSON() ([]byte, error) {
	type embed AssetTranscriptionData
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
