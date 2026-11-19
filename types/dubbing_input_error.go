package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingInputError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (d *DubbingInputError) GetError() (value string) {
	if d == nil {
		return
	}
	return d.Error
}

func (d *DubbingInputError) SetError(error string) {
	d.Error = error
}

func (d *DubbingInputError) MessageType() string {
	return d.messageType
}

func (d *DubbingInputError) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingInputError) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingInputError to string"
	}
	return string(jsonData)
}

func (d *DubbingInputError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingInputError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingInputError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "input_error" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "input_error", *literals.MessageType)
	}
	d.messageType = "input_error"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingInputError) MarshalJSON() ([]byte, error) {
	type embed DubbingInputError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "input_error",
	}
	return json.Marshal(marshaler)
}
