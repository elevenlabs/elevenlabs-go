package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingEndOfStream struct {
	messageType     string
	extraProperties map[string]interface{}
}

func (d *DubbingEndOfStream) MessageType() string {
	return d.messageType
}

func (d *DubbingEndOfStream) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingEndOfStream) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingEndOfStream to string"
	}
	return string(jsonData)
}

func (d *DubbingEndOfStream) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingEndOfStream
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingEndOfStream(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "end_of_stream" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "end_of_stream", *literals.MessageType)
	}
	d.messageType = "end_of_stream"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingEndOfStream) MarshalJSON() ([]byte, error) {
	type embed DubbingEndOfStream
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "end_of_stream",
	}
	return json.Marshal(marshaler)
}
