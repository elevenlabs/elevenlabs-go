package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A media generation that has not finished yet.
type MediaGenerationInProgressResponse struct {
	// The unique identifier of the generation.
	Id string `json:"id" xml:"id" required:"true"`
	// The lifecycle status of the generation. It ends at `completed` or `failed`.
	Status          MediaGenerationInProgressResponseStatus `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (m *MediaGenerationInProgressResponse) GetId() (value string) {
	if m == nil {
		return
	}
	return m.Id
}

func (m *MediaGenerationInProgressResponse) GetStatus() (value MediaGenerationInProgressResponseStatus) {
	if m == nil {
		return
	}
	return m.Status
}

func (m *MediaGenerationInProgressResponse) SetId(id string) {
	m.Id = id
}

func (m *MediaGenerationInProgressResponse) SetStatus(status MediaGenerationInProgressResponseStatus) {
	m.Status = status
}

func (m *MediaGenerationInProgressResponse) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MediaGenerationInProgressResponse) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MediaGenerationInProgressResponse to string"
	}
	return string(jsonData)
}

func (m *MediaGenerationInProgressResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias MediaGenerationInProgressResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = MediaGenerationInProgressResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MediaGenerationInProgressResponse) MarshalJSON() ([]byte, error) {
	type embed MediaGenerationInProgressResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
