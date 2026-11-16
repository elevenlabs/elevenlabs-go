package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EndCallToolResultModel struct {
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	Reason          *string `json:"reason,omitempty" xml:"reason,omitempty"`
	Message         *string `json:"message,omitempty" xml:"message,omitempty"`
	extraProperties map[string]interface{}
}

func (e *EndCallToolResultModel) GetStatus() (value *string) {
	if e == nil {
		return
	}
	return e.Status
}

func (e *EndCallToolResultModel) GetReason() (value *string) {
	if e == nil {
		return
	}
	return e.Reason
}

func (e *EndCallToolResultModel) GetMessage() (value *string) {
	if e == nil {
		return
	}
	return e.Message
}

func (e *EndCallToolResultModel) SetStatus(status *string) {
	e.Status = status
}

func (e *EndCallToolResultModel) SetReason(reason *string) {
	e.Reason = reason
}

func (e *EndCallToolResultModel) SetMessage(message *string) {
	e.Message = message
}

func (e *EndCallToolResultModel) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EndCallToolResultModel) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EndCallToolResultModel to string"
	}
	return string(jsonData)
}

func (e *EndCallToolResultModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, e); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "result_type")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EndCallToolResultModel) MarshalJSON() ([]byte, error) {
	type embed EndCallToolResultModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
