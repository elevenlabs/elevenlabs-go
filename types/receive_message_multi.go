package types

import (
	"encoding/json"
	"fmt"
)

type ReceiveMessageMulti struct {
	AudioOutputMulti *AudioOutputMulti
	FinalOutputMulti *FinalOutputMulti
	typ              string
}

func (r *ReceiveMessageMulti) GetAudioOutputMulti() *AudioOutputMulti {
	if r == nil {
		return nil
	}
	if r.AudioOutputMulti == nil {
		return nil
	}
	return r.AudioOutputMulti
}

func (r *ReceiveMessageMulti) GetFinalOutputMulti() *FinalOutputMulti {
	if r == nil {
		return nil
	}
	if r.FinalOutputMulti == nil {
		return nil
	}
	return r.FinalOutputMulti
}

func (r *ReceiveMessageMulti) UnmarshalJSON(data []byte) error {
	var unionKeys map[string]json.RawMessage
	_ = json.Unmarshal(data, &unionKeys)
	hasAnyKey := func(keys ...string) bool {
		for _, key := range keys {
			if _, ok := unionKeys[key]; ok {
				return true
			}
		}
		return false
	}
	var valueAudioOutputMulti *AudioOutputMulti
	if err := json.Unmarshal(data, &valueAudioOutputMulti); err == nil {
		r.typ = "AudioOutputMulti"
		r.AudioOutputMulti = valueAudioOutputMulti
		return nil
	}
	if hasAnyKey("is_final", "context_id") {
		var valueFinalOutputMulti *FinalOutputMulti
		if err := json.Unmarshal(data, &valueFinalOutputMulti); err == nil {
			r.typ = "FinalOutputMulti"
			r.FinalOutputMulti = valueFinalOutputMulti
			return nil
		}
	}

	// Retried without their key guards, for a payload that matched none of them.
	var valueFinalOutputMultiFallback *FinalOutputMulti
	if err := json.Unmarshal(data, &valueFinalOutputMultiFallback); err == nil {
		r.typ = "FinalOutputMulti"
		r.FinalOutputMulti = valueFinalOutputMultiFallback
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, r)
}

func (r ReceiveMessageMulti) MarshalJSON() ([]byte, error) {
	if r.typ == "AudioOutputMulti" || r.AudioOutputMulti != nil {
		return json.Marshal(r.AudioOutputMulti)
	}
	if r.typ == "FinalOutputMulti" || r.FinalOutputMulti != nil {
		return json.Marshal(r.FinalOutputMulti)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", r)
}

type ReceiveMessageMultiVisitor interface {
	VisitAudioOutputMulti(*AudioOutputMulti) error
	VisitFinalOutputMulti(*FinalOutputMulti) error
}

func (r *ReceiveMessageMulti) Accept(visitor ReceiveMessageMultiVisitor) error {
	if r.typ == "AudioOutputMulti" || r.AudioOutputMulti != nil {
		return visitor.VisitAudioOutputMulti(r.AudioOutputMulti)
	}
	if r.typ == "FinalOutputMulti" || r.FinalOutputMulti != nil {
		return visitor.VisitFinalOutputMulti(r.FinalOutputMulti)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", r)
}
