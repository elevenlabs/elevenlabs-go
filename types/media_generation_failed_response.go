package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A failed media generation and why it failed.
type MediaGenerationFailedResponse struct {
	// The unique identifier of the generation.
	Id string `json:"id" xml:"id" required:"true"`
	// The lifecycle status of the generation.
	status string
	// The category of failure.
	FailureReason MediaGenerationFailedResponseFailureReason `json:"failure_reason" xml:"failure_reason" required:"true"`
	// A human-readable description of the failure. Failed generations are not charged.
	ErrorMessage    string `json:"error_message" xml:"error_message" required:"true"`
	extraProperties map[string]interface{}
}

func (m *MediaGenerationFailedResponse) GetId() (value string) {
	if m == nil {
		return
	}
	return m.Id
}

func (m *MediaGenerationFailedResponse) GetFailureReason() (value MediaGenerationFailedResponseFailureReason) {
	if m == nil {
		return
	}
	return m.FailureReason
}

func (m *MediaGenerationFailedResponse) GetErrorMessage() (value string) {
	if m == nil {
		return
	}
	return m.ErrorMessage
}

func (m *MediaGenerationFailedResponse) SetId(id string) {
	m.Id = id
}

func (m *MediaGenerationFailedResponse) SetFailureReason(failureReason MediaGenerationFailedResponseFailureReason) {
	m.FailureReason = failureReason
}

func (m *MediaGenerationFailedResponse) SetErrorMessage(errorMessage string) {
	m.ErrorMessage = errorMessage
}

func (m *MediaGenerationFailedResponse) Status() string {
	return m.status
}

func (m *MediaGenerationFailedResponse) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MediaGenerationFailedResponse) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MediaGenerationFailedResponse to string"
	}
	return string(jsonData)
}

func (m *MediaGenerationFailedResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias MediaGenerationFailedResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = MediaGenerationFailedResponse(tmp)
	var literals struct {
		Status *string `json:"status"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Status != nil && *literals.Status != "failed" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", m, "failed", *literals.Status)
	}
	m.status = "failed"
	extraProperties, err := internal.ExtractExtraProperties(data, *m, "status")
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MediaGenerationFailedResponse) MarshalJSON() ([]byte, error) {
	type embed MediaGenerationFailedResponse
	var marshaler = struct {
		embed
		Status string `json:"status"`
	}{
		embed:  embed(*m),
		Status: "failed",
	}
	return json.Marshal(marshaler)
}
