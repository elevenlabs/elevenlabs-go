package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PodcastUrlSource struct {
	// The type of source to create.
	type_ string
	// The URL to create the podcast from.
	Url             string `json:"url" xml:"url" required:"true"`
	extraProperties map[string]interface{}
}

func (p *PodcastUrlSource) GetUrl() (value string) {
	if p == nil {
		return
	}
	return p.Url
}

func (p *PodcastUrlSource) SetUrl(url string) {
	p.Url = url
}

func (p *PodcastUrlSource) Type() string {
	return p.type_
}

func (p *PodcastUrlSource) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PodcastUrlSource) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PodcastUrlSource to string"
	}
	return string(jsonData)
}

func (p *PodcastUrlSource) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, p); err != nil {
		return err
	}
	type alias PodcastUrlSource
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*p = PodcastUrlSource(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "url" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", p, "url", *literals.Type)
	}
	p.type_ = "url"
	extraProperties, err := internal.ExtractExtraProperties(data, *p, "type")
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PodcastUrlSource) MarshalJSON() ([]byte, error) {
	type embed PodcastUrlSource
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*p),
		Type:  "url",
	}
	return json.Marshal(marshaler)
}
