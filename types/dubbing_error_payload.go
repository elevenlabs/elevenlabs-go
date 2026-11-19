package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingErrorPayload struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (d *DubbingErrorPayload) GetError() (value string) {
	if d == nil {
		return
	}
	return d.Error
}

func (d *DubbingErrorPayload) SetError(error string) {
	d.Error = error
}

func (d *DubbingErrorPayload) MessageType() string {
	return d.messageType
}

func (d *DubbingErrorPayload) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingErrorPayload) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingErrorPayload to string"
	}
	return string(jsonData)
}

func (d *DubbingErrorPayload) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingErrorPayload
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingErrorPayload(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "error" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "error", *literals.MessageType)
	}
	d.messageType = "error"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingErrorPayload) MarshalJSON() ([]byte, error) {
	type embed DubbingErrorPayload
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "error",
	}
	return json.Marshal(marshaler)
}
