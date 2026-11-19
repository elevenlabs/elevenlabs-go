package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingRateLimitedError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (d *DubbingRateLimitedError) GetError() (value string) {
	if d == nil {
		return
	}
	return d.Error
}

func (d *DubbingRateLimitedError) SetError(error string) {
	d.Error = error
}

func (d *DubbingRateLimitedError) MessageType() string {
	return d.messageType
}

func (d *DubbingRateLimitedError) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingRateLimitedError) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingRateLimitedError to string"
	}
	return string(jsonData)
}

func (d *DubbingRateLimitedError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingRateLimitedError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingRateLimitedError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "rate_limited" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "rate_limited", *literals.MessageType)
	}
	d.messageType = "rate_limited"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingRateLimitedError) MarshalJSON() ([]byte, error) {
	type embed DubbingRateLimitedError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "rate_limited",
	}
	return json.Marshal(marshaler)
}
