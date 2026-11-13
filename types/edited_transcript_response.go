package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A transcript produced by the optional transcript edit (text only, untimed).
type EditedTranscriptResponse struct {
	kind string
	// The edited transcript text.
	Text string `json:"text" xml:"text" required:"true"`
	// The edited transcript text. Same value as `text`; matches the `edited_text` field of the realtime `edited_transcript` event.
	EditedText      *string `json:"edited_text,omitempty" xml:"edited_text,omitempty"`
	extraProperties map[string]interface{}
}

func (e *EditedTranscriptResponse) GetText() (value string) {
	if e == nil {
		return
	}
	return e.Text
}

func (e *EditedTranscriptResponse) GetEditedText() (value *string) {
	if e == nil {
		return
	}
	return e.EditedText
}

func (e *EditedTranscriptResponse) SetText(text string) {
	e.Text = text
}

func (e *EditedTranscriptResponse) SetEditedText(editedText *string) {
	e.EditedText = editedText
}

func (e *EditedTranscriptResponse) Kind() string {
	return e.kind
}

func (e *EditedTranscriptResponse) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EditedTranscriptResponse) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditedTranscriptResponse to string"
	}
	return string(jsonData)
}

func (e *EditedTranscriptResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias EditedTranscriptResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = EditedTranscriptResponse(tmp)
	var literals struct {
		Kind *string `json:"kind"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Kind != nil && *literals.Kind != "transcript" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", e, "transcript", *literals.Kind)
	}
	e.kind = "transcript"
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "kind")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EditedTranscriptResponse) MarshalJSON() ([]byte, error) {
	type embed EditedTranscriptResponse
	var marshaler = struct {
		embed
		Kind string `json:"kind"`
	}{
		embed: embed(*e),
		Kind:  "transcript",
	}
	return json.Marshal(marshaler)
}
