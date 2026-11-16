package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentResponseAttachment struct {
	Url             string  `json:"url" xml:"url" required:"true"`
	Name            string  `json:"name" xml:"name" required:"true"`
	MimeType        *string `json:"mime_type,omitempty" xml:"mime_type,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentResponseAttachment) GetUrl() (value string) {
	if a == nil {
		return
	}
	return a.Url
}

func (a *AgentResponseAttachment) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AgentResponseAttachment) GetMimeType() (value *string) {
	if a == nil {
		return
	}
	return a.MimeType
}

func (a *AgentResponseAttachment) SetUrl(url string) {
	a.Url = url
}

func (a *AgentResponseAttachment) SetName(name string) {
	a.Name = name
}

func (a *AgentResponseAttachment) SetMimeType(mimeType *string) {
	a.MimeType = mimeType
}

func (a *AgentResponseAttachment) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentResponseAttachment) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentResponseAttachment to string"
	}
	return string(jsonData)
}

func (a *AgentResponseAttachment) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentResponseAttachment) MarshalJSON() ([]byte, error) {
	type embed AgentResponseAttachment
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
