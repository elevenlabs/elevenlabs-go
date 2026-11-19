package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslateAudio struct {
	messageType     string
	Data            string `json:"data" xml:"data" required:"true"`
	SampleRate      int    `json:"sample_rate" xml:"sample_rate" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TranslateAudio) GetData() (value string) {
	if t == nil {
		return
	}
	return t.Data
}

func (t *TranslateAudio) GetSampleRate() (value int) {
	if t == nil {
		return
	}
	return t.SampleRate
}

func (t *TranslateAudio) SetData(data string) {
	t.Data = data
}

func (t *TranslateAudio) SetSampleRate(sampleRate int) {
	t.SampleRate = sampleRate
}

func (t *TranslateAudio) MessageType() string {
	return t.messageType
}

func (t *TranslateAudio) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslateAudio) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslateAudio to string"
	}
	return string(jsonData)
}

func (t *TranslateAudio) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslateAudio
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslateAudio(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "audio" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "audio", *literals.MessageType)
	}
	t.messageType = "audio"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslateAudio) MarshalJSON() ([]byte, error) {
	type embed TranslateAudio
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "audio",
	}
	return json.Marshal(marshaler)
}
