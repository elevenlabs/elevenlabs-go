package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslateSessionStarted struct {
	messageType     string
	SessionId       string  `json:"session_id" xml:"session_id" required:"true"`
	ClientSessionId *string `json:"client_session_id,omitempty" xml:"client_session_id,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TranslateSessionStarted) GetSessionId() (value string) {
	if t == nil {
		return
	}
	return t.SessionId
}

func (t *TranslateSessionStarted) GetClientSessionId() (value *string) {
	if t == nil {
		return
	}
	return t.ClientSessionId
}

func (t *TranslateSessionStarted) SetSessionId(sessionId string) {
	t.SessionId = sessionId
}

func (t *TranslateSessionStarted) SetClientSessionId(clientSessionId *string) {
	t.ClientSessionId = clientSessionId
}

func (t *TranslateSessionStarted) MessageType() string {
	return t.messageType
}

func (t *TranslateSessionStarted) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslateSessionStarted) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslateSessionStarted to string"
	}
	return string(jsonData)
}

func (t *TranslateSessionStarted) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "session_started" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "session_started", *literals.MessageType)
	}
	t.messageType = "session_started"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslateSessionStarted) MarshalJSON() ([]byte, error) {
	type embed TranslateSessionStarted
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "session_started",
	}
	return json.Marshal(marshaler)
}
