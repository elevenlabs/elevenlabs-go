package types

import (
	"encoding/json"
	"fmt"
)

type ReceiveTextToDialogueWebsocketMessageMulti struct {
	AudioChunkMulti        *AudioChunkMulti
	FinalAudioForTurnMulti *FinalAudioForTurnMulti
	FinalMulti             *FinalMulti
	Error                  *Error
	typ                    string
}

func (r *ReceiveTextToDialogueWebsocketMessageMulti) GetAudioChunkMulti() *AudioChunkMulti {
	if r == nil {
		return nil
	}
	if r.AudioChunkMulti == nil {
		return nil
	}
	return r.AudioChunkMulti
}

func (r *ReceiveTextToDialogueWebsocketMessageMulti) GetFinalAudioForTurnMulti() *FinalAudioForTurnMulti {
	if r == nil {
		return nil
	}
	if r.FinalAudioForTurnMulti == nil {
		return nil
	}
	return r.FinalAudioForTurnMulti
}

func (r *ReceiveTextToDialogueWebsocketMessageMulti) GetFinalMulti() *FinalMulti {
	if r == nil {
		return nil
	}
	if r.FinalMulti == nil {
		return nil
	}
	return r.FinalMulti
}

func (r *ReceiveTextToDialogueWebsocketMessageMulti) GetError() *Error {
	if r == nil {
		return nil
	}
	if r.Error == nil {
		return nil
	}
	return r.Error
}

func (r *ReceiveTextToDialogueWebsocketMessageMulti) UnmarshalJSON(data []byte) error {
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
	if hasAnyKey("audio", "alignment", "normalized_alignment", "context_id") {
		var valueAudioChunkMulti *AudioChunkMulti
		if err := json.Unmarshal(data, &valueAudioChunkMulti); err == nil {
			r.typ = "AudioChunkMulti"
			r.AudioChunkMulti = valueAudioChunkMulti
			return nil
		}
	}
	if hasAnyKey("is_final_audio_for_turn", "context_id") {
		var valueFinalAudioForTurnMulti *FinalAudioForTurnMulti
		if err := json.Unmarshal(data, &valueFinalAudioForTurnMulti); err == nil {
			r.typ = "FinalAudioForTurnMulti"
			r.FinalAudioForTurnMulti = valueFinalAudioForTurnMulti
			return nil
		}
	}
	var valueFinalMulti *FinalMulti
	if err := json.Unmarshal(data, &valueFinalMulti); err == nil {
		r.typ = "FinalMulti"
		r.FinalMulti = valueFinalMulti
		return nil
	}
	var valueError *Error
	if err := json.Unmarshal(data, &valueError); err == nil {
		r.typ = "Error"
		r.Error = valueError
		return nil
	}

	// Retried without their key guards, for a payload that matched none of them.
	var valueAudioChunkMultiFallback *AudioChunkMulti
	if err := json.Unmarshal(data, &valueAudioChunkMultiFallback); err == nil {
		r.typ = "AudioChunkMulti"
		r.AudioChunkMulti = valueAudioChunkMultiFallback
		return nil
	}
	var valueFinalAudioForTurnMultiFallback *FinalAudioForTurnMulti
	if err := json.Unmarshal(data, &valueFinalAudioForTurnMultiFallback); err == nil {
		r.typ = "FinalAudioForTurnMulti"
		r.FinalAudioForTurnMulti = valueFinalAudioForTurnMultiFallback
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, r)
}

func (r ReceiveTextToDialogueWebsocketMessageMulti) MarshalJSON() ([]byte, error) {
	if r.typ == "AudioChunkMulti" || r.AudioChunkMulti != nil {
		return json.Marshal(r.AudioChunkMulti)
	}
	if r.typ == "FinalAudioForTurnMulti" || r.FinalAudioForTurnMulti != nil {
		return json.Marshal(r.FinalAudioForTurnMulti)
	}
	if r.typ == "FinalMulti" || r.FinalMulti != nil {
		return json.Marshal(r.FinalMulti)
	}
	if r.typ == "Error" || r.Error != nil {
		return json.Marshal(r.Error)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", r)
}

type ReceiveTextToDialogueWebsocketMessageMultiVisitor interface {
	VisitAudioChunkMulti(*AudioChunkMulti) error
	VisitFinalAudioForTurnMulti(*FinalAudioForTurnMulti) error
	VisitFinalMulti(*FinalMulti) error
	VisitError(*Error) error
}

func (r *ReceiveTextToDialogueWebsocketMessageMulti) Accept(visitor ReceiveTextToDialogueWebsocketMessageMultiVisitor) error {
	if r.typ == "AudioChunkMulti" || r.AudioChunkMulti != nil {
		return visitor.VisitAudioChunkMulti(r.AudioChunkMulti)
	}
	if r.typ == "FinalAudioForTurnMulti" || r.FinalAudioForTurnMulti != nil {
		return visitor.VisitFinalAudioForTurnMulti(r.FinalAudioForTurnMulti)
	}
	if r.typ == "FinalMulti" || r.FinalMulti != nil {
		return visitor.VisitFinalMulti(r.FinalMulti)
	}
	if r.typ == "Error" || r.Error != nil {
		return visitor.VisitError(r.Error)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", r)
}
