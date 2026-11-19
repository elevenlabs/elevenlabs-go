package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type BodyCreatePodcastV1StudioPodcastsPostSourceTwoItem struct {
	Type string
	Text *PodcastTextSource
	Url  *PodcastUrlSource
}

func (b *BodyCreatePodcastV1StudioPodcastsPostSourceTwoItem) GetText() *PodcastTextSource {
	if b == nil {
		return nil
	}
	if b.Text == nil {
		return nil
	}
	return b.Text
}

func (b *BodyCreatePodcastV1StudioPodcastsPostSourceTwoItem) GetUrl() *PodcastUrlSource {
	if b == nil {
		return nil
	}
	if b.Url == nil {
		return nil
	}
	return b.Url
}

func (b *BodyCreatePodcastV1StudioPodcastsPostSourceTwoItem) GetType() string {
	if b == nil {
		return ""
	}
	return b.Type
}

func (b *BodyCreatePodcastV1StudioPodcastsPostSourceTwoItem) UnmarshalJSON(data []byte) error {
	var unmarshaler struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(data, &unmarshaler); err != nil {
		return err
	}
	b.Type = unmarshaler.Type
	if unmarshaler.Type == "" {
		return fmt.Errorf("%T did not include discriminant type", b)
	}
	switch unmarshaler.Type {
	case "text":
		value := new(PodcastTextSource)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		b.Text = value
	case "url":
		value := new(PodcastUrlSource)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		b.Url = value
	}
	return nil
}

func (b BodyCreatePodcastV1StudioPodcastsPostSourceTwoItem) MarshalJSON() ([]byte, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}
	if b.Text != nil {
		return internal.MarshalJSONWithDiscriminant(b.Text, "type", "text")
	}
	if b.Url != nil {
		return internal.MarshalJSONWithDiscriminant(b.Url, "type", "url")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", b)
}

type BodyCreatePodcastV1StudioPodcastsPostSourceTwoItemVisitor interface {
	VisitText(*PodcastTextSource) error
	VisitUrl(*PodcastUrlSource) error
}

func (b *BodyCreatePodcastV1StudioPodcastsPostSourceTwoItem) Accept(visitor BodyCreatePodcastV1StudioPodcastsPostSourceTwoItemVisitor) error {
	if b.Text != nil {
		return visitor.VisitText(b.Text)
	}
	if b.Url != nil {
		return visitor.VisitUrl(b.Url)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", b)
}

func (b *BodyCreatePodcastV1StudioPodcastsPostSourceTwoItem) validate() error {
	if b == nil {
		return fmt.Errorf("type %T is nil", b)
	}
	var fields []string
	if b.Text != nil {
		fields = append(fields, "text")
	}
	if b.Url != nil {
		fields = append(fields, "url")
	}
	if len(fields) == 0 {
		if b.Type != "" {
			return fmt.Errorf("type %T defines a discriminant set to %q but the field is not set", b, b.Type)
		}
		return fmt.Errorf("type %T is empty", b)
	}
	if len(fields) > 1 {
		return fmt.Errorf("type %T defines values for %s, but only one value is allowed", b, fields)
	}
	if b.Type != "" {
		field := fields[0]
		if b.Type != field {
			return fmt.Errorf(
				"type %T defines a discriminant set to %q, but it does not match the %T field; either remove or update the discriminant to match",
				b,
				b.Type,
				b,
			)
		}
	}
	return nil
}
