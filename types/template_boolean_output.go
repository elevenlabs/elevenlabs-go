package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A `boolean` output, delivered as `content` once completed.
type TemplateBooleanOutput struct {
	// The id of the generation behind this output. Pass it as a `generation` reference to use the output as an input elsewhere.
	Id string `json:"id" xml:"id" required:"true"`
	// Where a run, or one of its outputs, is in its lifecycle. The run's is rolled
	// up from its outputs' (``aggregate_output_statuses``). A class rather than a
	// ``Literal`` so it publishes as one named enum instead of being inlined into
	// every kind.
	Status TemplateRunStatus `json:"status" xml:"status" required:"true"`
	// The category of failure. Present only when `status` is `failed`.
	FailureReason *TemplateBooleanOutputFailureReason `json:"failure_reason,omitempty" xml:"failure_reason,omitempty"`
	// A human-readable description of the failure. Present only when `status` is `failed`. Failed generations are not charged.
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// The kind of output: the `type` of the port's `content_schema`.
	type_ string
	// The generated boolean. Present only when `status` is `completed`. Reserved: no template produces this kind of output yet; it is published so that templates which do can be run with the same client.
	Content         *bool `json:"content,omitempty" xml:"content,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TemplateBooleanOutput) GetId() (value string) {
	if t == nil {
		return
	}
	return t.Id
}

func (t *TemplateBooleanOutput) GetStatus() (value TemplateRunStatus) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TemplateBooleanOutput) GetFailureReason() (value *TemplateBooleanOutputFailureReason) {
	if t == nil {
		return
	}
	return t.FailureReason
}

func (t *TemplateBooleanOutput) GetErrorMessage() (value *string) {
	if t == nil {
		return
	}
	return t.ErrorMessage
}

func (t *TemplateBooleanOutput) GetContent() (value *bool) {
	if t == nil {
		return
	}
	return t.Content
}

func (t *TemplateBooleanOutput) SetId(id string) {
	t.Id = id
}

func (t *TemplateBooleanOutput) SetStatus(status TemplateRunStatus) {
	t.Status = status
}

func (t *TemplateBooleanOutput) SetFailureReason(failureReason *TemplateBooleanOutputFailureReason) {
	t.FailureReason = failureReason
}

func (t *TemplateBooleanOutput) SetErrorMessage(errorMessage *string) {
	t.ErrorMessage = errorMessage
}

func (t *TemplateBooleanOutput) SetContent(content *bool) {
	t.Content = content
}

func (t *TemplateBooleanOutput) Type() string {
	return t.type_
}

func (t *TemplateBooleanOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TemplateBooleanOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateBooleanOutput to string"
	}
	return string(jsonData)
}

func (t *TemplateBooleanOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "boolean" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "boolean", *literals.Type)
	}
	t.type_ = "boolean"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TemplateBooleanOutput) MarshalJSON() ([]byte, error) {
	type embed TemplateBooleanOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*t),
		Type:  "boolean",
	}
	return json.Marshal(marshaler)
}
