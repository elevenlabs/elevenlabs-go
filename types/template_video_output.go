package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A `video` output, downloaded from `content_url` once completed.
type TemplateVideoOutput struct {
	// The id of the generation behind this output. Pass it as a `generation` reference to use the output as an input elsewhere.
	Id string `json:"id" xml:"id" required:"true"`
	// Where a run, or one of its outputs, is in its lifecycle. The run's is rolled
	// up from its outputs' (``aggregate_output_statuses``). A class rather than a
	// ``Literal`` so it publishes as one named enum instead of being inlined into
	// every kind.
	Status TemplateRunStatus `json:"status" xml:"status" required:"true"`
	// The category of failure. Present only when `status` is `failed`.
	FailureReason *TemplateVideoOutputFailureReason `json:"failure_reason,omitempty" xml:"failure_reason,omitempty"`
	// A human-readable description of the failure. Present only when `status` is `failed`. Failed generations are not charged.
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// A signed URL to download the generated media from. Present only when `status` is `completed`. It expires about an hour after this response is returned; fetch the run again for a fresh URL.
	ContentUrl *string `json:"content_url,omitempty" xml:"content_url,omitempty"`
	// The MIME type of the generated media. Present only when `status` is `completed`.
	ContentMimeType *string `json:"content_mime_type,omitempty" xml:"content_mime_type,omitempty"`
	// The kind of output: the `type` of the port's `content_schema`.
	type_           string
	extraProperties map[string]interface{}
}

func (t *TemplateVideoOutput) GetId() (value string) {
	if t == nil {
		return
	}
	return t.Id
}

func (t *TemplateVideoOutput) GetStatus() (value TemplateRunStatus) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TemplateVideoOutput) GetFailureReason() (value *TemplateVideoOutputFailureReason) {
	if t == nil {
		return
	}
	return t.FailureReason
}

func (t *TemplateVideoOutput) GetErrorMessage() (value *string) {
	if t == nil {
		return
	}
	return t.ErrorMessage
}

func (t *TemplateVideoOutput) GetContentUrl() (value *string) {
	if t == nil {
		return
	}
	return t.ContentUrl
}

func (t *TemplateVideoOutput) GetContentMimeType() (value *string) {
	if t == nil {
		return
	}
	return t.ContentMimeType
}

func (t *TemplateVideoOutput) SetId(id string) {
	t.Id = id
}

func (t *TemplateVideoOutput) SetStatus(status TemplateRunStatus) {
	t.Status = status
}

func (t *TemplateVideoOutput) SetFailureReason(failureReason *TemplateVideoOutputFailureReason) {
	t.FailureReason = failureReason
}

func (t *TemplateVideoOutput) SetErrorMessage(errorMessage *string) {
	t.ErrorMessage = errorMessage
}

func (t *TemplateVideoOutput) SetContentUrl(contentUrl *string) {
	t.ContentUrl = contentUrl
}

func (t *TemplateVideoOutput) SetContentMimeType(contentMimeType *string) {
	t.ContentMimeType = contentMimeType
}

func (t *TemplateVideoOutput) Type() string {
	return t.type_
}

func (t *TemplateVideoOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TemplateVideoOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateVideoOutput to string"
	}
	return string(jsonData)
}

func (t *TemplateVideoOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "video" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "video", *literals.Type)
	}
	t.type_ = "video"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TemplateVideoOutput) MarshalJSON() ([]byte, error) {
	type embed TemplateVideoOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*t),
		Type:  "video",
	}
	return json.Marshal(marshaler)
}
