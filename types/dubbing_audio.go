package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingAudio struct {
	messageType     string
	Data            string `json:"data" xml:"data" required:"true" minLength:"1"`
	SampleRate      int    `json:"sample_rate" xml:"sample_rate" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingAudio) GetData() (value string) {
	if d == nil {
		return
	}
	return d.Data
}

func (d *DubbingAudio) GetSampleRate() (value int) {
	if d == nil {
		return
	}
	return d.SampleRate
}

func (d *DubbingAudio) SetData(data string) {
	d.Data = data
}

func (d *DubbingAudio) SetSampleRate(sampleRate int) {
	d.SampleRate = sampleRate
}

func (d *DubbingAudio) MessageType() string {
	return d.messageType
}

func (d *DubbingAudio) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingAudio) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingAudio to string"
	}
	return string(jsonData)
}

func (d *DubbingAudio) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingAudio
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingAudio(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "audio" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "audio", *literals.MessageType)
	}
	d.messageType = "audio"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingAudio) MarshalJSON() ([]byte, error) {
	type embed DubbingAudio
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "audio",
	}
	return json.Marshal(marshaler)
}
