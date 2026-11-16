package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingSessionStarted struct {
	messageType     string
	SessionId       string  `json:"session_id" xml:"session_id" required:"true" minLength:"1"`
	ClientSessionId *string `json:"client_session_id,omitempty" xml:"client_session_id,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DubbingSessionStarted) GetSessionId() (value string) {
	if d == nil {
		return
	}
	return d.SessionId
}

func (d *DubbingSessionStarted) GetClientSessionId() (value *string) {
	if d == nil {
		return
	}
	return d.ClientSessionId
}

func (d *DubbingSessionStarted) SetSessionId(sessionId string) {
	d.SessionId = sessionId
}

func (d *DubbingSessionStarted) SetClientSessionId(clientSessionId *string) {
	d.ClientSessionId = clientSessionId
}

func (d *DubbingSessionStarted) MessageType() string {
	return d.messageType
}

func (d *DubbingSessionStarted) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingSessionStarted) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingSessionStarted to string"
	}
	return string(jsonData)
}

func (d *DubbingSessionStarted) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "session_started" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "session_started", *literals.MessageType)
	}
	d.messageType = "session_started"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingSessionStarted) MarshalJSON() ([]byte, error) {
	type embed DubbingSessionStarted
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "session_started",
	}
	return json.Marshal(marshaler)
}
