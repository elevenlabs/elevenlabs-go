package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// An error returned in place of an edited transcript when the edit could not be produced.
type TranscriptEditError struct {
	kind string
	// edit_failed: the edit could not be produced.
	errorType string
	// A short, user-facing explanation of the failure.
	Message         string `json:"message" xml:"message" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TranscriptEditError) GetMessage() (value string) {
	if t == nil {
		return
	}
	return t.Message
}

func (t *TranscriptEditError) SetMessage(message string) {
	t.Message = message
}

func (t *TranscriptEditError) Kind() string {
	return t.kind
}

func (t *TranscriptEditError) ErrorType() string {
	return t.errorType
}

func (t *TranscriptEditError) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranscriptEditError) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranscriptEditError to string"
	}
	return string(jsonData)
}

func (t *TranscriptEditError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranscriptEditError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranscriptEditError(tmp)
	var literals struct {
		Kind      *string `json:"kind"`
		ErrorType *string `json:"error_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Kind != nil && *literals.Kind != "error" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "error", *literals.Kind)
	}
	if literals.ErrorType != nil && *literals.ErrorType != "edit_failed" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "edit_failed", *literals.ErrorType)
	}
	t.kind = "error"
	t.errorType = "edit_failed"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "kind", "error_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranscriptEditError) MarshalJSON() ([]byte, error) {
	type embed TranscriptEditError
	var marshaler = struct {
		embed
		Kind      string `json:"kind"`
		ErrorType string `json:"error_type"`
	}{
		embed:     embed(*t),
		Kind:      "error",
		ErrorType: "edit_failed",
	}
	return json.Marshal(marshaler)
}
