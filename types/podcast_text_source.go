package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PodcastTextSource struct {
	// The type of source to create.
	type_ string
	// The text to create the podcast from.
	Text            string `json:"text" xml:"text" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PodcastTextSource) GetText() (value string) {
	if p == nil {
		return
	}
	return p.Text
}

func (p *PodcastTextSource) SetText(text string) {
	p.Text = text
}

func (p *PodcastTextSource) Type() string {
	return p.type_
}

func (p *PodcastTextSource) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PodcastTextSource) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PodcastTextSource to string"
	}
	return string(jsonData)
}

func (p *PodcastTextSource) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PodcastTextSource
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PodcastTextSource(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "text" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "text", *literals.Type)
	}
	p.type_ = "text"
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PodcastTextSource) MarshalJSON() ([]byte, error) {
	type embed PodcastTextSource
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "text",
	}
	return json.Marshal(marshaler)
}
