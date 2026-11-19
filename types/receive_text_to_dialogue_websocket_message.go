package types

import (
	"encoding/json"
	"fmt"
)

type ReceiveTextToDialogueWebsocketMessage struct {
	AudioChunk        *AudioChunk
	FinalAudioForTurn *FinalAudioForTurn
	Final             *Final
	Error             *Error
	typ               string
}

func (r *ReceiveTextToDialogueWebsocketMessage) GetAudioChunk() *AudioChunk {
	if r == nil {
		return nil
	}
	if r.AudioChunk == nil {
		return nil
	}
	return r.AudioChunk
}

func (r *ReceiveTextToDialogueWebsocketMessage) GetFinalAudioForTurn() *FinalAudioForTurn {
	if r == nil {
		return nil
	}
	if r.FinalAudioForTurn == nil {
		return nil
	}
	return r.FinalAudioForTurn
}

func (r *ReceiveTextToDialogueWebsocketMessage) GetFinal() *Final {
	if r == nil {
		return nil
	}
	if r.Final == nil {
		return nil
	}
	return r.Final
}

func (r *ReceiveTextToDialogueWebsocketMessage) GetError() *Error {
	if r == nil {
		return nil
	}
	if r.Error == nil {
		return nil
	}
	return r.Error
}

func (r *ReceiveTextToDialogueWebsocketMessage) UnmarshalJSON(data []byte) error {
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
	if hasAnyKey("audio", "alignment", "normalized_alignment") {
		var valueAudioChunk *AudioChunk
		if err := json.Unmarshal(data, &valueAudioChunk); err == nil {
			r.typ = "AudioChunk"
			r.AudioChunk = valueAudioChunk
			return nil
		}
	}
	if hasAnyKey("is_final_audio_for_turn") {
		var valueFinalAudioForTurn *FinalAudioForTurn
		if err := json.Unmarshal(data, &valueFinalAudioForTurn); err == nil {
			r.typ = "FinalAudioForTurn"
			r.FinalAudioForTurn = valueFinalAudioForTurn
			return nil
		}
	}
	var valueFinal *Final
	if err := json.Unmarshal(data, &valueFinal); err == nil {
		r.typ = "Final"
		r.Final = valueFinal
		return nil
	}
	var valueError *Error
	if err := json.Unmarshal(data, &valueError); err == nil {
		r.typ = "Error"
		r.Error = valueError
		return nil
	}

	// Retried without their key guards, for a payload that matched none of them.
	var valueAudioChunkFallback *AudioChunk
	if err := json.Unmarshal(data, &valueAudioChunkFallback); err == nil {
		r.typ = "AudioChunk"
		r.AudioChunk = valueAudioChunkFallback
		return nil
	}
	var valueFinalAudioForTurnFallback *FinalAudioForTurn
	if err := json.Unmarshal(data, &valueFinalAudioForTurnFallback); err == nil {
		r.typ = "FinalAudioForTurn"
		r.FinalAudioForTurn = valueFinalAudioForTurnFallback
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, r)
}

func (r ReceiveTextToDialogueWebsocketMessage) MarshalJSON() ([]byte, error) {
	if r.typ == "AudioChunk" || r.AudioChunk != nil {
		return json.Marshal(r.AudioChunk)
	}
	if r.typ == "FinalAudioForTurn" || r.FinalAudioForTurn != nil {
		return json.Marshal(r.FinalAudioForTurn)
	}
	if r.typ == "Final" || r.Final != nil {
		return json.Marshal(r.Final)
	}
	if r.typ == "Error" || r.Error != nil {
		return json.Marshal(r.Error)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", r)
}

type ReceiveTextToDialogueWebsocketMessageVisitor interface {
	VisitAudioChunk(*AudioChunk) error
	VisitFinalAudioForTurn(*FinalAudioForTurn) error
	VisitFinal(*Final) error
	VisitError(*Error) error
}

func (r *ReceiveTextToDialogueWebsocketMessage) Accept(visitor ReceiveTextToDialogueWebsocketMessageVisitor) error {
	if r.typ == "AudioChunk" || r.AudioChunk != nil {
		return visitor.VisitAudioChunk(r.AudioChunk)
	}
	if r.typ == "FinalAudioForTurn" || r.FinalAudioForTurn != nil {
		return visitor.VisitFinalAudioForTurn(r.FinalAudioForTurn)
	}
	if r.typ == "Final" || r.Final != nil {
		return visitor.VisitFinal(r.Final)
	}
	if r.typ == "Error" || r.Error != nil {
		return visitor.VisitError(r.Error)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", r)
}
