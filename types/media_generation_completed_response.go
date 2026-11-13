package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A completed media generation and its output.
type MediaGenerationCompletedResponse struct {
	// The unique identifier of the generation.
	Id string `json:"id" xml:"id" required:"true"`
	// The lifecycle status of the generation.
	status string
	// A signed URL to download the generated media from. It expires about an hour after this response is returned; fetch the generation again for a fresh URL.
	ContentUrl string `json:"content_url" xml:"content_url" required:"true"`
	// The MIME type of the generated media.
	ContentMimeType string `json:"content_mime_type" xml:"content_mime_type" required:"true"`
	extraProperties map[string]interface{}
}

func (m *MediaGenerationCompletedResponse) GetId() (value string) {
	if m == nil {
		return
	}
	return m.Id
}

func (m *MediaGenerationCompletedResponse) GetContentUrl() (value string) {
	if m == nil {
		return
	}
	return m.ContentUrl
}

func (m *MediaGenerationCompletedResponse) GetContentMimeType() (value string) {
	if m == nil {
		return
	}
	return m.ContentMimeType
}

func (m *MediaGenerationCompletedResponse) SetId(id string) {
	m.Id = id
}

func (m *MediaGenerationCompletedResponse) SetContentUrl(contentUrl string) {
	m.ContentUrl = contentUrl
}

func (m *MediaGenerationCompletedResponse) SetContentMimeType(contentMimeType string) {
	m.ContentMimeType = contentMimeType
}

func (m *MediaGenerationCompletedResponse) Status() string {
	return m.status
}

func (m *MediaGenerationCompletedResponse) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MediaGenerationCompletedResponse) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MediaGenerationCompletedResponse to string"
	}
	return string(jsonData)
}

func (m *MediaGenerationCompletedResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias MediaGenerationCompletedResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = MediaGenerationCompletedResponse(tmp)
	var literals struct {
		Status *string `json:"status"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Status != nil && *literals.Status != "completed" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", m, "completed", *literals.Status)
	}
	m.status = "completed"
	extraProperties, err := internal.ExtractExtraProperties(data, *m, "status")
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MediaGenerationCompletedResponse) MarshalJSON() ([]byte, error) {
	type embed MediaGenerationCompletedResponse
	var marshaler = struct {
		embed
		Status string `json:"status"`
	}{
		embed:  embed(*m),
		Status: "completed",
	}
	return json.Marshal(marshaler)
}
