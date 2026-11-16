package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// An `array` output, delivered as `content` once completed: one output per
// element, each an output in its own right with its own `status`. A list can
// be arbitrarily long, so `content` is its first page; when `has_more` is
// true, `next_cursor` leads to the rest.
type TemplateArrayOutput struct {
	// The id of the generation behind this output. Pass it as a `generation` reference to use the output as an input elsewhere.
	Id string `json:"id" xml:"id" required:"true"`
	// Where a run, or one of its outputs, is in its lifecycle. The run's is rolled
	// up from its outputs' (``aggregate_output_statuses``). A class rather than a
	// ``Literal`` so it publishes as one named enum instead of being inlined into
	// every kind.
	Status TemplateRunStatus `json:"status" xml:"status" required:"true"`
	// The category of failure. Present only when `status` is `failed`.
	FailureReason *TemplateArrayOutputFailureReason `json:"failure_reason,omitempty" xml:"failure_reason,omitempty"`
	// A human-readable description of the failure. Present only when `status` is `failed`. Failed generations are not charged.
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// The kind of output: the `type` of the port's `content_schema`.
	type_ string
	// The first elements of the list, in order, each following the shape of the schema's `items`. Present only when `status` is `completed`. When `has_more` is true this is not the whole list. Reserved: no template produces this kind of output yet; it is published so that templates which do can be run with the same client.
	Content []*TemplateOutput `json:"content,omitempty" xml:"content,omitempty"`
	// Pass as `cursor` to a later endpoint to fetch the elements after `content`. `null` when `content` holds the whole list, or before the output completes. Reserved: no template produces this kind of output yet; it is published so that templates which do can be run with the same client.
	NextCursor *string `json:"next_cursor,omitempty" xml:"next_cursor,omitempty"`
	// Whether the list has elements beyond `content`.
	HasMore         *bool `json:"has_more,omitempty" xml:"has_more,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TemplateArrayOutput) GetId() (value string) {
	if t == nil {
		return
	}
	return t.Id
}

func (t *TemplateArrayOutput) GetStatus() (value TemplateRunStatus) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TemplateArrayOutput) GetFailureReason() (value *TemplateArrayOutputFailureReason) {
	if t == nil {
		return
	}
	return t.FailureReason
}

func (t *TemplateArrayOutput) GetErrorMessage() (value *string) {
	if t == nil {
		return
	}
	return t.ErrorMessage
}

func (t *TemplateArrayOutput) GetContent() (value []*TemplateOutput) {
	if t == nil {
		return
	}
	return t.Content
}

func (t *TemplateArrayOutput) GetNextCursor() (value *string) {
	if t == nil {
		return
	}
	return t.NextCursor
}

func (t *TemplateArrayOutput) GetHasMore() (value *bool) {
	if t == nil {
		return
	}
	return t.HasMore
}

func (t *TemplateArrayOutput) SetId(id string) {
	t.Id = id
}

func (t *TemplateArrayOutput) SetStatus(status TemplateRunStatus) {
	t.Status = status
}

func (t *TemplateArrayOutput) SetFailureReason(failureReason *TemplateArrayOutputFailureReason) {
	t.FailureReason = failureReason
}

func (t *TemplateArrayOutput) SetErrorMessage(errorMessage *string) {
	t.ErrorMessage = errorMessage
}

func (t *TemplateArrayOutput) SetContent(content []*TemplateOutput) {
	t.Content = content
}

func (t *TemplateArrayOutput) SetNextCursor(nextCursor *string) {
	t.NextCursor = nextCursor
}

func (t *TemplateArrayOutput) SetHasMore(hasMore *bool) {
	t.HasMore = hasMore
}

func (t *TemplateArrayOutput) Type() string {
	return t.type_
}

func (t *TemplateArrayOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TemplateArrayOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateArrayOutput to string"
	}
	return string(jsonData)
}

func (t *TemplateArrayOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "array" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "array", *literals.Type)
	}
	t.type_ = "array"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TemplateArrayOutput) MarshalJSON() ([]byte, error) {
	type embed TemplateArrayOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*t),
		Type:  "array",
	}
	return json.Marshal(marshaler)
}
