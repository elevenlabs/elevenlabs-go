package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// An `integer` output, delivered as `content` once completed.
type TemplateIntegerOutput struct {
	// The id of the generation behind this output. Pass it as a `generation` reference to use the output as an input elsewhere.
	Id string `json:"id" xml:"id" required:"true"`
	// Where a run, or one of its outputs, is in its lifecycle. The run's is rolled
	// up from its outputs' (``aggregate_output_statuses``). A class rather than a
	// ``Literal`` so it publishes as one named enum instead of being inlined into
	// every kind.
	Status TemplateRunStatus `json:"status" xml:"status" required:"true"`
	// The category of failure. Present only when `status` is `failed`.
	FailureReason *TemplateIntegerOutputFailureReason `json:"failure_reason,omitempty" xml:"failure_reason,omitempty"`
	// A human-readable description of the failure. Present only when `status` is `failed`. Failed generations are not charged.
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// The kind of output: the `type` of the port's `content_schema`.
	type_ string
	// The generated integer. Present only when `status` is `completed`. Reserved: no template produces this kind of output yet; it is published so that templates which do can be run with the same client.
	Content         *int `json:"content,omitempty" xml:"content,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TemplateIntegerOutput) GetId() (value string) {
	if t == nil {
		return
	}
	return t.Id
}

func (t *TemplateIntegerOutput) GetStatus() (value TemplateRunStatus) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TemplateIntegerOutput) GetFailureReason() (value *TemplateIntegerOutputFailureReason) {
	if t == nil {
		return
	}
	return t.FailureReason
}

func (t *TemplateIntegerOutput) GetErrorMessage() (value *string) {
	if t == nil {
		return
	}
	return t.ErrorMessage
}

func (t *TemplateIntegerOutput) GetContent() (value *int) {
	if t == nil {
		return
	}
	return t.Content
}

func (t *TemplateIntegerOutput) SetId(id string) {
	t.Id = id
}

func (t *TemplateIntegerOutput) SetStatus(status TemplateRunStatus) {
	t.Status = status
}

func (t *TemplateIntegerOutput) SetFailureReason(failureReason *TemplateIntegerOutputFailureReason) {
	t.FailureReason = failureReason
}

func (t *TemplateIntegerOutput) SetErrorMessage(errorMessage *string) {
	t.ErrorMessage = errorMessage
}

func (t *TemplateIntegerOutput) SetContent(content *int) {
	t.Content = content
}

func (t *TemplateIntegerOutput) Type() string {
	return t.type_
}

func (t *TemplateIntegerOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TemplateIntegerOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateIntegerOutput to string"
	}
	return string(jsonData)
}

func (t *TemplateIntegerOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "integer" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "integer", *literals.Type)
	}
	t.type_ = "integer"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TemplateIntegerOutput) MarshalJSON() ([]byte, error) {
	type embed TemplateIntegerOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*t),
		Type:  "integer",
	}
	return json.Marshal(marshaler)
}
