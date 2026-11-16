package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// The type of podcast to generate. Can be 'conversation', an interaction between two voices, or 'bulletin', a monologue.
type BodyCreatePodcastV1StudioPodcastsPostMode struct {
	Type         string
	Bulletin     *PodcastBulletinMode
	Conversation *PodcastConversationMode
}

func (b *BodyCreatePodcastV1StudioPodcastsPostMode) GetBulletin() *PodcastBulletinMode {
	if b == nil {
		return nil
	}
	if b.Bulletin == nil {
		return nil
	}
	return b.Bulletin
}

func (b *BodyCreatePodcastV1StudioPodcastsPostMode) GetConversation() *PodcastConversationMode {
	if b == nil {
		return nil
	}
	if b.Conversation == nil {
		return nil
	}
	return b.Conversation
}

func (b *BodyCreatePodcastV1StudioPodcastsPostMode) GetType() string {
	if b == nil {
		return ""
	}
	return b.Type
}

func (b *BodyCreatePodcastV1StudioPodcastsPostMode) UnmarshalJSON(data []byte) error {
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
	case "bulletin":
		value := new(PodcastBulletinMode)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		b.Bulletin = value
	case "conversation":
		value := new(PodcastConversationMode)
		if err := json.Unmarshal(data, &value); err != nil {
			return err
		}
		b.Conversation = value
	}
	return nil
}

func (b BodyCreatePodcastV1StudioPodcastsPostMode) MarshalJSON() ([]byte, error) {
	if err := b.validate(); err != nil {
		return nil, err
	}
	if b.Bulletin != nil {
		return internal.MarshalJSONWithDiscriminant(b.Bulletin, "type", "bulletin")
	}
	if b.Conversation != nil {
		return internal.MarshalJSONWithDiscriminant(b.Conversation, "type", "conversation")
	}
	return nil, fmt.Errorf("type %T does not define a non-empty union type", b)
}

type BodyCreatePodcastV1StudioPodcastsPostModeVisitor interface {
	VisitBulletin(*PodcastBulletinMode) error
	VisitConversation(*PodcastConversationMode) error
}

func (b *BodyCreatePodcastV1StudioPodcastsPostMode) Accept(visitor BodyCreatePodcastV1StudioPodcastsPostModeVisitor) error {
	if b.Bulletin != nil {
		return visitor.VisitBulletin(b.Bulletin)
	}
	if b.Conversation != nil {
		return visitor.VisitConversation(b.Conversation)
	}
	return fmt.Errorf("type %T does not define a non-empty union type", b)
}

func (b *BodyCreatePodcastV1StudioPodcastsPostMode) validate() error {
	if b == nil {
		return fmt.Errorf("type %T is nil", b)
	}
	var fields []string
	if b.Bulletin != nil {
		fields = append(fields, "bulletin")
	}
	if b.Conversation != nil {
		fields = append(fields, "conversation")
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
