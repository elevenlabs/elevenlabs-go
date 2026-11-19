package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslateEndOfStream struct {
	messageType     string
	extraProperties map[string]interface{}
}

func (t *TranslateEndOfStream) MessageType() string {
	return t.messageType
}

func (t *TranslateEndOfStream) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslateEndOfStream) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslateEndOfStream to string"
	}
	return string(jsonData)
}

func (t *TranslateEndOfStream) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslateEndOfStream
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslateEndOfStream(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "end_of_stream" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "end_of_stream", *literals.MessageType)
	}
	t.messageType = "end_of_stream"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslateEndOfStream) MarshalJSON() ([]byte, error) {
	type embed TranslateEndOfStream
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "end_of_stream",
	}
	return json.Marshal(marshaler)
}
