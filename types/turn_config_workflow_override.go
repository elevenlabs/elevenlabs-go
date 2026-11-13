package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TurnConfigWorkflowOverride struct {
	// Maximum wait time for the user's reply before re-engaging the user
	TurnTimeout *float64 `json:"turn_timeout,omitempty" xml:"turn_timeout,omitempty"`
	// How long the agent will wait for the user to start the conversation if the first message is empty. If not set, uses the regular turn_timeout.
	InitialWaitTime *float64 `json:"initial_wait_time,omitempty" xml:"initial_wait_time,omitempty"`
	// Maximum wait time since the user last spoke before terminating the call
	SilenceEndCallTimeout *float64 `json:"silence_end_call_timeout,omitempty" xml:"silence_end_call_timeout,omitempty"`
	// Controls how eager the agent is to respond. Low = less eager (waits longer), Standard = default eagerness, High = more eager (responds sooner)
	TurnEagerness *TurnEagerness `json:"turn_eagerness,omitempty" xml:"turn_eagerness,omitempty"`
	// Controls if the agent should be more patient when user is spelling numbers and named entities. Auto = model based, Off = never wait extra
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
	// Configuration for soft timeout functionality. Provides immediate feedback during longer LLM responses.
	SoftTimeoutConfig *SoftTimeoutConfigWorkflowOverride `json:"soft_timeout_config,omitempty" xml:"soft_timeout_config,omitempty"`
	// What the agent says when it detects a noisy environment.
	NoisyEnvironmentMessage *string `json:"noisy_environment_message,omitempty" xml:"noisy_environment_message,omitempty"`
	extraProperties         map[string]interface{}
}

func (t *TurnConfigWorkflowOverride) GetTurnTimeout() (value *float64) {
	if t == nil {
		return
	}
	return t.TurnTimeout
}

func (t *TurnConfigWorkflowOverride) GetInitialWaitTime() (value *float64) {
	if t == nil {
		return
	}
	return t.InitialWaitTime
}

func (t *TurnConfigWorkflowOverride) GetSilenceEndCallTimeout() (value *float64) {
	if t == nil {
		return
	}
	return t.SilenceEndCallTimeout
}

func (t *TurnConfigWorkflowOverride) GetTurnEagerness() (value *TurnEagerness) {
	if t == nil {
		return
	}
	return t.TurnEagerness
}

func (t *TurnConfigWorkflowOverride) GetSpellingPatience() (value *SpellingPatience) {
	if t == nil {
		return
	}
	return t.SpellingPatience
}

func (t *TurnConfigWorkflowOverride) GetSpeculativeTurn() (value *bool) {
	if t == nil {
		return
	}
	return t.SpeculativeTurn
}

func (t *TurnConfigWorkflowOverride) GetRetranscribeOnTurnTimeout() (value *bool) {
	if t == nil {
		return
	}
	return t.RetranscribeOnTurnTimeout
}

func (t *TurnConfigWorkflowOverride) GetTurnModel() (value *TurnModel) {
	if t == nil {
		return
	}
	return t.TurnModel
}

func (t *TurnConfigWorkflowOverride) GetInterruptionIgnoreTerms() (value []string) {
	if t == nil {
		return
	}
	return t.InterruptionIgnoreTerms
}

func (t *TurnConfigWorkflowOverride) GetInterruptionIgnoreTermLanguages() (value []string) {
	if t == nil {
		return
	}
	return t.InterruptionIgnoreTermLanguages
}

func (t *TurnConfigWorkflowOverride) GetMergeWithDefaultIgnoreTerms() (value *bool) {
	if t == nil {
		return
	}
	return t.MergeWithDefaultIgnoreTerms
}

func (t *TurnConfigWorkflowOverride) GetTranscribeOnDisabledInterruptions() (value *bool) {
	if t == nil {
		return
	}
	return t.TranscribeOnDisabledInterruptions
}

func (t *TurnConfigWorkflowOverride) GetSoftTimeoutConfig() (value *SoftTimeoutConfigWorkflowOverride) {
	if t == nil {
		return
	}
	return t.SoftTimeoutConfig
}

func (t *TurnConfigWorkflowOverride) GetNoisyEnvironmentMessage() (value *string) {
	if t == nil {
		return
	}
	return t.NoisyEnvironmentMessage
}

func (t *TurnConfigWorkflowOverride) SetTurnTimeout(turnTimeout *float64) {
	t.TurnTimeout = turnTimeout
}

func (t *TurnConfigWorkflowOverride) SetInitialWaitTime(initialWaitTime *float64) {
	t.InitialWaitTime = initialWaitTime
}

func (t *TurnConfigWorkflowOverride) SetSilenceEndCallTimeout(silenceEndCallTimeout *float64) {
	t.SilenceEndCallTimeout = silenceEndCallTimeout
}

func (t *TurnConfigWorkflowOverride) SetTurnEagerness(turnEagerness *TurnEagerness) {
	t.TurnEagerness = turnEagerness
}

func (t *TurnConfigWorkflowOverride) SetSpellingPatience(spellingPatience *SpellingPatience) {
	t.SpellingPatience = spellingPatience
}

func (t *TurnConfigWorkflowOverride) SetSpeculativeTurn(speculativeTurn *bool) {
	t.SpeculativeTurn = speculativeTurn
}

func (t *TurnConfigWorkflowOverride) SetRetranscribeOnTurnTimeout(retranscribeOnTurnTimeout *bool) {
	t.RetranscribeOnTurnTimeout = retranscribeOnTurnTimeout
}

func (t *TurnConfigWorkflowOverride) SetTurnModel(turnModel *TurnModel) {
	t.TurnModel = turnModel
}

func (t *TurnConfigWorkflowOverride) SetInterruptionIgnoreTerms(interruptionIgnoreTerms []string) {
	t.InterruptionIgnoreTerms = interruptionIgnoreTerms
}

func (t *TurnConfigWorkflowOverride) SetInterruptionIgnoreTermLanguages(interruptionIgnoreTermLanguages []string) {
	t.InterruptionIgnoreTermLanguages = interruptionIgnoreTermLanguages
}

func (t *TurnConfigWorkflowOverride) SetMergeWithDefaultIgnoreTerms(mergeWithDefaultIgnoreTerms *bool) {
	t.MergeWithDefaultIgnoreTerms = mergeWithDefaultIgnoreTerms
}

func (t *TurnConfigWorkflowOverride) SetTranscribeOnDisabledInterruptions(transcribeOnDisabledInterruptions *bool) {
	t.TranscribeOnDisabledInterruptions = transcribeOnDisabledInterruptions
}

func (t *TurnConfigWorkflowOverride) SetSoftTimeoutConfig(softTimeoutConfig *SoftTimeoutConfigWorkflowOverride) {
	t.SoftTimeoutConfig = softTimeoutConfig
}

func (t *TurnConfigWorkflowOverride) SetNoisyEnvironmentMessage(noisyEnvironmentMessage *string) {
	t.NoisyEnvironmentMessage = noisyEnvironmentMessage
}

func (t *TurnConfigWorkflowOverride) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TurnConfigWorkflowOverride) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TurnConfigWorkflowOverride to string"
	}
	return string(jsonData)
}

func (t *TurnConfigWorkflowOverride) UnmarshalJSON(data []byte) error {
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

func (t *TurnConfigWorkflowOverride) MarshalJSON() ([]byte, error) {
	type embed TurnConfigWorkflowOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
