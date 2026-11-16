package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslateAuthError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TranslateAuthError) GetError() (value string) {
	if t == nil {
		return
	}
	return t.Error
}

func (t *TranslateAuthError) SetError(error string) {
	t.Error = error
}

func (t *TranslateAuthError) MessageType() string {
	return t.messageType
}

func (t *TranslateAuthError) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslateAuthError) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslateAuthError to string"
	}
	return string(jsonData)
}

func (t *TranslateAuthError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslateAuthError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslateAuthError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "auth_error" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "auth_error", *literals.MessageType)
	}
	t.messageType = "auth_error"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslateAuthError) MarshalJSON() ([]byte, error) {
	type embed TranslateAuthError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "auth_error",
	}
	return json.Marshal(marshaler)
}
