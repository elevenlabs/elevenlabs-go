package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslateUnacceptedTermsError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TranslateUnacceptedTermsError) GetError() (value string) {
	if t == nil {
		return
	}
	return t.Error
}

func (t *TranslateUnacceptedTermsError) SetError(error string) {
	t.Error = error
}

func (t *TranslateUnacceptedTermsError) MessageType() string {
	return t.messageType
}

func (t *TranslateUnacceptedTermsError) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslateUnacceptedTermsError) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslateUnacceptedTermsError to string"
	}
	return string(jsonData)
}

func (t *TranslateUnacceptedTermsError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslateUnacceptedTermsError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslateUnacceptedTermsError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "unaccepted_terms" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "unaccepted_terms", *literals.MessageType)
	}
	t.messageType = "unaccepted_terms"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslateUnacceptedTermsError) MarshalJSON() ([]byte, error) {
	type embed TranslateUnacceptedTermsError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "unaccepted_terms",
	}
	return json.Marshal(marshaler)
}
