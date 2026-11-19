package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingQueueOverflowError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (d *DubbingQueueOverflowError) GetError() (value string) {
	if d == nil {
		return
	}
	return d.Error
}

func (d *DubbingQueueOverflowError) SetError(error string) {
	d.Error = error
}

func (d *DubbingQueueOverflowError) MessageType() string {
	return d.messageType
}

func (d *DubbingQueueOverflowError) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingQueueOverflowError) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingQueueOverflowError to string"
	}
	return string(jsonData)
}

func (d *DubbingQueueOverflowError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingQueueOverflowError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingQueueOverflowError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "queue_overflow" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "queue_overflow", *literals.MessageType)
	}
	d.messageType = "queue_overflow"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingQueueOverflowError) MarshalJSON() ([]byte, error) {
	type embed DubbingQueueOverflowError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "queue_overflow",
	}
	return json.Marshal(marshaler)
}
