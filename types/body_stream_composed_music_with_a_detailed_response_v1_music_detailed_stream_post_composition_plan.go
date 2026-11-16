package types

import (
	"encoding/json"
	"fmt"
)

type BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPostCompositionPlan struct {
	MusicPrompt     *MusicPrompt
	CompositionPlan *CompositionPlan
	typ             string
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPostCompositionPlan) GetMusicPrompt() *MusicPrompt {
	if b == nil {
		return nil
	}
	if b.MusicPrompt == nil {
		return nil
	}
	return b.MusicPrompt
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPostCompositionPlan) GetCompositionPlan() *CompositionPlan {
	if b == nil {
		return nil
	}
	if b.CompositionPlan == nil {
		return nil
	}
	return b.CompositionPlan
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPostCompositionPlan) UnmarshalJSON(data []byte) error {
	var valueMusicPrompt *MusicPrompt
	if err := json.Unmarshal(data, &valueMusicPrompt); err == nil {
		b.typ = "MusicPrompt"
		b.MusicPrompt = valueMusicPrompt
		return nil
	}
	var valueCompositionPlan *CompositionPlan
	if err := json.Unmarshal(data, &valueCompositionPlan); err == nil {
		b.typ = "CompositionPlan"
		b.CompositionPlan = valueCompositionPlan
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, b)
}

func (b BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPostCompositionPlan) MarshalJSON() ([]byte, error) {
	if b.typ == "MusicPrompt" || b.MusicPrompt != nil {
		return json.Marshal(b.MusicPrompt)
	}
	if b.typ == "CompositionPlan" || b.CompositionPlan != nil {
		return json.Marshal(b.CompositionPlan)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", b)
}

type BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPostCompositionPlanVisitor interface {
	VisitMusicPrompt(*MusicPrompt) error
	VisitCompositionPlan(*CompositionPlan) error
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPostCompositionPlan) Accept(visitor BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPostCompositionPlanVisitor) error {
	if b.typ == "MusicPrompt" || b.MusicPrompt != nil {
		return visitor.VisitMusicPrompt(b.MusicPrompt)
	}
	if b.typ == "CompositionPlan" || b.CompositionPlan != nil {
		return visitor.VisitCompositionPlan(b.CompositionPlan)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", b)
}
