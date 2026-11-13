package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// An `object` output, delivered as `content` once completed: one output per
// field of the schema's `properties`, each an output in its own right with
// its own `status`. There are no dynamic fields.
type TemplateObjectOutput struct {
	// The id of the generation behind this output. Pass it as a `generation` reference to use the output as an input elsewhere.
	Id string `json:"id" xml:"id" required:"true"`
	// Where a run, or one of its outputs, is in its lifecycle. The run's is rolled
	// up from its outputs' (``aggregate_output_statuses``). A class rather than a
	// ``Literal`` so it publishes as one named enum instead of being inlined into
	// every kind.
	Status TemplateRunStatus `json:"status" xml:"status" required:"true"`
	// The category of failure. Present only when `status` is `failed`.
	FailureReason *TemplateObjectOutputFailureReason `json:"failure_reason,omitempty" xml:"failure_reason,omitempty"`
	// A human-readable description of the failure. Present only when `status` is `failed`. Failed generations are not charged.
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// The kind of output: the `type` of the port's `content_schema`.
	type_ string
	// One output per field, keyed by field name, each following the shape of that field's schema. Present only when `status` is `completed`. Reserved: no template produces this kind of output yet; it is published so that templates which do can be run with the same client.
	Content         map[string]*TemplateOutput `json:"content,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TemplateObjectOutput) GetId() (value string) {
	if t == nil {
		return
	}
	return t.Id
}

func (t *TemplateObjectOutput) GetStatus() (value TemplateRunStatus) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TemplateObjectOutput) GetFailureReason() (value *TemplateObjectOutputFailureReason) {
	if t == nil {
		return
	}
	return t.FailureReason
}

func (t *TemplateObjectOutput) GetErrorMessage() (value *string) {
	if t == nil {
		return
	}
	return t.ErrorMessage
}

func (t *TemplateObjectOutput) GetContent() (value map[string]*TemplateOutput) {
	if t == nil {
		return
	}
	return t.Content
}

func (t *TemplateObjectOutput) SetId(id string) {
	t.Id = id
}

func (t *TemplateObjectOutput) SetStatus(status TemplateRunStatus) {
	t.Status = status
}

func (t *TemplateObjectOutput) SetFailureReason(failureReason *TemplateObjectOutputFailureReason) {
	t.FailureReason = failureReason
}

func (t *TemplateObjectOutput) SetErrorMessage(errorMessage *string) {
	t.ErrorMessage = errorMessage
}

func (t *TemplateObjectOutput) SetContent(content map[string]*TemplateOutput) {
	t.Content = content
}

func (t *TemplateObjectOutput) Type() string {
	return t.type_
}

func (t *TemplateObjectOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TemplateObjectOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateObjectOutput to string"
	}
	return string(jsonData)
}

func (t *TemplateObjectOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "object" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "object", *literals.Type)
	}
	t.type_ = "object"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TemplateObjectOutput) MarshalJSON() ([]byte, error) {
	type embed TemplateObjectOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*t),
		Type:  "object",
	}
	return json.Marshal(marshaler)
}
