package types

import (
	"encoding/json"
	"fmt"
)

// The composition plan extracted from the uploaded song. Only present if `extract_composition_plan` was provided in the request body.
type MusicUploadResponseCompositionPlan struct {
	MusicPrompt     *MusicPrompt
	CompositionPlan *CompositionPlan
	typ             string
}

func (m *MusicUploadResponseCompositionPlan) GetMusicPrompt() *MusicPrompt {
	if m == nil {
		return nil
	}
	if m.MusicPrompt == nil {
		return nil
	}
	return m.MusicPrompt
}

func (m *MusicUploadResponseCompositionPlan) GetCompositionPlan() *CompositionPlan {
	if m == nil {
		return nil
	}
	if m.CompositionPlan == nil {
		return nil
	}
	return m.CompositionPlan
}

func (m *MusicUploadResponseCompositionPlan) UnmarshalJSON(data []byte) error {
	var valueMusicPrompt *MusicPrompt
	if err := json.Unmarshal(data, &valueMusicPrompt); err == nil {
		m.typ = "MusicPrompt"
		m.MusicPrompt = valueMusicPrompt
		return nil
	}
	var valueCompositionPlan *CompositionPlan
	if err := json.Unmarshal(data, &valueCompositionPlan); err == nil {
		m.typ = "CompositionPlan"
		m.CompositionPlan = valueCompositionPlan
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, m)
}

func (m MusicUploadResponseCompositionPlan) MarshalJSON() ([]byte, error) {
	if m.typ == "MusicPrompt" || m.MusicPrompt != nil {
		return json.Marshal(m.MusicPrompt)
	}
	if m.typ == "CompositionPlan" || m.CompositionPlan != nil {
		return json.Marshal(m.CompositionPlan)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", m)
}

type MusicUploadResponseCompositionPlanVisitor interface {
	VisitMusicPrompt(*MusicPrompt) error
	VisitCompositionPlan(*CompositionPlan) error
}

func (m *MusicUploadResponseCompositionPlan) Accept(visitor MusicUploadResponseCompositionPlanVisitor) error {
	if m.typ == "MusicPrompt" || m.MusicPrompt != nil {
		return visitor.VisitMusicPrompt(m.MusicPrompt)
	}
	if m.typ == "CompositionPlan" || m.CompositionPlan != nil {
		return visitor.VisitCompositionPlan(m.CompositionPlan)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", m)
}
