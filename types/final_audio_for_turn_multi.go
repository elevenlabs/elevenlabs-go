package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type FinalAudioForTurnMulti struct {
	// Signals the end of audio for a turn. Emitted after each explicit flush, new_turn/voice-switch boundary, and on connection close. All of the turn's audio, including any bytes buffered by streaming codecs (MP3, Opus, etc.), is delivered before this marker.
	IsFinalAudioForTurn *bool `json:"is_final_audio_for_turn,omitempty" xml:"is_final_audio_for_turn,omitempty"`
	// Identifies which context this response belongs to. Only set on the multi-context endpoint; omitted (None) on the single-context endpoint.
	ContextId       *string `json:"context_id,omitempty" xml:"context_id,omitempty"`
	extraProperties map[string]interface{}
}

func (f *FinalAudioForTurnMulti) GetIsFinalAudioForTurn() (value *bool) {
	if f == nil {
		return
	}
	return f.IsFinalAudioForTurn
}

func (f *FinalAudioForTurnMulti) GetContextId() (value *string) {
	if f == nil {
		return
	}
	return f.ContextId
}

func (f *FinalAudioForTurnMulti) SetIsFinalAudioForTurn(isFinalAudioForTurn *bool) {
	f.IsFinalAudioForTurn = isFinalAudioForTurn
}

func (f *FinalAudioForTurnMulti) SetContextId(contextId *string) {
	f.ContextId = contextId
}

func (f *FinalAudioForTurnMulti) GetExtraProperties() map[string]interface{} {
	if f == nil {
		return nil
	}
	return f.extraProperties
}

func (f FinalAudioForTurnMulti) String() string {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return "error converting struct: FinalAudioForTurnMulti to string"
	}
	return string(jsonData)
}

func (f *FinalAudioForTurnMulti) UnmarshalJSON(data []byte) error {
	type unmarshaler FinalAudioForTurnMulti
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*f = FinalAudioForTurnMulti(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *f)
	if err != nil {
		return err
	}
	f.extraProperties = extraProperties
	return nil
}

func (f *FinalAudioForTurnMulti) MarshalJSON() ([]byte, error) {
	type embed FinalAudioForTurnMulti
	var marshaler = struct {
		embed
	}{
		embed: embed(*f),
	}
	return json.Marshal(marshaler)
}
