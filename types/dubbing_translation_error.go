package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingTranslationError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (d *DubbingTranslationError) GetError() (value string) {
	if d == nil {
		return
	}
	return d.Error
}

func (d *DubbingTranslationError) SetError(error string) {
	d.Error = error
}

func (d *DubbingTranslationError) MessageType() string {
	return d.messageType
}

func (d *DubbingTranslationError) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingTranslationError) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingTranslationError to string"
	}
	return string(jsonData)
}

func (d *DubbingTranslationError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingTranslationError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingTranslationError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "translation_error" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "translation_error", *literals.MessageType)
	}
	d.messageType = "translation_error"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingTranslationError) MarshalJSON() ([]byte, error) {
	type embed DubbingTranslationError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "translation_error",
	}
	return json.Marshal(marshaler)
}
