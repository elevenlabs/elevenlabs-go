package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A newly queued media generation; fetch the GET endpoint for the output.
type MediaGenerationCreateResponse struct {
	// The unique identifier of the generation. Pass it to the corresponding GET endpoint to retrieve the output.
	Id string `json:"id" xml:"id" required:"true"`
	// A newly created generation is always `pending`.
	status          string
	extraProperties map[string]interface{}
}

func (m *MediaGenerationCreateResponse) GetId() (value string) {
	if m == nil {
		return
	}
	return m.Id
}

func (m *MediaGenerationCreateResponse) SetId(id string) {
	m.Id = id
}

func (m *MediaGenerationCreateResponse) Status() string {
	return m.status
}

func (m *MediaGenerationCreateResponse) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MediaGenerationCreateResponse) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MediaGenerationCreateResponse to string"
	}
	return string(jsonData)
}

func (m *MediaGenerationCreateResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias MediaGenerationCreateResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = MediaGenerationCreateResponse(tmp)
	var literals struct {
		Status *string `json:"status"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Status != nil && *literals.Status != "pending" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", m, "pending", *literals.Status)
	}
	m.status = "pending"
	extraProperties, err := internal.ExtractExtraProperties(data, *m, "status")
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MediaGenerationCreateResponse) MarshalJSON() ([]byte, error) {
	type embed MediaGenerationCreateResponse
	var marshaler = struct {
		embed
		Status string `json:"status"`
	}{
		embed:  embed(*m),
		Status: "pending",
	}
	return json.Marshal(marshaler)
}
