package types

import (
	"encoding/json"
	"fmt"
)

// The source content for the Podcast.
type BodyCreatePodcastV1StudioPodcastsPostSource struct {
	PodcastTextSource                                      *PodcastTextSource
	PodcastUrlSource                                       *PodcastUrlSource
	BodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList []*BodyCreatePodcastV1StudioPodcastsPostSourceTwoItem
	typ                                                    string
}

func (b *BodyCreatePodcastV1StudioPodcastsPostSource) GetPodcastTextSource() *PodcastTextSource {
	if b == nil {
		return nil
	}
	if b.PodcastTextSource == nil {
		return nil
	}
	return b.PodcastTextSource
}

func (b *BodyCreatePodcastV1StudioPodcastsPostSource) GetPodcastUrlSource() *PodcastUrlSource {
	if b == nil {
		return nil
	}
	if b.PodcastUrlSource == nil {
		return nil
	}
	return b.PodcastUrlSource
}

func (b *BodyCreatePodcastV1StudioPodcastsPostSource) GetBodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList() []*BodyCreatePodcastV1StudioPodcastsPostSourceTwoItem {
	if b == nil {
		return nil
	}
	return b.BodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList
}

func (b *BodyCreatePodcastV1StudioPodcastsPostSource) UnmarshalJSON(data []byte) error {
	var unionKeys map[string]json.RawMessage
	_ = json.Unmarshal(data, &unionKeys)
	hasLiteral := func(key string, want interface{}) bool {
		raw, ok := unionKeys[key]
		if !ok {
			return false
		}
		var got interface{}
		return json.Unmarshal(raw, &got) == nil && got == want
	}
	matchedLiteral := false
	if hasLiteral("type", "text") {
		matchedLiteral = true
		var valuePodcastTextSourceLiteral *PodcastTextSource
		if err := json.Unmarshal(data, &valuePodcastTextSourceLiteral); err == nil {
			b.typ = "PodcastTextSource"
			b.PodcastTextSource = valuePodcastTextSourceLiteral
			return nil
		}
	}
	if hasLiteral("type", "url") {
		matchedLiteral = true
		var valuePodcastUrlSourceLiteral *PodcastUrlSource
		if err := json.Unmarshal(data, &valuePodcastUrlSourceLiteral); err == nil {
			b.typ = "PodcastUrlSource"
			b.PodcastUrlSource = valuePodcastUrlSourceLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
	}
	var valuePodcastTextSource *PodcastTextSource
	if err := json.Unmarshal(data, &valuePodcastTextSource); err == nil {
		b.typ = "PodcastTextSource"
		b.PodcastTextSource = valuePodcastTextSource
		return nil
	}
	var valuePodcastUrlSource *PodcastUrlSource
	if err := json.Unmarshal(data, &valuePodcastUrlSource); err == nil {
		b.typ = "PodcastUrlSource"
		b.PodcastUrlSource = valuePodcastUrlSource
		return nil
	}
	var valueBodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList []*BodyCreatePodcastV1StudioPodcastsPostSourceTwoItem
	if err := json.Unmarshal(data, &valueBodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList); err == nil {
		b.typ = "BodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList"
		b.BodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList = valueBodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BodyCreatePodcastV1StudioPodcastsPostSource) MarshalJSON() ([]byte, error) {
	if b.typ == "PodcastTextSource" || b.PodcastTextSource != nil {
		return json.Marshal(b.PodcastTextSource)
	}
	if b.typ == "PodcastUrlSource" || b.PodcastUrlSource != nil {
		return json.Marshal(b.PodcastUrlSource)
	}
	if b.typ == "BodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList" || b.BodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList != nil {
		return json.Marshal(b.BodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyCreatePodcastV1StudioPodcastsPostSourceVisitor interface {
	VisitPodcastTextSource(*PodcastTextSource) error
	VisitPodcastUrlSource(*PodcastUrlSource) error
	VisitBodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList([]*BodyCreatePodcastV1StudioPodcastsPostSourceTwoItem) error
}

func (b *BodyCreatePodcastV1StudioPodcastsPostSource) Accept(visitor BodyCreatePodcastV1StudioPodcastsPostSourceVisitor) error {
	if b.typ == "PodcastTextSource" || b.PodcastTextSource != nil {
		return visitor.VisitPodcastTextSource(b.PodcastTextSource)
	}
	if b.typ == "PodcastUrlSource" || b.PodcastUrlSource != nil {
		return visitor.VisitPodcastUrlSource(b.PodcastUrlSource)
	}
	if b.typ == "BodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList" || b.BodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList != nil {
		return visitor.VisitBodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList(b.BodyCreatePodcastV1StudioPodcastsPostSourceTwoItemList)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}
