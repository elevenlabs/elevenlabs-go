package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslateRateLimitedError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TranslateRateLimitedError) GetError() (value string) {
	if t == nil {
		return
	}
	return t.Error
}

func (t *TranslateRateLimitedError) SetError(error string) {
	t.Error = error
}

func (t *TranslateRateLimitedError) MessageType() string {
	return t.messageType
}

func (t *TranslateRateLimitedError) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslateRateLimitedError) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslateRateLimitedError to string"
	}
	return string(jsonData)
}

func (t *TranslateRateLimitedError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslateRateLimitedError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslateRateLimitedError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "rate_limited" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "rate_limited", *literals.MessageType)
	}
	t.messageType = "rate_limited"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslateRateLimitedError) MarshalJSON() ([]byte, error) {
	type embed TranslateRateLimitedError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "rate_limited",
	}
	return json.Marshal(marshaler)
}
