package types

import (
	"encoding/json"
	"fmt"
)

// The composition plan used to generate the song
type DetailedMusicResponseCompositionPlan struct {
	MusicPrompt     *MusicPrompt
	CompositionPlan *CompositionPlan
	typ             string
}

func (d *DetailedMusicResponseCompositionPlan) GetMusicPrompt() *MusicPrompt {
	if d == nil {
		return nil
	}
	if d.MusicPrompt == nil {
		return nil
	}
	return d.MusicPrompt
}

func (d *DetailedMusicResponseCompositionPlan) GetCompositionPlan() *CompositionPlan {
	if d == nil {
		return nil
	}
	if d.CompositionPlan == nil {
		return nil
	}
	return d.CompositionPlan
}

func (d *DetailedMusicResponseCompositionPlan) UnmarshalJSON(data []byte) error {
	var valueMusicPrompt *MusicPrompt
	if err := json.Unmarshal(data, &valueMusicPrompt); err == nil {
		d.typ = "MusicPrompt"
		d.MusicPrompt = valueMusicPrompt
		return nil
	}
	var valueCompositionPlan *CompositionPlan
	if err := json.Unmarshal(data, &valueCompositionPlan); err == nil {
		d.typ = "CompositionPlan"
		d.CompositionPlan = valueCompositionPlan
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, d)
}

func (d DetailedMusicResponseCompositionPlan) MarshalJSON() ([]byte, error) {
	if d.typ == "MusicPrompt" || d.MusicPrompt != nil {
		return json.Marshal(d.MusicPrompt)
	}
	if d.typ == "CompositionPlan" || d.CompositionPlan != nil {
		return json.Marshal(d.CompositionPlan)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", d)
}

type DetailedMusicResponseCompositionPlanVisitor interface {
	VisitMusicPrompt(*MusicPrompt) error
	VisitCompositionPlan(*CompositionPlan) error
}

func (d *DetailedMusicResponseCompositionPlan) Accept(visitor DetailedMusicResponseCompositionPlanVisitor) error {
	if d.typ == "MusicPrompt" || d.MusicPrompt != nil {
		return visitor.VisitMusicPrompt(d.MusicPrompt)
	}
	if d.typ == "CompositionPlan" || d.CompositionPlan != nil {
		return visitor.VisitCompositionPlan(d.CompositionPlan)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", d)
}
