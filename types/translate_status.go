package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslateStatus struct {
	messageType     string
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TranslateStatus) GetStatus() (value string) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TranslateStatus) SetStatus(status string) {
	t.Status = status
}

func (t *TranslateStatus) MessageType() string {
	return t.messageType
}

func (t *TranslateStatus) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslateStatus) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslateStatus to string"
	}
	return string(jsonData)
}

func (t *TranslateStatus) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslateStatus
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslateStatus(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "status" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "status", *literals.MessageType)
	}
	t.messageType = "status"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslateStatus) MarshalJSON() ([]byte, error) {
	type embed TranslateStatus
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "status",
	}
	return json.Marshal(marshaler)
}
