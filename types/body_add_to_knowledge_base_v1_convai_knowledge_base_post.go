package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodyAddToKnowledgeBaseV1ConvaiKnowledgeBasePost struct {
	// A custom, human-readable name for the document.
	Name *string `json:"name,omitempty" xml:"name,omitempty" minLength:"1"`
	// URL to a page of documentation that the agent will have access to in order to interact with users.
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
	// Documentation that the agent will have access to in order to interact with users.
	File            io.Reader `json:"file,omitempty" xml:"file,omitempty"`
	extraProperties map[string]interface{}
}

func (b *BodyAddToKnowledgeBaseV1ConvaiKnowledgeBasePost) GetName() (value *string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyAddToKnowledgeBaseV1ConvaiKnowledgeBasePost) GetUrl() (value *string) {
	if b == nil {
		return
	}
	return b.Url
}

func (b *BodyAddToKnowledgeBaseV1ConvaiKnowledgeBasePost) GetFile() (value io.Reader) {
	if b == nil {
		return
	}
	return b.File
}

func (b *BodyAddToKnowledgeBaseV1ConvaiKnowledgeBasePost) SetName(name *string) {
	b.Name = name
}

func (b *BodyAddToKnowledgeBaseV1ConvaiKnowledgeBasePost) SetUrl(url *string) {
	b.Url = url
}

func (b *BodyAddToKnowledgeBaseV1ConvaiKnowledgeBasePost) SetFile(file io.Reader) {
	b.File = file
}

func (b *BodyAddToKnowledgeBaseV1ConvaiKnowledgeBasePost) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BodyAddToKnowledgeBaseV1ConvaiKnowledgeBasePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyAddToKnowledgeBaseV1ConvaiKnowledgeBasePost to string"
	}
	return string(jsonData)
}

func (b *BodyAddToKnowledgeBaseV1ConvaiKnowledgeBasePost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BodyAddToKnowledgeBaseV1ConvaiKnowledgeBasePost) MarshalJSON() ([]byte, error) {
	type embed BodyAddToKnowledgeBaseV1ConvaiKnowledgeBasePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
