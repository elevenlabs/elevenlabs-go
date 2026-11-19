package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingAudioOutputError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (d *DubbingAudioOutputError) GetError() (value string) {
	if d == nil {
		return
	}
	return d.Error
}

func (d *DubbingAudioOutputError) SetError(error string) {
	d.Error = error
}

func (d *DubbingAudioOutputError) MessageType() string {
	return d.messageType
}

func (d *DubbingAudioOutputError) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingAudioOutputError) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingAudioOutputError to string"
	}
	return string(jsonData)
}

func (d *DubbingAudioOutputError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingAudioOutputError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingAudioOutputError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "audio_output_error" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "audio_output_error", *literals.MessageType)
	}
	d.messageType = "audio_output_error"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingAudioOutputError) MarshalJSON() ([]byte, error) {
	type embed DubbingAudioOutputError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "audio_output_error",
	}
	return json.Marshal(marshaler)
}
