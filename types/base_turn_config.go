package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BaseTurnConfig struct {
	// Maximum wait time for the user's reply before re-engaging the user
	TurnTimeout *float64 `json:"turn_timeout,omitempty" xml:"turn_timeout,omitempty"`
	// How long the agent will wait for the user to start the conversation if the first message is empty. If not set, uses the regular turn_timeout.
	InitialWaitTime *float64 `json:"initial_wait_time,omitempty" xml:"initial_wait_time,omitempty"`
	// Maximum wait time since the user last spoke before terminating the call
	SilenceEndCallTimeout *float64 `json:"silence_end_call_timeout,omitempty" xml:"silence_end_call_timeout,omitempty"`
	// Agent's eagerness to respond. Higher values make agent wait for higher turn probability.
	TurnEagerness *TurnEagerness `json:"turn_eagerness,omitempty" xml:"turn_eagerness,omitempty"`
	// Controls if the agent should be more patient when user is spelling numbers and named entities.
	SpellingPatience *SpellingPatience `json:"spelling_patience,omitempty" xml:"spelling_patience,omitempty"`
	// When enabled, starts generating LLM responses during silence before full turn confidence is reached, reducing perceived latency. May increase LLM costs.
	SpeculativeTurn *bool `json:"speculative_turn,omitempty" xml:"speculative_turn,omitempty"`
	// When enabled, if VAD detects no speech, attempts to re-transcribe accumulated audio at turn timeout. Disables silence discount billing while enabled for the active agent or workflow node.
	RetranscribeOnTurnTimeout *bool `json:"retranscribe_on_turn_timeout,omitempty" xml:"retranscribe_on_turn_timeout,omitempty"`
	// Version of the turn detection model to use.
	TurnModel *TurnModel `json:"turn_model,omitempty" xml:"turn_model,omitempty"`
	// List of terms that should not trigger an interruption when spoken by the user (e.g. 'gotcha', 'understood'). Uses case-insensitive exact matching.
	InterruptionIgnoreTerms []string `json:"interruption_ignore_terms,omitempty" xml:"interruption_ignore_terms,omitempty"`
	// Language codes for which preset ignore-term categories have been activated. Stored explicitly so display is not inferred from term overlap.
	InterruptionIgnoreTermLanguages []string `json:"interruption_ignore_term_languages,omitempty" xml:"interruption_ignore_term_languages,omitempty"`
	// When enabled, the curated default terms for interruption_ignore_term_languages are used in addition to interruption_ignore_terms.
	MergeWithDefaultIgnoreTerms *bool `json:"merge_with_default_ignore_terms,omitempty" xml:"merge_with_default_ignore_terms,omitempty"`
	// When interruptions are disabled, still transcribe what the user says so it can carry into the next turn. When off, user speech during a non-interruptible turn is ignored and won't trigger a turn.
	TranscribeOnDisabledInterruptions *bool `json:"transcribe_on_disabled_interruptions,omitempty" xml:"transcribe_on_disabled_interruptions,omitempty"`
	extraProperties                   map[string]interface{}
}

func (b *BaseTurnConfig) GetTurnTimeout() (value *float64) {
	if b == nil {
		return
	}
	return b.TurnTimeout
}

func (b *BaseTurnConfig) GetInitialWaitTime() (value *float64) {
	if b == nil {
		return
	}
	return b.InitialWaitTime
}

func (b *BaseTurnConfig) GetSilenceEndCallTimeout() (value *float64) {
	if b == nil {
		return
	}
	return b.SilenceEndCallTimeout
}

func (b *BaseTurnConfig) GetTurnEagerness() (value *TurnEagerness) {
	if b == nil {
		return
	}
	return b.TurnEagerness
}

func (b *BaseTurnConfig) GetSpellingPatience() (value *SpellingPatience) {
	if b == nil {
		return
	}
	return b.SpellingPatience
}

func (b *BaseTurnConfig) GetSpeculativeTurn() (value *bool) {
	if b == nil {
		return
	}
	return b.SpeculativeTurn
}

func (b *BaseTurnConfig) GetRetranscribeOnTurnTimeout() (value *bool) {
	if b == nil {
		return
	}
	return b.RetranscribeOnTurnTimeout
}

func (b *BaseTurnConfig) GetTurnModel() (value *TurnModel) {
	if b == nil {
		return
	}
	return b.TurnModel
}

func (b *BaseTurnConfig) GetInterruptionIgnoreTerms() (value []string) {
	if b == nil {
		return
	}
	return b.InterruptionIgnoreTerms
}

func (b *BaseTurnConfig) GetInterruptionIgnoreTermLanguages() (value []string) {
	if b == nil {
		return
	}
	return b.InterruptionIgnoreTermLanguages
}

func (b *BaseTurnConfig) GetMergeWithDefaultIgnoreTerms() (value *bool) {
	if b == nil {
		return
	}
	return b.MergeWithDefaultIgnoreTerms
}

func (b *BaseTurnConfig) GetTranscribeOnDisabledInterruptions() (value *bool) {
	if b == nil {
		return
	}
	return b.TranscribeOnDisabledInterruptions
}

func (b *BaseTurnConfig) SetTurnTimeout(turnTimeout *float64) {
	b.TurnTimeout = turnTimeout
}

func (b *BaseTurnConfig) SetInitialWaitTime(initialWaitTime *float64) {
	b.InitialWaitTime = initialWaitTime
}

func (b *BaseTurnConfig) SetSilenceEndCallTimeout(silenceEndCallTimeout *float64) {
	b.SilenceEndCallTimeout = silenceEndCallTimeout
}

func (b *BaseTurnConfig) SetTurnEagerness(turnEagerness *TurnEagerness) {
	b.TurnEagerness = turnEagerness
}

func (b *BaseTurnConfig) SetSpellingPatience(spellingPatience *SpellingPatience) {
	b.SpellingPatience = spellingPatience
}

func (b *BaseTurnConfig) SetSpeculativeTurn(speculativeTurn *bool) {
	b.SpeculativeTurn = speculativeTurn
}

func (b *BaseTurnConfig) SetRetranscribeOnTurnTimeout(retranscribeOnTurnTimeout *bool) {
	b.RetranscribeOnTurnTimeout = retranscribeOnTurnTimeout
}

func (b *BaseTurnConfig) SetTurnModel(turnModel *TurnModel) {
	b.TurnModel = turnModel
}

func (b *BaseTurnConfig) SetInterruptionIgnoreTerms(interruptionIgnoreTerms []string) {
	b.InterruptionIgnoreTerms = interruptionIgnoreTerms
}

func (b *BaseTurnConfig) SetInterruptionIgnoreTermLanguages(interruptionIgnoreTermLanguages []string) {
	b.InterruptionIgnoreTermLanguages = interruptionIgnoreTermLanguages
}

func (b *BaseTurnConfig) SetMergeWithDefaultIgnoreTerms(mergeWithDefaultIgnoreTerms *bool) {
	b.MergeWithDefaultIgnoreTerms = mergeWithDefaultIgnoreTerms
}

func (b *BaseTurnConfig) SetTranscribeOnDisabledInterruptions(transcribeOnDisabledInterruptions *bool) {
	b.TranscribeOnDisabledInterruptions = transcribeOnDisabledInterruptions
}

func (b *BaseTurnConfig) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BaseTurnConfig) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BaseTurnConfig to string"
	}
	return string(jsonData)
}

func (b *BaseTurnConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BaseTurnConfig) MarshalJSON() ([]byte, error) {
	type embed BaseTurnConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
