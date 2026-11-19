package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PartialTranscript struct {
	MessageType *string `json:"message_type,omitempty" xml:"message_type,omitempty"`
	Text        string  `json:"text" xml:"text" required:"true"`
	// Probability that the speaker has finished their turn. Only sent for models that predict turns.
	TurnProbability *float64 `json:"turn_probability,omitempty" xml:"turn_probability,omitempty"`
	// The language of the speech so far in the current turn. Only present when include_language_detection is true, and then always present: null until the text is long enough to tell and the same language has been seen on consecutive partials. The committed transcript's language_code is authoritative.
	LanguageCode *string `json:"language_code,omitempty" xml:"language_code,omitempty"`
	// Draft words with speaker IDs when diarize is enabled. Speaker IDs and timestamps may change as more audio is received.
	Words           []*PartialTranscriptWordsItem `json:"words,omitempty" xml:"words,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PartialTranscript) GetMessageType() (value *string) {
	if p == nil {
		return
	}
	return p.MessageType
}

func (p *PartialTranscript) GetText() (value string) {
	if p == nil {
		return
	}
	return p.Text
}

func (p *PartialTranscript) GetTurnProbability() (value *float64) {
	if p == nil {
		return
	}
	return p.TurnProbability
}

func (p *PartialTranscript) GetLanguageCode() (value *string) {
	if p == nil {
		return
	}
	return p.LanguageCode
}

func (p *PartialTranscript) GetWords() (value []*PartialTranscriptWordsItem) {
	if p == nil {
		return
	}
	return p.Words
}

func (p *PartialTranscript) SetMessageType(messageType *string) {
	p.MessageType = messageType
}

func (p *PartialTranscript) SetText(text string) {
	p.Text = text
}

func (p *PartialTranscript) SetTurnProbability(turnProbability *float64) {
	p.TurnProbability = turnProbability
}

func (p *PartialTranscript) SetLanguageCode(languageCode *string) {
	p.LanguageCode = languageCode
}

func (p *PartialTranscript) SetWords(words []*PartialTranscriptWordsItem) {
	p.Words = words
}

func (p *PartialTranscript) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PartialTranscript) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PartialTranscript to string"
	}
	return string(jsonData)
}

func (p *PartialTranscript) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PartialTranscript) MarshalJSON() ([]byte, error) {
	type embed PartialTranscript
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
