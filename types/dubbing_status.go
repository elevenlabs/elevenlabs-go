package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingStatus struct {
	messageType     string
	Status          DubbingStatusStatus `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingStatus) GetStatus() (value DubbingStatusStatus) {
	if d == nil {
		return
	}
	return d.Status
}

func (d *DubbingStatus) SetStatus(status DubbingStatusStatus) {
	d.Status = status
}

func (d *DubbingStatus) MessageType() string {
	return d.messageType
}

func (d *DubbingStatus) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingStatus) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingStatus to string"
	}
	return string(jsonData)
}

func (d *DubbingStatus) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingStatus
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingStatus(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "status" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "status", *literals.MessageType)
	}
	d.messageType = "status"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingStatus) MarshalJSON() ([]byte, error) {
	type embed DubbingStatus
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "status",
	}
	return json.Marshal(marshaler)
}
