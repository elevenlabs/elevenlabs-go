package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingTranscriberError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (d *DubbingTranscriberError) GetError() (value string) {
	if d == nil {
		return
	}
	return d.Error
}

func (d *DubbingTranscriberError) SetError(error string) {
	d.Error = error
}

func (d *DubbingTranscriberError) MessageType() string {
	return d.messageType
}

func (d *DubbingTranscriberError) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingTranscriberError) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingTranscriberError to string"
	}
	return string(jsonData)
}

func (d *DubbingTranscriberError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingTranscriberError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingTranscriberError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "transcriber_error" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "transcriber_error", *literals.MessageType)
	}
	d.messageType = "transcriber_error"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingTranscriberError) MarshalJSON() ([]byte, error) {
	type embed DubbingTranscriberError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "transcriber_error",
	}
	return json.Marshal(marshaler)
}
