package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A `string` output, delivered as `content` once completed.
type TemplateStringOutput struct {
	// The id of the generation behind this output. Pass it as a `generation` reference to use the output as an input elsewhere.
	Id string `json:"id" xml:"id" required:"true"`
	// Where a run, or one of its outputs, is in its lifecycle. The run's is rolled
	// up from its outputs' (``aggregate_output_statuses``). A class rather than a
	// ``Literal`` so it publishes as one named enum instead of being inlined into
	// every kind.
	Status TemplateRunStatus `json:"status" xml:"status" required:"true"`
	// The category of failure. Present only when `status` is `failed`.
	FailureReason *TemplateStringOutputFailureReason `json:"failure_reason,omitempty" xml:"failure_reason,omitempty"`
	// A human-readable description of the failure. Present only when `status` is `failed`. Failed generations are not charged.
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// The kind of output: the `type` of the port's `content_schema`.
	type_ string
	// The generated text. Present only when `status` is `completed`.
	Content         *string `json:"content,omitempty" xml:"content,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TemplateStringOutput) GetId() (value string) {
	if t == nil {
		return
	}
	return t.Id
}

func (t *TemplateStringOutput) GetStatus() (value TemplateRunStatus) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TemplateStringOutput) GetFailureReason() (value *TemplateStringOutputFailureReason) {
	if t == nil {
		return
	}
	return t.FailureReason
}

func (t *TemplateStringOutput) GetErrorMessage() (value *string) {
	if t == nil {
		return
	}
	return t.ErrorMessage
}

func (t *TemplateStringOutput) GetContent() (value *string) {
	if t == nil {
		return
	}
	return t.Content
}

func (t *TemplateStringOutput) SetId(id string) {
	t.Id = id
}

func (t *TemplateStringOutput) SetStatus(status TemplateRunStatus) {
	t.Status = status
}

func (t *TemplateStringOutput) SetFailureReason(failureReason *TemplateStringOutputFailureReason) {
	t.FailureReason = failureReason
}

func (t *TemplateStringOutput) SetErrorMessage(errorMessage *string) {
	t.ErrorMessage = errorMessage
}

func (t *TemplateStringOutput) SetContent(content *string) {
	t.Content = content
}

func (t *TemplateStringOutput) Type() string {
	return t.type_
}

func (t *TemplateStringOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TemplateStringOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateStringOutput to string"
	}
	return string(jsonData)
}

func (t *TemplateStringOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "string" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "string", *literals.Type)
	}
	t.type_ = "string"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TemplateStringOutput) MarshalJSON() ([]byte, error) {
	type embed TemplateStringOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*t),
		Type:  "string",
	}
	return json.Marshal(marshaler)
}
