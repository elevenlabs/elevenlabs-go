package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// An `image` output, downloaded from `content_url` once completed.
type TemplateImageOutput struct {
	// The id of the generation behind this output. Pass it as a `generation` reference to use the output as an input elsewhere.
	Id string `json:"id" xml:"id" required:"true"`
	// Where a run, or one of its outputs, is in its lifecycle. The run's is rolled
	// up from its outputs' (``aggregate_output_statuses``). A class rather than a
	// ``Literal`` so it publishes as one named enum instead of being inlined into
	// every kind.
	Status TemplateRunStatus `json:"status" xml:"status" required:"true"`
	// The category of failure. Present only when `status` is `failed`.
	FailureReason *TemplateImageOutputFailureReason `json:"failure_reason,omitempty" xml:"failure_reason,omitempty"`
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

func (t *TemplateImageOutput) GetId() (value string) {
	if t == nil {
		return
	}
	return t.Id
}

func (t *TemplateImageOutput) GetStatus() (value TemplateRunStatus) {
	if t == nil {
		return
	}
	return t.Status
}

func (t *TemplateImageOutput) GetFailureReason() (value *TemplateImageOutputFailureReason) {
	if t == nil {
		return
	}
	return t.FailureReason
}

func (t *TemplateImageOutput) GetErrorMessage() (value *string) {
	if t == nil {
		return
	}
	return t.ErrorMessage
}

func (t *TemplateImageOutput) GetContentUrl() (value *string) {
	if t == nil {
		return
	}
	return t.ContentUrl
}

func (t *TemplateImageOutput) GetContentMimeType() (value *string) {
	if t == nil {
		return
	}
	return t.ContentMimeType
}

func (t *TemplateImageOutput) SetId(id string) {
	t.Id = id
}

func (t *TemplateImageOutput) SetStatus(status TemplateRunStatus) {
	t.Status = status
}

func (t *TemplateImageOutput) SetFailureReason(failureReason *TemplateImageOutputFailureReason) {
	t.FailureReason = failureReason
}

func (t *TemplateImageOutput) SetErrorMessage(errorMessage *string) {
	t.ErrorMessage = errorMessage
}

func (t *TemplateImageOutput) SetContentUrl(contentUrl *string) {
	t.ContentUrl = contentUrl
}

func (t *TemplateImageOutput) SetContentMimeType(contentMimeType *string) {
	t.ContentMimeType = contentMimeType
}

func (t *TemplateImageOutput) Type() string {
	return t.type_
}

func (t *TemplateImageOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TemplateImageOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TemplateImageOutput to string"
	}
	return string(jsonData)
}

func (t *TemplateImageOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "image" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "image", *literals.Type)
	}
	t.type_ = "image"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TemplateImageOutput) MarshalJSON() ([]byte, error) {
	type embed TemplateImageOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*t),
		Type:  "image",
	}
	return json.Marshal(marshaler)
}
