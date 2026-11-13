package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TurnConfig struct {
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
	// Configuration for soft timeout functionality during LLM response generation.
	SoftTimeoutConfig *SoftTimeoutConfig `json:"soft_timeout_config,omitempty" xml:"soft_timeout_config,omitempty"`
	// When enabled, the agent politely suggests the user move somewhere quieter when it hears other voices in the background. This improves transcription and turn-taking accuracy.
	NoisyEnvironmentDetection *bool `json:"noisy_environment_detection,omitempty" xml:"noisy_environment_detection,omitempty"`
	// What the agent says when it detects a noisy environment.
	NoisyEnvironmentMessage *string `json:"noisy_environment_message,omitempty" xml:"noisy_environment_message,omitempty"`
	extraProperties         map[string]interface{}
}

func (t *TurnConfig) GetTurnTimeout() (value *float64) {
	if t == nil {
		return
	}
	return t.TurnTimeout
}

func (t *TurnConfig) GetInitialWaitTime() (value *float64) {
	if t == nil {
		return
	}
	return t.InitialWaitTime
}

func (t *TurnConfig) GetSilenceEndCallTimeout() (value *float64) {
	if t == nil {
		return
	}
	return t.SilenceEndCallTimeout
}

func (t *TurnConfig) GetTurnEagerness() (value *TurnEagerness) {
	if t == nil {
		return
	}
	return t.TurnEagerness
}

func (t *TurnConfig) GetSpellingPatience() (value *SpellingPatience) {
	if t == nil {
		return
	}
	return t.SpellingPatience
}

func (t *TurnConfig) GetSpeculativeTurn() (value *bool) {
	if t == nil {
		return
	}
	return t.SpeculativeTurn
}

func (t *TurnConfig) GetRetranscribeOnTurnTimeout() (value *bool) {
	if t == nil {
		return
	}
	return t.RetranscribeOnTurnTimeout
}

func (t *TurnConfig) GetTurnModel() (value *TurnModel) {
	if t == nil {
		return
	}
	return t.TurnModel
}

func (t *TurnConfig) GetInterruptionIgnoreTerms() (value []string) {
	if t == nil {
		return
	}
	return t.InterruptionIgnoreTerms
}

func (t *TurnConfig) GetInterruptionIgnoreTermLanguages() (value []string) {
	if t == nil {
		return
	}
	return t.InterruptionIgnoreTermLanguages
}

func (t *TurnConfig) GetMergeWithDefaultIgnoreTerms() (value *bool) {
	if t == nil {
		return
	}
	return t.MergeWithDefaultIgnoreTerms
}

func (t *TurnConfig) GetTranscribeOnDisabledInterruptions() (value *bool) {
	if t == nil {
		return
	}
	return t.TranscribeOnDisabledInterruptions
}

func (t *TurnConfig) GetSoftTimeoutConfig() (value *SoftTimeoutConfig) {
	if t == nil {
		return
	}
	return t.SoftTimeoutConfig
}

func (t *TurnConfig) GetNoisyEnvironmentDetection() (value *bool) {
	if t == nil {
		return
	}
	return t.NoisyEnvironmentDetection
}

func (t *TurnConfig) GetNoisyEnvironmentMessage() (value *string) {
	if t == nil {
		return
	}
	return t.NoisyEnvironmentMessage
}

func (t *TurnConfig) SetTurnTimeout(turnTimeout *float64) {
	t.TurnTimeout = turnTimeout
}

func (t *TurnConfig) SetInitialWaitTime(initialWaitTime *float64) {
	t.InitialWaitTime = initialWaitTime
}

func (t *TurnConfig) SetSilenceEndCallTimeout(silenceEndCallTimeout *float64) {
	t.SilenceEndCallTimeout = silenceEndCallTimeout
}

func (t *TurnConfig) SetTurnEagerness(turnEagerness *TurnEagerness) {
	t.TurnEagerness = turnEagerness
}

func (t *TurnConfig) SetSpellingPatience(spellingPatience *SpellingPatience) {
	t.SpellingPatience = spellingPatience
}

func (t *TurnConfig) SetSpeculativeTurn(speculativeTurn *bool) {
	t.SpeculativeTurn = speculativeTurn
}

func (t *TurnConfig) SetRetranscribeOnTurnTimeout(retranscribeOnTurnTimeout *bool) {
	t.RetranscribeOnTurnTimeout = retranscribeOnTurnTimeout
}

func (t *TurnConfig) SetTurnModel(turnModel *TurnModel) {
	t.TurnModel = turnModel
}

func (t *TurnConfig) SetInterruptionIgnoreTerms(interruptionIgnoreTerms []string) {
	t.InterruptionIgnoreTerms = interruptionIgnoreTerms
}

func (t *TurnConfig) SetInterruptionIgnoreTermLanguages(interruptionIgnoreTermLanguages []string) {
	t.InterruptionIgnoreTermLanguages = interruptionIgnoreTermLanguages
}

func (t *TurnConfig) SetMergeWithDefaultIgnoreTerms(mergeWithDefaultIgnoreTerms *bool) {
	t.MergeWithDefaultIgnoreTerms = mergeWithDefaultIgnoreTerms
}

func (t *TurnConfig) SetTranscribeOnDisabledInterruptions(transcribeOnDisabledInterruptions *bool) {
	t.TranscribeOnDisabledInterruptions = transcribeOnDisabledInterruptions
}

func (t *TurnConfig) SetSoftTimeoutConfig(softTimeoutConfig *SoftTimeoutConfig) {
	t.SoftTimeoutConfig = softTimeoutConfig
}

func (t *TurnConfig) SetNoisyEnvironmentDetection(noisyEnvironmentDetection *bool) {
	t.NoisyEnvironmentDetection = noisyEnvironmentDetection
}

func (t *TurnConfig) SetNoisyEnvironmentMessage(noisyEnvironmentMessage *string) {
	t.NoisyEnvironmentMessage = noisyEnvironmentMessage
}

func (t *TurnConfig) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TurnConfig) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TurnConfig to string"
	}
	return string(jsonData)
}

func (t *TurnConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TurnConfig) MarshalJSON() ([]byte, error) {
	type embed TurnConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
