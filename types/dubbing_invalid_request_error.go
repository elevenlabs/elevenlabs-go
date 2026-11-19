package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingInvalidRequestError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (d *DubbingInvalidRequestError) GetError() (value string) {
	if d == nil {
		return
	}
	return d.Error
}

func (d *DubbingInvalidRequestError) SetError(error string) {
	d.Error = error
}

func (d *DubbingInvalidRequestError) MessageType() string {
	return d.messageType
}

func (d *DubbingInvalidRequestError) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingInvalidRequestError) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingInvalidRequestError to string"
	}
	return string(jsonData)
}

func (d *DubbingInvalidRequestError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingInvalidRequestError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingInvalidRequestError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "invalid_request" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "invalid_request", *literals.MessageType)
	}
	d.messageType = "invalid_request"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingInvalidRequestError) MarshalJSON() ([]byte, error) {
	type embed DubbingInvalidRequestError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "invalid_request",
	}
	return json.Marshal(marshaler)
}
