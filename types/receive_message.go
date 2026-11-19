package types

import (
	"encoding/json"
	"fmt"
)

type ReceiveMessage struct {
	AudioOutput *AudioOutput
	FinalOutput *FinalOutput
	typ         string
}

func (r *ReceiveMessage) GetAudioOutput() *AudioOutput {
	if r == nil {
		return nil
	}
	if r.AudioOutput == nil {
		return nil
	}
	return r.AudioOutput
}

func (r *ReceiveMessage) GetFinalOutput() *FinalOutput {
	if r == nil {
		return nil
	}
	if r.FinalOutput == nil {
		return nil
	}
	return r.FinalOutput
}

func (r *ReceiveMessage) UnmarshalJSON(data []byte) error {
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
	var valueAudioOutput *AudioOutput
	if err := json.Unmarshal(data, &valueAudioOutput); err == nil {
		r.typ = "AudioOutput"
		r.AudioOutput = valueAudioOutput
		return nil
	}
	if hasAnyKey("is_final") {
		var valueFinalOutput *FinalOutput
		if err := json.Unmarshal(data, &valueFinalOutput); err == nil {
			r.typ = "FinalOutput"
			r.FinalOutput = valueFinalOutput
			return nil
		}
	}

	// Retried without their key guards, for a payload that matched none of them.
	var valueFinalOutputFallback *FinalOutput
	if err := json.Unmarshal(data, &valueFinalOutputFallback); err == nil {
		r.typ = "FinalOutput"
		r.FinalOutput = valueFinalOutputFallback
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, r)
}

func (r ReceiveMessage) MarshalJSON() ([]byte, error) {
	if r.typ == "AudioOutput" || r.AudioOutput != nil {
		return json.Marshal(r.AudioOutput)
	}
	if r.typ == "FinalOutput" || r.FinalOutput != nil {
		return json.Marshal(r.FinalOutput)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", r)
}

type ReceiveMessageVisitor interface {
	VisitAudioOutput(*AudioOutput) error
	VisitFinalOutput(*FinalOutput) error
}

func (r *ReceiveMessage) Accept(visitor ReceiveMessageVisitor) error {
	if r.typ == "AudioOutput" || r.AudioOutput != nil {
		return visitor.VisitAudioOutput(r.AudioOutput)
	}
	if r.typ == "FinalOutput" || r.FinalOutput != nil {
		return visitor.VisitFinalOutput(r.FinalOutput)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", r)
}
