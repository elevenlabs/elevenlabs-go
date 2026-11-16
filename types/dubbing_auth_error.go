package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingAuthError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (d *DubbingAuthError) GetError() (value string) {
	if d == nil {
		return
	}
	return d.Error
}

func (d *DubbingAuthError) SetError(error string) {
	d.Error = error
}

func (d *DubbingAuthError) MessageType() string {
	return d.messageType
}

func (d *DubbingAuthError) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingAuthError) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingAuthError to string"
	}
	return string(jsonData)
}

func (d *DubbingAuthError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingAuthError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingAuthError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "auth_error" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "auth_error", *literals.MessageType)
	}
	d.messageType = "auth_error"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingAuthError) MarshalJSON() ([]byte, error) {
	type embed DubbingAuthError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "auth_error",
	}
	return json.Marshal(marshaler)
}
